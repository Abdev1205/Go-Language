## Statically Typed Language in Go 🧐

In a statically typed language like Go, variable types are explicitly declared at compile-time. This means that the data type of a variable is known and checked by the compiler before the program runs.

### Example in Go

```go
package main

import "fmt"

func main() {
    // Statically typed variables
    var age int = 25
    var name string = "John"

    // Attempting to assign a different type results in a compile-time error
    // var invalidAssignment string = 42  // Uncommenting this line will cause a compile-time error

    // Print values
    fmt.Printf("Name: %s, Age: %d\n", name, age)
}
```

In this example:

- `age` is explicitly declared as an `int` variable.
- `name` is explicitly declared as a `string` variable.
- If you uncomment the line with `invalidAssignment`, a compile-time error will occur because it attempts to assign an `int` to a `string`.

### Advantages of Statically Typed Languages

- **Early Error Detection**: Type errors are caught at compile-time, reducing the likelihood of runtime errors.
- **Improved Performance**: The compiler can optimize the code more effectively when it knows the data types in advance.
- **Better Code Quality**: Type information enhances code readability and provides clear documentation.

Using a statically typed language like Go ensures a more robust and reliable codebase. 👩‍💻✨
