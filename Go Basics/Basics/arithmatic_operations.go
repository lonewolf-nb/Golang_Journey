package basics

// Importing required packages
import (
	"fmt"  // For printing output to the console
	"math" // For mathematical constants and functions
)

func main() {
	// Declaring and initializing two integer variables
	var a, b int = 10, 2

	// Declaring a variable to store results of operations
	var result int

	// Performing addition
	result = a + b
	fmt.Println("Addition:", result)

	// Performing subtraction
	result = a - b
	fmt.Println("Subtraction:", result)

	// Performing multiplication
	result = a * b
	fmt.Println("Multiplication:", result)

	// Performing integer division (result will be an integer)
	result = a / b
	fmt.Println("Division:", result)

	// Finding remainder using modulus operator
	result = a % b
	fmt.Println("Remainder:", result)

	// Declaring a constant for an approximate value of π (pi)
	const p float64 = 22.0 / 7.0
	fmt.Println("Approximate value of π:", p)

	// Demonstrating overflow with signed integers
	var maxInt int64 = 9223372036854775807 // Maximum value for int64
	fmt.Println("Max int64 value:", maxInt)

	// Adding 1 causes overflow, wraps around to negative value
	maxInt = maxInt + 1
	fmt.Println("After overflow (int64):", maxInt)

	// Demonstrating overflow with unsigned integers
	var uMaxInt uint64 = 18446744073709551615 // Maximum value for uint64
	fmt.Println("Max uint64 value:", uMaxInt)

	// Adding 1 causes overflow, wraps around to 0
	uMaxInt = uMaxInt + 1
	fmt.Println("After overflow (uint64):", uMaxInt)

	// Demonstrating underflow with floating-point numbers
	var smallFloat float64 = 1.0e-323 // Very small number close to zero
	fmt.Println("Small float value:", smallFloat)

	// Dividing by a large number causes underflow (value becomes even smaller)
	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println("After underflow (float64):", smallFloat)
}


/*

Sure Navneet! Here's a slightly expanded version of your notes—still concise, but with a bit more depth for clarity and future reference:

---

### 📝 Go Notes: `package main` vs Custom Package Names

#### 🔹 `package main`
- Required for **standalone executable programs**.
- Must include a `main()` function—this is the entry point.
- Used when running files directly with:
  ```bash
  go run filename.go
  ```
- Without `package main`, Go won’t know where to start execution.

#### 🔹 Custom Package Names (e.g., `package arithmatic_operations`)
- Used for **modular code** or libraries.
- Cannot be run directly with `go run`—will throw:
  ```
  go: cannot run non-main package
  ```
- Intended to be **imported** into another file that uses `package main`.

#### 🔹 Example Use Case
**Library file:**
```go
// arithmatic_operations.go
package arithmatic_operations

func Add(a, b int) int {
    return a + b
}
```

**Main file:**
```go
// main.go
package main

import (
    "fmt"
    "yourmodule/arithmatic_operations"
)

func main() {
    fmt.Println(arithmatic_operations.Add(5, 3))
}
```

#### 🔹 Summary
| Use Case                     | Package Name         | Can Run Directly? |
|-----------------------------|----------------------|-------------------|
| Executable program           | `main`               | ✅ Yes            |
| Reusable module/library      | Custom (e.g. `arithmatic_operations`) | ❌ No             |

---

Let me know if you want to modularize your current project or set up a multi-file structure with reusable logic. I can help scaffold it out for you.
 */