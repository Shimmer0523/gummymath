# Quickstart: Get Vector Length

This guide shows how to use the `Length()` function to get the magnitude of a `Vector2D`.

```go
package main

import (
	"fmt"
	"github.com/shimmer0523/gummymath/vector2"
)

func main() {
	v := vector2.Vector2D{X: 3, Y: 4}
	length := v.Length()
	fmt.Println(length) // Output: 5
}
```