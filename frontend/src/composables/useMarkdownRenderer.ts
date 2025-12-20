/**
 * Composable for rendering markdown content to HTML.
 * Consolidates markdown rendering logic that was duplicated across components.
 */

// CSS classes for consistent styling across components
const linkClasses = 'text-go-blue hover:text-go-blue-dark underline underline-offset-2';
const inlineCodeClasses = 'px-1.5 py-0.5 bg-neutral-100 dark:bg-neutral-800 text-go-blue rounded text-sm font-mono';
const boldClasses = 'font-semibold';

/**
 * Escapes HTML entities to prevent XSS.
 */
export function escapeHtml(text: string): string {
  const div = document.createElement('div');
  div.textContent = text;
  return div.innerHTML;
}

/**
 * Converts basic markdown syntax to HTML.
 * Handles: links, inline code, bold text.
 */
export function renderInlineMarkdown(text: string): string {
  let html = text;

  // Convert markdown links [text](url) to HTML anchors
  html = html.replace(
    /\[([^\]]+)\]\(([^)]+)\)/g,
    `<a href="$2" target="_blank" rel="noopener noreferrer" class="${linkClasses}">$1</a>`
  );

  // Convert inline code `code` to styled spans
  html = html.replace(
    /`([^`]+)`/g,
    `<code class="${inlineCodeClasses}">$1</code>`
  );

  // Convert bold **text** to strong
  html = html.replace(
    /\*\*(.+?)\*\*/g,
    `<strong class="${boldClasses}">$1</strong>`
  );

  // Convert italic *text* to em (after bold to avoid conflicts)
  html = html.replace(
    /\*(.+?)\*/g,
    '<em>$1</em>'
  );

  return html;
}

/**
 * Composable that provides markdown rendering functions.
 */
export function useMarkdownRenderer() {
  /**
   * Renders basic inline markdown (links, code, bold).
   */
  const renderMarkdown = (text: string): string => {
    return renderInlineMarkdown(text);
  };

  /**
   * Renders markdown with code block support for instructor notes.
   * Uses amber color scheme for instructor content.
   */
  const renderInstructorMarkdown = (content: string): string => {
    let html = content;

    // Convert headers
    html = html.replace(/^### (.+)$/gm, '<h3 class="text-lg font-semibold text-amber-900 dark:text-amber-100 mt-6 mb-3">$1</h3>');
    html = html.replace(/^## (.+)$/gm, '<h2 class="text-xl font-bold text-amber-900 dark:text-amber-100 mt-6 mb-4">$1</h2>');
    html = html.replace(/^# (.+)$/gm, '<h1 class="text-2xl font-bold text-amber-900 dark:text-amber-100 mb-4">$1</h1>');

    // Convert bold and italic
    html = html.replace(/\*\*(.+?)\*\*/g, '<strong class="font-semibold text-amber-900 dark:text-amber-100">$1</strong>');
    html = html.replace(/\*(.+?)\*/g, '<em>$1</em>');

    // Convert inline code
    html = html.replace(/`([^`]+)`/g, '<code class="px-1.5 py-0.5 bg-amber-100 dark:bg-amber-900/50 text-amber-800 dark:text-amber-200 rounded text-sm font-mono">$1</code>');

    // Convert code blocks
    html = html.replace(/```(\w*)\n([\s\S]*?)```/g, (_, _lang, code) => {
      return `<pre class="p-4 bg-neutral-900 rounded-lg overflow-x-auto my-4"><code class="text-sm font-mono text-neutral-100">${escapeHtml(code.trim())}</code></pre>`;
    });

    // Convert bullet lists
    html = html.replace(/^- (.+)$/gm, '<li class="text-amber-800 dark:text-amber-200 ml-4">$1</li>');
    html = html.replace(/(<li.*?<\/li>\n?)+/g, '<ul class="list-disc space-y-1 my-3">$&</ul>');

    // Convert paragraphs (lines not already converted)
    html = html.split('\n\n').map(para => {
      if (para.trim() && !para.startsWith('<')) {
        return `<p class="text-amber-800 dark:text-amber-200 leading-relaxed my-3">${para}</p>`;
      }
      return para;
    }).join('\n');

    return html;
  };

  /**
   * Renders section content markdown, excluding code blocks (which are rendered separately).
   * Handles headings, paragraphs, links, inline code, bold, italic, and lists.
   */
  const renderSectionContent = (content: string): string => {
    let html = content;

    // Remove code blocks (they're rendered separately via CodeRunner components)
    html = html.replace(/```[\s\S]*?```/g, '');

    // Convert headers (## becomes h2, ### becomes h3, etc.)
    html = html.replace(/^#### (.+)$/gm, '<h4 class="text-base font-semibold text-neutral-900 dark:text-neutral-100 mt-6 mb-3">$1</h4>');
    html = html.replace(/^### (.+)$/gm, '<h3 class="text-lg font-semibold text-neutral-900 dark:text-neutral-100 mt-6 mb-3">$1</h3>');
    html = html.replace(/^## (.+)$/gm, '<h2 class="text-xl font-bold text-neutral-900 dark:text-neutral-100 mt-8 mb-4">$1</h2>');
    html = html.replace(/^# (.+)$/gm, '<h1 class="text-2xl font-bold text-neutral-900 dark:text-neutral-100 mt-8 mb-4">$1</h1>');

    // Convert markdown links [text](url) to HTML anchors
    html = html.replace(
      /\[([^\]]+)\]\(([^)]+)\)/g,
      `<a href="$2" target="_blank" rel="noopener noreferrer" class="${linkClasses}">$1</a>`
    );

    // Convert inline code `code` to styled spans
    html = html.replace(
      /`([^`]+)`/g,
      `<code class="${inlineCodeClasses}">$1</code>`
    );

    // Convert bold **text** to strong (before italic to avoid conflicts)
    html = html.replace(
      /\*\*(.+?)\*\*/g,
      `<strong class="${boldClasses}">$1</strong>`
    );

    // Convert italic *text* to em (after bold to avoid conflicts)
    html = html.replace(
      /(?<!\*)\*(.+?)\*(?!\*)/g,
      '<em>$1</em>'
    );

    // Convert bullet lists (- item)
    html = html.replace(/^- (.+)$/gm, '<li class="text-neutral-900 dark:text-neutral-100 ml-4">$1</li>');
    html = html.replace(/(<li.*?<\/li>\n?)+/g, '<ul class="list-disc space-y-2 my-4 pl-6">$&</ul>');

    // Convert paragraphs (non-empty lines not already converted)
    const lines = html.split('\n');
    const processedLines: string[] = [];
    let currentParagraph: string[] = [];

    for (const line of lines) {
      const trimmed = line.trim();
      if (!trimmed) {
        // Empty line - flush current paragraph
        if (currentParagraph.length > 0) {
          processedLines.push(`<p class="text-neutral-900 dark:text-neutral-100 leading-relaxed my-4">${currentParagraph.join(' ')}</p>`);
          currentParagraph = [];
        }
        processedLines.push('');
      } else if (trimmed.startsWith('<')) {
        // Already processed (heading, list, etc.)
        if (currentParagraph.length > 0) {
          processedLines.push(`<p class="text-neutral-900 dark:text-neutral-100 leading-relaxed my-4">${currentParagraph.join(' ')}</p>`);
          currentParagraph = [];
        }
        processedLines.push(line);
      } else {
        // Regular text line
        currentParagraph.push(trimmed);
      }
    }

    // Flush remaining paragraph
    if (currentParagraph.length > 0) {
      processedLines.push(`<p class="text-neutral-900 dark:text-neutral-100 leading-relaxed my-4">${currentParagraph.join(' ')}</p>`);
    }

    html = processedLines.join('\n');

    // Clean up multiple empty lines
    html = html.replace(/\n{3,}/g, '\n\n');

    return html.trim();
  };

  return {
    renderMarkdown,
    renderInstructorMarkdown,
    renderSectionContent,
    escapeHtml,
  };
}
