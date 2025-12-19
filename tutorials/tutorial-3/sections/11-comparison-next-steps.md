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
- Go: No inheritance keyword, no `extends`, no `super`
- Go: Composition is explicit and visible
- Go: No virtual methods - method resolution is simple
- Go: Can't accidentally break parent class
- Go: Interface satisfaction is implicit
- Go: Simpler mental model - just nested structs

## Recap What Was Covered:
- Basic composition with named fields
- Struct embedding (anonymous fields)
- Field and method promotion
- Multiple embedding and conflicts
- Embedding with interfaces
- Practical patterns (mixins, decorators, base entities)
- When to use embedding vs composition
- Common pitfalls to avoid

## Preview Next Topics:
- Interfaces in depth
- Polymorphism in Go
- Type assertions and type switches
- Interface composition
- Error handling patterns

## Practice Suggestions:
1. **Easy:** Create a Shape hierarchy (Shape → Rectangle, Circle) using embedding
2. **Medium:** Build a notification system with different notifier types
3. **Challenge:** Create a plugin system where plugins embed common functionality
4. **Advanced:** Implement a middleware chain using embedding

## Resources:
- Effective Go on Embedding: golang.org/doc/effective_go#embedding
- Go blog on embedding: blog.golang.org/json-and-go
