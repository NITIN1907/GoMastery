Go (Golang) is an open-source, statically typed, compiled programming language designed by Google.

It’s known for simplicity, speed, concurrency support, and strong performance.


🚀 Key Features of Go
-----------------------------------------------------------------------------------
| Feature                  |                 Description                          |
|--------------------------|----------------------------------------------------- |
| **Compiled Language**    | Compiles to native machine code, very fast execution |
| **Static Typing**        | Catches errors at compile time                       |
| **Garbage Collected**    | Automatic memory management                          |
| **Built-in Concurrency** | Uses goroutines and channels for lightweight concurrency |
| **Fast Compilation**     | Much faster build times compared to C/C++            |
| **Simple Syntax**        | Easy to learn and read                               |
| **Standard Library**     | Rich and consistent standard library                 |
| **Cross-Platform**       | Easily compiles to Windows, macOS, Linux, etc.       |


🔄 How Go is Different from Other Languages
| Language  | Comparison with Go |
|-----------|--------------------|
| **C/C++** | Go has garbage collection, simpler syntax, built-in concurrency, and no manual memory management |
| **Java**  | No JVM required; faster startup; no class-based OOP, uses composition |
| **Python**| Go is compiled and much faster; strongly typed; built-in concurrency |
| **Rust**  | Rust offers better memory safety but is more complex; Go is simpler and faster to write for many apps |

📦 package main

Every Go file starts with a package declaration.

main is a special package that tells Go this is the entry point to build an executable program.

Only package main can contain a main() function to run directly.

📥 import "fmt"

import is used to bring in packages.

fmt stands for format – it’s a standard library package used for formatted I/O operations like:

fmt.Println() – print with a newline

fmt.Printf() – formatted output

fmt.Scan() – read user input


✅ Rules for Writing Functions

|                      Rules                            | 
|-------------------------------------------------------|
| Use `func` keyword                                    |
| Function names must be unique within the same package |
| Capitalized function names are **exported** (public)  |
| Lowercase function names are **unexported** (private) |
| Parameter types must be declared                      |
| Return type must be declared (or omitted if void)     |
| Can return multiple values                            |
| Supports named return values                          |

🧑‍💻 Function Naming Rules
|              RULE                        |         Example                    |
|------------------------------------------|------------------------------------|
| Must start with a letter                 | `func add()` ✅                    |
| Can contain letters, digits, underscores | `func addTwoNumbers()` ✅          |
| Cannot use Go keywords                   | `func if()` ❌                     |
| Case matters                             | `Add()` and `add()` are different  |


## 📘 GoLang Variable Declaration Rules

### 1. **Basic Declaration Syntax**
```go
var name type
```
Example:
```go
var age int
```

### 2. **Declaration with Initialization**
```go
var name type = value
```
Example:
```go
var age int = 25
```

### 3. **Type Inference (type is optional if initialized)**
```go
var name = value
```
Example:
```go
var city = "New York" // Go infers the type as string
```

### 4. **Short Variable Declaration (`:=`)**
- Used **inside functions only**
```go
name := value
```
Example:
```go
message := "Hello, Go!"
```

### 5. **Multiple Variable Declaration**
```go
var a, b, c int
```
Or with different types:
```go
var (
    name string
    age  int
    isGoAwesome bool
)
```

### 6. **Zero Values**
- If no value is assigned, Go sets a **zero value**:
  - `0` for numbers
  - `""` for strings
  - `false` for booleans
  - `nil` for pointers, slices, maps, interfaces, etc.

Example:
```go
var n int        // 0
var s string     // ""
var b bool       // false
```

### 7. **Variable Names**
- Must begin with a letter or underscore (_)
- Can include letters, digits, and underscores
- Are case-sensitive (`Name` ≠ `name`)
- Use **camelCase** for local variables and **PascalCase** for exported variables

### 8. **Exported Variables**
- If a variable starts with a **capital letter**, it is **exported** (accessible from other packages).


