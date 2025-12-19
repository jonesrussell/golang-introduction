# Practical Example: Plugin System

**Duration:** 10-12 minutes

## Build Together

An extensible notification system demonstrating interface patterns.

```go runnable
package main

import (
    "fmt"
    "time"
)

// ========================================
// Core Interface - small and focused
// ========================================

// Notifier is the core interface for all notification types
type Notifier interface {
    Notify(message string) error
    Name() string
}

// ========================================
// Concrete Implementations
// ========================================

// EmailNotifier sends notifications via email
type EmailNotifier struct {
    From    string
    To      string
    SMTPHost string
}

func NewEmailNotifier(from, to, host string) *EmailNotifier {
    return &EmailNotifier{From: from, To: to, SMTPHost: host}
}

func (e *EmailNotifier) Notify(message string) error {
    fmt.Printf("[EMAIL] From: %s, To: %s\n", e.From, e.To)
    fmt.Printf("[EMAIL] Message: %s\n", message)
    return nil
}

func (e *EmailNotifier) Name() string {
    return "Email"
}

// SlackNotifier sends notifications to Slack
type SlackNotifier struct {
    WebhookURL string
    Channel    string
}

func NewSlackNotifier(webhookURL, channel string) *SlackNotifier {
    return &SlackNotifier{WebhookURL: webhookURL, Channel: channel}
}

func (s *SlackNotifier) Notify(message string) error {
    fmt.Printf("[SLACK] Channel: %s\n", s.Channel)
    fmt.Printf("[SLACK] Message: %s\n", message)
    return nil
}

func (s *SlackNotifier) Name() string {
    return "Slack"
}

// ConsoleNotifier for development/testing
type ConsoleNotifier struct{}

func NewConsoleNotifier() *ConsoleNotifier {
    return &ConsoleNotifier{}
}

func (c *ConsoleNotifier) Notify(message string) error {
    fmt.Printf("[CONSOLE] %s: %s\n", time.Now().Format("15:04:05"), message)
    return nil
}

func (c *ConsoleNotifier) Name() string {
    return "Console"
}

// ========================================
// Optional Interface - for additional capabilities
// ========================================

// Validator can validate before sending
type Validator interface {
    Validate() error
}

// EmailNotifier implements Validator
func (e *EmailNotifier) Validate() error {
    if e.To == "" {
        return fmt.Errorf("email recipient cannot be empty")
    }
    return nil
}

// ========================================
// Notification Manager
// ========================================

type NotificationManager struct {
    notifiers []Notifier
}

func NewNotificationManager() *NotificationManager {
    return &NotificationManager{
        notifiers: make([]Notifier, 0),
    }
}

// Register adds a notifier (accepts interface)
func (nm *NotificationManager) Register(n Notifier) {
    // Check if notifier implements Validator
    if v, ok := n.(Validator); ok {
        if err := v.Validate(); err != nil {
            fmt.Printf("Warning: %s notifier validation failed: %v\n", n.Name(), err)
            return
        }
    }
    nm.notifiers = append(nm.notifiers, n)
    fmt.Printf("Registered: %s notifier\n", n.Name())
}

// NotifyAll sends message to all registered notifiers
func (nm *NotificationManager) NotifyAll(message string) error {
    fmt.Printf("\n--- Sending notification to %d channels ---\n", len(nm.notifiers))

    var lastErr error
    for _, n := range nm.notifiers {
        if err := n.Notify(message); err != nil {
            fmt.Printf("Error with %s: %v\n", n.Name(), err)
            lastErr = err
        }
    }
    return lastErr
}

// ========================================
// Decorator Pattern with Interfaces
// ========================================

// LoggingNotifier wraps any Notifier with logging
type LoggingNotifier struct {
    wrapped Notifier
}

func WithLogging(n Notifier) *LoggingNotifier {
    return &LoggingNotifier{wrapped: n}
}

func (l *LoggingNotifier) Notify(message string) error {
    start := time.Now()
    fmt.Printf("[LOG] Starting %s notification...\n", l.wrapped.Name())

    err := l.wrapped.Notify(message)

    elapsed := time.Since(start)
    if err != nil {
        fmt.Printf("[LOG] %s failed after %v: %v\n", l.wrapped.Name(), elapsed, err)
    } else {
        fmt.Printf("[LOG] %s completed in %v\n", l.wrapped.Name(), elapsed)
    }
    return err
}

func (l *LoggingNotifier) Name() string {
    return l.wrapped.Name() + " (logged)"
}

// ========================================
// Main - Demonstration
// ========================================

func main() {
    fmt.Println("=== Notification Plugin System ===\n")

    // Create manager
    manager := NewNotificationManager()

    // Register various notifiers
    manager.Register(NewConsoleNotifier())
    manager.Register(NewEmailNotifier("system@company.com", "admin@company.com", "smtp.company.com"))
    manager.Register(NewSlackNotifier("https://hooks.slack.com/...", "#alerts"))

    // This one will fail validation
    manager.Register(NewEmailNotifier("system@company.com", "", "smtp.company.com"))

    // Send to all
    manager.NotifyAll("Server CPU usage exceeded 90%!")

    fmt.Println("\n=== Using Decorators ===\n")

    // Wrap with logging
    loggedEmail := WithLogging(NewEmailNotifier("alerts@company.com", "ops@company.com", "smtp.company.com"))
    loggedEmail.Notify("Database connection lost")

    fmt.Println("\n=== Type Assertions ===\n")

    // Get specific notifier and check capabilities
    var notifier Notifier = NewEmailNotifier("test@test.com", "user@test.com", "smtp")
    
    // Check if notifier implements optional Validator interface
    if validator, ok := notifier.(Validator); ok {
        if err := validator.Validate(); err != nil {
            fmt.Println("Validation error:", err)
        } else {
            fmt.Println("Notifier is valid")
        }
    }
}
```

## Walk Through:
- Core `Notifier` interface (small, focused)
- Multiple implementations (Email, Slack, Console)
- Optional `Validator` interface
- Type assertions to check capabilities
- Decorator pattern (logging wrapper)
- Manager works with any Notifier
- Easy to add new notification types
