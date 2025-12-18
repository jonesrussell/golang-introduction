# Technical Stack

This document details the complete technical stack used in the Go Tutorials platform.

## Backend Stack

### Core Language & Runtime
- **Go 1.25+** ✅
  - Version: `1.25.5` (as specified in `go.mod`)
  - Standard library focused approach

### HTTP Server & Routing
- **Standard library `net/http`** ✅
  - Using `http.ServeMux` for routing
  - Custom route handlers with path parameter parsing
  - CORS middleware implementation
  - Alternative option available: `github.com/gorilla/mux` (not currently used, but compatible)

### Markdown Parsing
- **github.com/yuin/goldmark** ✅
  - Version: `v1.7.13`
  - AST-based markdown parsing
  - Located in: `internal/parser/goldmark_parser.go`
  - Also includes regex-based parser for compatibility

### Code Execution
- **os/exec** ✅
  - Standard library for executing Go code
  - `exec.CommandContext` for timeout support
  - Isolated execution in temporary directories
  - Located in: `internal/executor/runner.go`

### Additional Backend Dependencies
- Standard library packages:
  - `context` - For cancellation and timeouts
  - `encoding/json` - For JSON serialization
  - `os` - For file operations
  - `path/filepath` - For path manipulation
  - `regexp` - For pattern matching
  - `strings` - For string manipulation
  - `sync` - For concurrency control
  - `time` - For timeouts and durations

## Frontend Stack

### Core Framework
- **Vue 3** ✅
  - Version: `^3.5.25`
  - Composition API (not Options API)
  - TypeScript support
  - Located in: `frontend/src/`

### Build Tooling
- **Vite** ✅
  - Version: `^7.2.4`
  - Fast HMR (Hot Module Replacement)
  - Optimized production builds
  - Configuration: `frontend/vite.config.ts`

### Type System
- **TypeScript** ✅
  - Version: `~5.9.0`
  - Strict type checking
  - Type definitions for all components
  - Type-safe API calls

### State Management
- **Pinia** ✅
  - Version: `^3.0.4`
  - Vue 3's official state management
  - Stores: `frontend/src/stores/progress.ts`
  - Used for: Progress tracking, tutorial state

### Routing
- **Vue Router** ✅
  - Version: `^4.6.3`
  - Client-side routing
  - Configuration: `frontend/src/router/index.ts`

### Utility Libraries
- **@vueuse/core** ⚠️
  - Version: `^14.1.0`
  - Installed but not currently used
  - Vue composition utilities available for future use
  - Ready-to-use composables library

### Code Syntax Highlighting
- **Shiki** ✅
  - Version: `^3.20.0`
  - Syntax highlighting library
  - Used in: `frontend/src/composables/useSyntaxHighlight.ts`
  - Supports: Go, JavaScript, TypeScript, JSON, Markdown
  - Theme: GitHub Dark

### Code Editing
- **Monaco Editor** ⚠️
  - Version: `^0.55.1`
  - Installed but not currently used
  - Currently using simple `<textarea>` in `CodeEditor.vue`
  - Available for future upgrade to advanced editing features
  - Alternative: CodeMirror (not installed)
  - Location: `frontend/src/components/CodeEditor.vue`

### HTTP Client
- **Axios** ✅
  - Version: `^1.13.2`
  - Promise-based HTTP client
  - API service layer: `frontend/src/services/api.ts`
  - Request/response interceptors support

### Styling
- **Tailwind CSS** ✅
  - Version: `^4.1.18`
  - Utility-first CSS framework
  - Configuration: `frontend/tailwind.config.js`
  - Responsive design utilities
  - Dark mode support ready

### Development Tools
- **vue-tsc** ✅
  - Version: `^3.1.5`
  - TypeScript type checking for Vue files
- **@vitejs/plugin-vue** ✅
  - Version: `^6.0.2`
  - Vite plugin for Vue SFC support
- **vite-plugin-vue-devtools** ✅
  - Version: `^8.0.5`
  - Vue DevTools integration

## Architecture Patterns

### Backend
- **Layered Architecture**
  - `internal/api/` - HTTP handlers and routes
  - `internal/parser/` - Markdown parsing logic
  - `internal/executor/` - Code execution service
  - `internal/storage/` - Data persistence
  - `pkg/models/` - Shared data models

### Frontend
- **Component-Based Architecture**
  - Vue Single File Components (SFC)
  - Composition API with `<script setup>`
  - Composable functions for reusable logic
  - Pinia stores for global state

### Code Organization
```
frontend/src/
├── components/     # Vue components
├── composables/    # Reusable composition functions
├── stores/         # Pinia stores
├── services/        # API service layer
├── types/           # TypeScript type definitions
└── views/           # Route components
```

## API Communication

- **RESTful API Design**
  - JSON request/response format
  - Path-based routing (`/api/tutorials/:id`)
  - HTTP methods: GET, POST
  - CORS enabled for cross-origin requests

## Security Features

- **Code Execution Security**
  - Input validation
  - Dangerous import blocking
  - Network access restrictions
  - File I/O restrictions
  - Timeout protection
  - Isolated execution environments

## Performance Optimizations

- **Frontend**
  - Code splitting with Vite
  - Component lazy loading
  - Tutorial data caching (5-minute TTL)
  - Optimistic UI updates
  - LocalStorage persistence

- **Backend**
  - Efficient markdown parsing
  - In-memory tutorial cache
  - Isolated code execution
  - Resource limits

## Development Workflow

### Backend
```bash
go build ./...        # Build all packages
go run cmd/server/main.go  # Run server
```

### Frontend
```bash
npm install          # Install dependencies
npm run dev          # Development server
npm run build        # Production build
npm run type-check   # TypeScript checking
```

## Version Compatibility

- **Node.js**: `^20.19.0 || >=22.12.0`
- **Go**: `1.25+`
- **Vue**: `3.x`
- **TypeScript**: `5.9.x`

## Future Considerations

### Potential Upgrades
- [ ] Integrate Monaco Editor for advanced code editing (already installed)
- [ ] Utilize @vueuse/core composables for common utilities (already installed)
- [ ] Consider gorilla/mux for more advanced routing needs
- [ ] Add Docker containerization for code execution
- [ ] Implement WebSocket for real-time updates
- [ ] Add GraphQL API layer (optional)

### Alternative Options Available
- **CodeMirror**: Alternative to Monaco Editor
- **Prism.js**: Alternative to Shiki (lighter weight)
- **CSS Modules**: Alternative to Tailwind CSS
- **gorilla/mux**: Alternative to net/http for routing

## Summary

✅ **All specified technologies are implemented and configured**

The platform uses a modern, production-ready stack with:
- Go standard library for backend (minimal dependencies)
- Vue 3 Composition API for reactive frontend
- TypeScript for type safety
- Vite for fast development and builds
- Tailwind CSS for styling
- Shiki for syntax highlighting
- Pinia for state management

The stack is optimized for:
- Developer experience
- Performance
- Type safety
- Maintainability
- Security
