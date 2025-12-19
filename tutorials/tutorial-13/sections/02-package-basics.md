# Package Basics

**Duration:** 6-7 minutes

## Topics:
- Package declaration
- Imports
- Exported vs unexported
- Package initialization

## Code Examples

```go
// math/calculator.go
package math  // Package declaration

import (
    "fmt"       // Standard library
    "strings"   // Multiple imports

    "github.com/user/project/internal/util"  // Project import
)

// Exported (capital letter) - accessible from other packages
func Add(a, b int) int {
    return a + b
}

// Unexported (lowercase) - only this package
func validateInput(n int) bool {
    return n >= 0
}

// Exported constant
const MaxValue = 1000

// Unexported constant
const defaultPrecision = 2

// Exported type
type Calculator struct {
    precision int
}

// Unexported type
type operation func(int, int) int

// Package-level variables
var (
    ErrOverflow = errors.New("overflow")   // Exported
    cache       = make(map[string]int)     // Unexported
)

// init() runs when package is imported
func init() {
    fmt.Println("math package initialized")
    // Setup, register, etc.
}
```

## Import Rules

```go
// Standard format
import "fmt"

// Grouped imports
import (
    "fmt"
    "os"
)

// Alias
import (
    f "fmt"  // Use as f.Println()
)

// Blank import (side effects only)
import (
    _ "github.com/lib/pq"  // Registers postgres driver
)

// Dot import (avoid - pollutes namespace)
import (
    . "fmt"  // Use Println() instead of fmt.Println()
)
```
