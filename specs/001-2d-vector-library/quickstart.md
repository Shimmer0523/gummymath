# Quickstart

This guide provides a quick overview of how to use the 2D vector library.

## Installation

```bash
go get github.com/Shimmer0523/gmath
```

## Usage

```go
package main

import (
	"fmt"
	"math"

	"github.com/Shimmer0523/gummymath/src/vector2"
)

func main() {
	// Create two vectors
	v1 := vector2.New(1, 2)
	v2 := vector2.New(3, 4)

	// Add vectors
	v3 := v1.Add(v2)
	fmt.Printf("v1 + v2 = %v\n", v3) // Expected: {4, 6}

	// Subtract vectors
	v4 := v1.Sub(v2)
	fmt.Printf("v1 - v2 = %v\n", v4) // Expected: {-2, -2}

	// Scale a vector
	v5 := v1.Scale(2)
	fmt.Printf("v1 * 2 = %v\n", v5) // Expected: {2, 4}

	// Dot product
	dot := v1.Dot(v2)
	fmt.Printf("v1 . v2 = %v\n", dot) // Expected: 11

	// Cross product
	cross := v1.Cross(v2)
	fmt.Printf("v1 x v2 = %v\n", cross) // Expected: -2

	// Rotate a vector
	v6 := vector2.New(1, 0)
	v7 := v6.Rotate(math.Pi / 2) // Rotate 90 degrees
	fmt.Printf("v6 rotated by 90 degrees = %v\n", v7) // Expected: {0, 1}
}
```

```
