package vector2

import "math"

// Vector2D represents a 2D vector.
type Vector2D struct {
	X float64
	Y float64
}

// New creates a new 2D vector.
func New(x, y float64) *Vector2D {
	return &Vector2D{X: x, Y: y}
}

// Add adds two vectors and returns a new vector.
func (v *Vector2D) Add(other *Vector2D) *Vector2D {
	return &Vector2D{X: v.X + other.X, Y: v.Y + other.Y}
}

// Sub subtracts two vectors and returns a new vector.
func (v *Vector2D) Sub(other *Vector2D) *Vector2D {
	return &Vector2D{X: v.X - other.X, Y: v.Y - other.Y}
}

// Scale scales a vector by a scalar and returns a new vector.
func (v *Vector2D) Scale(s float64) *Vector2D {
	return &Vector2D{X: v.X * s, Y: v.Y * s}
}

// Dot calculates the dot product of two vectors.
func (v *Vector2D) Dot(other *Vector2D) float64 {
	return v.X*other.X + v.Y*other.Y
}

// Cross calculates the cross product of two vectors.
// The cross product of 2D vectors is a scalar.
func (v *Vector2D) Cross(other *Vector2D) float64 {
	return v.X*other.Y - v.Y*other.X
}

// Rotate rotates a vector by a given angle in radians.
func (v *Vector2D) Rotate(angle float64) *Vector2D {
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	return &Vector2D{
		X: v.X*cos - v.Y*sin,
		Y: v.X*sin + v.Y*cos,
	}
}

// Length returns the length (magnitude) of the vector.
func (v *Vector2D) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
