# Instructor Notes: Practical Example

## Teaching Techniques
- Build incrementally, don't paste all at once
- Explain interface design as you go
- Show how new notifiers can be added easily
- Demonstrate polymorphism in action

## Build Order
1. Define Notifier interface (small and focused)
2. Create EmailNotifier implementation
3. Create SMSNotifier implementation
4. Create SlackNotifier implementation
5. Create NotificationService (accepts []Notifier)
6. Show how easy it is to add new notifiers
7. Demonstrate all working together

## Live Commentary
- "First, let's define our interface - keep it small..."
- "EmailNotifier implements Notifier - notice it's implicit"
- "Now we can add SMS - same interface!"
- "NotificationService accepts any Notifier - polymorphism!"

## Things to Emphasize
- Small interface enables many implementations
- Easy to add new notifiers (just implement interface)
- Service doesn't know about concrete types
- This is the power of interfaces

## Engagement
- "What if we want to add a Discord notifier?"
- "Notice how the service doesn't change - just add new type"
- Challenge: "Add a PushNotification notifier"

## Variations to Mention
- Could add priority levels
- Could add retry logic
- Could add batching
- Could add filtering

## Common Mistakes to Watch For
- Making interface too large
- Returning interface instead of struct
- Not keeping interface small
