# Code Quality Review - December 2025

A comprehensive review of DRY (Don't Repeat Yourself), SRP (Single Responsibility Principle), and best practices violations in the Go Fundamentals & Best Practices codebase.

---

## Executive Summary

| Area | DRY Violations | SRP Violations | Best Practice Issues | Severity |
|------|----------------|----------------|----------------------|----------|
| **Backend (Go)** | 10+ | 3 | 12+ | Medium-High |
| **Frontend (Vue)** | 12+ | 5 | 14+ | High |

**Top Priority Issues:**
1. Markdown rendering logic duplicated 4 times across frontend components
2. Missing entry point (`cmd/server/main.go`)
3. User ID extraction repeated in 4 backend handlers
4. Tutorial lookup logic repeated in 3 handlers
5. SectionViewer.vue has 10+ responsibilities

---

## Backend (Go) Issues

### 🔴 Critical: Missing Entry Point

**Location:** Project references `cmd/server/main.go` in `Taskfile.yml` but file doesn't exist.

**Impact:** Backend cannot be started with `task dev:backend`.

---

### DRY Violations

#### 1. Repeated User ID Extraction (`internal/api/handlers.go`)

**Lines:** 165-168, 187-190, 217-220, 265-269

```go
// This pattern appears 4 times:
userID := r.URL.Query().Get("userId")
if userID == "" {
    userID = DefaultUserID
}
```

**Fix:** Extract to helper function:
```go
func extractUserID(r *http.Request) string {
    if userID := r.URL.Query().Get("userId"); userID != "" {
        return userID
    }
    return DefaultUserID
}
```

---

#### 2. Repeated Tutorial Lookup (`internal/api/handlers.go`)

**Lines:** 79-84, 101-106, 234-244

```go
// Pattern repeated 3 times:
for _, tutorial := range h.tutorials {
    if tutorial.ID == tutorialID {
        respondJSON(w, tutorial)
        return
    }
}
```

**Fix:** Create helper method:
```go
func (h *Handlers) findTutorial(tutorialID string) *models.Tutorial {
    for i := range h.tutorials {
        if h.tutorials[i].ID == tutorialID {
            return &h.tutorials[i]
        }
    }
    return nil
}
```

---

#### 3. Error Response Duplication (`internal/api/handlers.go`)

13+ instances of similar error responses:
```go
http.Error(w, "tutorial not found", http.StatusNotFound)     // 4 times
http.Error(w, "method not allowed", http.StatusMethodNotAllowed)  // 2 times
http.Error(w, "tutorial ID required", http.StatusBadRequest) // 3 times
```

**Fix:** Create error response helpers:
```go
func respondNotFound(w http.ResponseWriter, resource string) {
    http.Error(w, resource+" not found", http.StatusNotFound)
}
```

---

#### 4. Progress Initialization (`internal/storage/progress.go`)

**Lines:** 68-78, 101-111

```go
// Identical initialization in MarkSectionComplete and MarkExerciseComplete:
if s.progress[userID] == nil {
    s.progress[userID] = &models.Progress{
        UserID:             userID,
        CompletedSections:  make(map[string][]string),
        CompletedExercises: make(map[string][]string),
    }
}
```

**Fix:** Extract to `ensureUserProgress(userID string) *models.Progress`

---

#### 5. Pattern Validation (`internal/executor/validator.go`)

**Lines:** 127-140, 143-155

Network and file operation validation use identical loops:
```go
for _, pattern := range networkPatterns {
    if strings.Contains(code, pattern) {
        v.logger.Warn(...)
        return fmt.Errorf(...)
    }
}
```

**Fix:** Generic pattern checker:
```go
func (v *Validator) checkPatterns(code string, patterns []string, errType string) error
```

---

#### 6. List Extraction in Parsers (`internal/parser/section.go`)

**Lines:** 243-288, 340-371

`extractTopics()` and `extractTeachingPoints()` have nearly identical structures.

**Fix:** Create generic `extractListSection(content, headerPattern string) []string`

---

### SRP Violations

#### 1. `GetExercisesByTutorialID` Does Too Much (`internal/api/handlers.go:231-260`)

Single function handles:
- Tutorial lookup in cache
- File system listing
- File matching by name
- Exercise parsing

**Fix:** Delegate file operations to parser; keep handler focused on HTTP concerns.

---

#### 2. `handleTutorialRoutes` Mixed Concerns (`internal/api/routes.go:39-56`)

Mixes routing decisions with delegation logic.

**Fix:** Use a proper router (chi, gorilla/mux) or cleaner switch statement.

---

### Best Practices Issues

| Issue | Location | Recommendation |
|-------|----------|----------------|
| Regex compiled per call | `section.go:299` | Move to module-level `var` |
| Silent error in sscanf | `tutorial.go:103` | Handle parse errors |
| Wildcard CORS | `routes.go:76` | Restrict origins in production |
| Incomplete import validation | `validator.go:177-186` | Use AST or better regex |
| Magic numbers undocumented | `markdown.go:14-18` | Add explanatory comments |
| Error ignored in JSON encode | `handlers.go:277` | Log encoding errors |

---

## Frontend (Vue/TypeScript) Issues

### DRY Violations

#### 1. Markdown Rendering (4 Copies!) ⚠️ Critical

**Locations:**
- `SectionViewer.vue:243-259` - `renderMarkdown()`
- `SectionViewer.vue:262-358` - `renderMarkdownContent()` (unused!)
- `InstructorPanel.vue:56-82` - `renderedNotes` computed

All contain identical regex patterns:
```typescript
// Link conversion (duplicated 3 times)
html = html.replace(
  /\[([^\]]+)\]\(([^)]+)\)/g,
  '<a href="$2" ... class="text-[#00ADD8] hover:text-[#007D9C] ...">$1</a>'
);

// Inline code (duplicated 3 times)
html = html.replace(/`([^`]+)`/g, '<code class="...">$1</code>');

// Bold text (duplicated 3 times)
html = html.replace(/\*\*(.+?)\*\*/g, '<strong class="font-semibold">$1</strong>');
```

**Fix:** Create `composables/useMarkdownRenderer.ts`:
```typescript
export function useMarkdownRenderer() {
  const renderMarkdown = (content: string): string => {
    // Single implementation of all regex transformations
  };
  return { renderMarkdown };
}
```

---

#### 2. Copy-to-Clipboard (2 Copies)

**Locations:**
- `CodeRunner.vue:178-184`
- `CodeEditor.vue:125-135`

```typescript
// Identical implementation in both files:
const copyCode = async () => {
  try {
    await navigator.clipboard.writeText(code.value);
    copied.value = true;
    setTimeout(() => { copied.value = false; }, 2000);
  } catch (err) {
    console.error('Failed to copy:', err);
  }
};
```

**Fix:** Create `composables/useCopyToClipboard.ts`

---

#### 3. Index Clamping (`TutorialViewer.vue`)

**Lines:** 186, 249

```typescript
// Same expression twice:
const clampedIndex = Math.max(0, Math.min(index, tutorial.value.sections.length - 1));
```

**Fix:** Extract to utility function:
```typescript
const clampIndex = (index: number, max: number) => Math.max(0, Math.min(index, max));
```

---

#### 4. Progress Initialization (`stores/progress.ts`)

**Lines:** 48-54, 88-95

Same initialization block duplicated:
```typescript
if (!progress.value) {
  progress.value = {
    userId,
    completedSections: {},
    completedExercises: {},
    lastAccessed: new Date().toISOString(),
  };
}
```

**Fix:** Create `initializeProgress(userId: string)` helper

---

#### 5. localStorage Persistence (`stores/progress.ts`)

**Lines:** 27, 73, 103

Same save pattern repeated 3 times:
```typescript
localStorage.setItem('tutorial-progress', JSON.stringify(progress.value));
```

**Fix:** Create `persistProgress()` helper or use Pinia persist plugin

---

### SRP Violations

#### 1. `SectionViewer.vue` (359 lines) ⚠️ Critical

**Current responsibilities (10+):**
- Parse and render Table of Contents
- Render section topics
- Render code examples via CodeRunner
- Render teaching points
- Handle ToC click navigation
- Track section completion
- Multiple markdown rendering functions
- Inline styling decisions

**Fix:** Split into:
- `SectionHeader.vue` - Title, completion status
- `TableOfContents.vue` - ToC parsing and display
- `SectionContent.vue` - Content rendering
- `TeachingPoints.vue` - Teaching points list

---

#### 2. `TutorialViewer.vue` (263 lines)

**Responsibilities:**
- Tutorial loading
- Section navigation (3 functions)
- Progress tracking
- Instructor mode toggle
- Route watching (2 watchers)
- Local storage persistence

**Fix:** Extract navigation logic to `useTutorialNavigation()` composable

---

### Best Practices Issues

#### Anti-Patterns

| Issue | Location | Recommendation |
|-------|----------|----------------|
| `alert()` for feedback | `ExerciseView.vue:169,171,175,177` | Use toast notifications |
| Module-level mutable state | `useSyntaxHighlight.ts:4` | Use Vue reactive state |
| Unused composable | `useRetry.ts` (entire file) | Remove or implement |
| Linear "exponential" backoff | `useRetry.ts:40` | Fix: `delay * 2^i` |
| Unused function | `SectionViewer.vue:renderMarkdownContent()` | Remove dead code |

---

#### Type Safety Issues

| Issue | Location | Recommendation |
|-------|----------|----------------|
| Unsafe parseInt | `TutorialView.vue:4` | Add undefined guard |
| Inconsistent nullability | `types/tutorial.ts:25-27` | Use `undefined` consistently |
| Unused type field | `types/progress.ts:20` | Remove `sectionProgress` or implement |

---

#### Styling Inconsistencies

Hard-coded colors appear 15+ times instead of using Tailwind config:
```typescript
// Found throughout components:
class="text-[#00ADD8]"  // Go blue
class="hover:text-[#007D9C]"  // Go dark blue
class="from-[#00ADD8] to-[#5DC9E2]"  // Gradient
```

**Fix:** Add to `tailwind.config.js`:
```javascript
theme: {
  extend: {
    colors: {
      'go-blue': '#00ADD8',
      'go-blue-dark': '#007D9C',
      'go-blue-light': '#5DC9E2',
    }
  }
}
```

---

## Priority Action Items

### Tier 1 - Critical (Should fix immediately)

1. **Create `cmd/server/main.go`** - Backend cannot start
2. **Extract markdown renderer composable** - 4 duplicated implementations
3. **Split SectionViewer.vue** - 10+ responsibilities
4. **Create user ID extraction helper** - 4 duplicates in handlers
5. **Create tutorial lookup helper** - 3 duplicates in handlers

### Tier 2 - Important (Fix soon)

6. Extract copy-to-clipboard composable
7. Extract progress initialization helpers (both Go and Vue)
8. Replace `alert()` with toast notifications
9. Fix CORS wildcard for production
10. Compile regex patterns at module level

### Tier 3 - Nice-to-have

11. Create Tailwind color scheme constants
12. Remove unused `useRetry.ts` or implement properly
13. Add proper error logging to JSON encoding
14. Document magic numbers in parser
15. Remove `renderMarkdownContent()` dead code

---

## Estimated Impact

| Refactoring | Lines Removed | Maintainability Improvement |
|-------------|---------------|----------------------------|
| Markdown renderer extraction | ~150 lines | High |
| User ID helper | ~12 lines | Medium |
| Tutorial lookup helper | ~18 lines | Medium |
| Copy-to-clipboard extraction | ~20 lines | Medium |
| SectionViewer split | Reorganized | High |
| Progress helpers | ~30 lines | Medium |

**Total estimated duplicate code:** ~230+ lines

---

## Appendix: Go Linting Alignment

The codebase mostly follows `CLAUDE.md` guidelines but could improve on:

- [ ] Magic numbers should use named constants
- [ ] Regex should be compiled at module level
- [ ] Error handling should be consistent (no `_ =` ignores)
- [ ] Switch statements preferred over if-else chains for HTTP methods

---

*Review conducted December 2025*
