# Comparison to Other Languages & Next Steps

**Duration:** 3-4 minutes

## Comparison to Other Languages

```go
// Go - Composition via Embedding
type Animal struct {
    Name string
}

func (a Animal) Speak() {
    fmt.Println("Some sound")
}

type Dog struct {
    Animal  // Composition, not inheritance!
    Breed string
}

func (d Dog) Speak() {  // Override
    fmt.Println("Woof!")
}

// Python - Class Inheritance
/*
class Animal:
    def __init__(self, name):
        self.name = name
    
    def speak(self):
        print("Some sound")

class Dog(Animal):  # Inheritance
    def __init__(self, name, breed):
        super().__init__(name)  # Call parent constructor
        self.breed = breed
    
    def speak(self):  # Override
        print("Woof!")
*/

// Java - Class Inheritance
/*
class Animal {
    String name;
    void speak() {
        System.out.println("Some sound");
    }
}

class Dog extends Animal {  // Inheritance
    String breed;
    
    @Override
    void speak() {
        System.out.println("Woof!");
    }
}
*/
```

## Key Differences:
- Go: [No inheritance keyword](https://go.dev/doc/faq#inheritance), no `extends`, no `super`
- Go: [Composition is explicit](https://go.dev/ref/spec#Struct_types) and visible
- Go: No virtual methods - method resolution is simple
- Go: Can't accidentally break parent class
- Go: [Interface satisfaction](https://go.dev/ref/spec#Interface_types) is implicit
- Go: Simpler mental model - just nested structs

## Recap What Was Covered:
- Basic [composition](https://go.dev/ref/spec#Struct_types) with named fields
- [Struct embedding](https://go.dev/ref/spec#Struct_types) (anonymous fields)
- [Field and method promotion](https://go.dev/ref/spec#Selectors)
- Multiple embedding and conflicts
- Embedding with [interfaces](https://go.dev/ref/spec#Interface_types)
- Practical patterns (mixins, decorators, base entities)
- When to use embedding vs composition
- Common pitfalls to avoid

## Preview Next Topics:
- [Interfaces](https://go.dev/ref/spec#Interface_types) in depth
- [Polymorphism](https://go.dev/doc/faq#polymorphism) in Go
- Type assertions and type switches
- [Interface composition](https://go.dev/ref/spec#Interface_types)
- [Error handling](https://go.dev/doc/effective_go#errors) patterns

## Practice Suggestions:
1. **Easy:** Create a Shape hierarchy (Shape → Rectangle, Circle) using embedding
2. **Medium:** Build a notification system with different notifier types
3. **Challenge:** Create a plugin system where plugins embed common functionality
4. **Advanced:** Implement a middleware chain using embedding

## Resources:
- [Effective Go on Embedding](https://go.dev/doc/effective_go#embedding): go.dev/doc/effective_go#embedding
- [Go Tour - Struct Embedding](https://go.dev/tour/moretypes/10): go.dev/tour/moretypes/10
- [Go FAQ - Why no inheritance?](https://go.dev/doc/faq#inheritance): go.dev/doc/faq#inheritance
