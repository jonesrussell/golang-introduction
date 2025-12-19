# Instructor Notes: Practical Example

## Teaching Techniques
- Build incrementally, don't paste all at once
- Explain each step as you type
- Run after each addition to show progress
- Show validation in action

## Build Order
1. Define User struct with fields
2. Create NewUser constructor with validation
3. Add methods: Activate, Deactivate, UpdateEmail
4. Add String() method for formatting
5. Create UserManager struct
6. Add methods: AddUser, FindUser, ListUsers
7. Show it all working together

## Live Commentary
- "First, let's define our User struct..."
- "Notice the constructor validates the email..."
- "Methods give us behavior attached to data..."
- "UserManager shows how structs work together..."

## Things to Emphasize
- Constructor pattern for validation
- Pointer receivers for methods that modify
- String() method for readable output
- Struct composition (UserManager contains Users)

## Engagement
- "What happens if we try to create a user with an invalid email?"
- "Let's add a method to check if a user is active"
- Challenge: "Add a method to get all active users"

## Variations to Mention
- Could add more validation (password strength, etc.)
- Could add database persistence
- Could add user roles/permissions

## Common Mistakes to Watch For
- Forgetting pointer receiver for methods that modify
- Not validating in constructor
- Nil pointer dereference
