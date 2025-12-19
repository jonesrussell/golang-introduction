# Instructor Notes: Pass by Value vs Reference

## Teaching Techniques
- Show the "surprise" first (value doesn't change)
- Then show pointer solution
- Emphasize: "Go is ALWAYS pass by value"
- Explain what "value" means for different types

## Demo Flow
1. Show `doubleValue(num)` - doesn't change original
2. Explain: "Function got a COPY"
3. Show `doublePointer(&num)` - changes original
4. Explain: "Function got a COPY of the address"
5. Show struct example with value vs pointer receiver

## Key Emphasis
- **Go is always pass by value**: Always copies the argument
- **Passing pointer**: Copies the address (8 bytes), not the data
- **Both access same memory**: Caller and function share the data
- **Value semantics**: Prevent accidental mutation

## Common Questions
- "Why doesn't my function change the value?" - You passed by value!
- "Is Go pass by reference?" - No! Always pass by value (even pointers)
- "When do I use pointer?" - When you need to modify the original

## Engagement
- "Watch what happens - the original doesn't change!"
- "Now with a pointer - see the difference?"
- "This is why Go methods use pointer receivers for modification"

## Real-World Context
- This is why methods that modify use pointer receivers
- Value receivers create copies (can't modify original)
- Pointer receivers share the data (can modify)

## Transition
- "Let's see how this works with structs..."
