# Complete Practical Example: Employee Management System

**Duration:** 10-12 minutes

## Build Together

```go runnable
package main

import (
    "fmt"
    "time"
)

// ========================================
// Mixins - Reusable functionality
// ========================================

type Timestamps struct {
    CreatedAt time.Time
    UpdatedAt time.Time
}

func (t *Timestamps) Touch() {
    t.UpdatedAt = time.Now()
    if t.CreatedAt.IsZero() {
        t.CreatedAt = t.UpdatedAt
    }
}

type Identifiable struct {
    ID int
}

// ========================================
// Base Person type
// ========================================

type Person struct {
    Timestamps   // Mixin
    Identifiable // Mixin
    FirstName    string
    LastName     string
    Email        string
    Phone        string
}

func (p Person) FullName() string {
    return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}

func (p Person) ContactInfo() string {
    return fmt.Sprintf("%s (%s)", p.Email, p.Phone)
}

// ========================================
// Employee types with different roles
// ========================================

type Employee struct {
    Person           // Embedded - promotes all Person fields/methods
    EmployeeID       string
    Department       string
    HireDate         time.Time
    Salary           float64
}

func (e Employee) YearsOfService() int {
    return int(time.Since(e.HireDate).Hours() / 24 / 365)
}

func (e Employee) String() string {
    return fmt.Sprintf("Employee[%s]: %s - %s", 
        e.EmployeeID, e.FullName(), e.Department)
}

type Manager struct {
    Employee            // Embedded - Manager IS an Employee
    TeamSize     int
    DirectReports []string
}

func (m Manager) CanApprove(amount float64) bool {
    // Managers can approve up to $10k per team member
    return amount <= float64(m.TeamSize) * 10000
}

// Override String method
func (m Manager) String() string {
    return fmt.Sprintf("Manager[%s]: %s - %s (Team: %d)", 
        m.EmployeeID, m.FullName(), m.Department, m.TeamSize)
}

type Developer struct {
    Employee              // Embedded
    ProgrammingLanguages []string
    SeniorityLevel      string
}

func (d Developer) HasSkill(language string) bool {
    for _, lang := range d.ProgrammingLanguages {
        if lang == language {
            return true
        }
    }
    return false
}

func (d Developer) String() string {
    return fmt.Sprintf("Developer[%s]: %s - %s %s", 
        d.EmployeeID, d.FullName(), d.SeniorityLevel, d.Department)
}

// ========================================
// Contractor - different from Employee
// ========================================

type Contractor struct {
    Person              // Embedded Person (not Employee!)
    ContractID   string
    HourlyRate   float64
    EndDate      time.Time
}

func (c Contractor) IsActive() bool {
    return time.Now().Before(c.EndDate)
}

func (c Contractor) String() string {
    status := "Active"
    if !c.IsActive() {
        status = "Expired"
    }
    return fmt.Sprintf("Contractor[%s]: %s - %s", 
        c.ContractID, c.FullName(), status)
}

// ========================================
// Interfaces that work with embedded types
// ========================================

type Worker interface {
    FullName() string
    String() string
}

func main() {
    fmt.Println("=== Employee Management System ===\n")

    // Create a Manager
    manager := Manager{
        Employee: Employee{
            Person: Person{
                Identifiable: Identifiable{ID: 1},
                FirstName:    "Alice",
                LastName:     "Johnson",
                Email:        "alice@company.com",
                Phone:        "555-0101",
            },
            EmployeeID: "E001",
            Department: "Engineering",
            HireDate:   time.Now().AddDate(-3, 0, 0),
            Salary:     120000,
        },
        TeamSize:      5,
        DirectReports: []string{"E002", "E003", "E004"},
    }
    manager.Touch()

    // Create a Developer
    dev1 := Developer{
        Employee: Employee{
            Person: Person{
                Identifiable: Identifiable{ID: 2},
                FirstName:    "Bob",
                LastName:     "Smith",
                Email:        "bob@company.com",
                Phone:        "555-0102",
            },
            EmployeeID: "E002",
            Department: "Engineering",
            HireDate:   time.Now().AddDate(-2, 0, 0),
            Salary:     95000,
        },
        ProgrammingLanguages: []string{"Go", "Python", "JavaScript"},
        SeniorityLevel:      "Senior",
    }
    dev1.Touch()

    // Create a Contractor
    contractor := Contractor{
        Person: Person{
            Identifiable: Identifiable{ID: 4},
            FirstName:    "Diana",
            LastName:     "Williams",
            Email:        "diana@contractor.com",
            Phone:        "555-0104",
        },
        ContractID: "C001",
        HourlyRate: 85,
        EndDate:    time.Now().AddDate(0, 6, 0),
    }
    contractor.Touch()

    // Demonstrate promoted methods
    fmt.Println("=== Promoted Method Access ===")
    fmt.Printf("Manager full name: %s\n", manager.FullName())
    fmt.Printf("Manager contact: %s\n", manager.ContactInfo())
    fmt.Printf("Manager ID: %d\n", manager.ID)
    fmt.Printf("Years of service: %d\n", manager.YearsOfService())

    // Demonstrate method override
    fmt.Println("\n=== Method Override ===")
    fmt.Println(manager.String())
    fmt.Println(dev1.String())

    // Demonstrate type-specific functionality
    fmt.Println("\n=== Type-Specific Methods ===")
    if manager.CanApprove(40000) {
        fmt.Printf("%s can approve $40,000 expense\n", manager.FullName())
    }

    if dev1.HasSkill("Go") {
        fmt.Printf("%s knows Go!\n", dev1.FullName())
    }

    if contractor.IsActive() {
        fmt.Printf("%s contract is active until %s\n", 
            contractor.FullName(), 
            contractor.EndDate.Format("2006-01-02"))
    }

    // Using interface with embedded types
    fmt.Println("\n=== Interface Usage ===")
    workers := []Worker{manager, dev1, contractor}
    for _, w := range workers {
        fmt.Printf("  - %s\n", w.String())
    }
}
```

## Walk Through:
- Mixin structs for common functionality (Timestamps, Identifiable)
- Base Person struct with common identity fields
- Employee embeds Person (promotes all fields/methods)
- Manager and Developer embed Employee (multi-level embedding)
- Contractor embeds Person directly (different hierarchy)
- Interfaces work with embedded types
- Method promotion in action
- Method overriding
