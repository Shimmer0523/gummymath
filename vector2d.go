package gummymath

import "math"

// Vector2D represents a 2D vector.
type Vector2D struct {
	X float64
	Y float64
}

// New creates a new 2D vector.
func NewVector2D(x, y float64) Vector2D {
	return Vector2D{X: x, Y: y}
}

// Length returns the length (magnitude) of the vector.
func (v Vector2D) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vector2D) Normalized() Vector2D {
	length := v.Length()
	return Vector2D{X: v.X / length, Y: v.Y / length}
}

// Add adds two vectors and returns a new vector.
func (v Vector2D) Added(other Vector2D) Vector2D {
	return Vector2D{X: v.X + other.X, Y: v.Y + other.Y}
}

// Sub subtracts two vectors and returns a new vector.
func (v Vector2D) Subed(other Vector2D) Vector2D {
	return Vector2D{X: v.X - other.X, Y: v.Y - other.Y}
}

// Scale scales a vector by a scalar and returns a new vector.
func (v Vector2D) Scaled(s float64) Vector2D {
	return Vector2D{X: v.X * s, Y: v.Y * s}
}

// Dot calculates the dot product of two vectors.
func (v Vector2D) Dotted(other Vector2D) float64 {
	return v.X*other.X + v.Y*other.Y
}

// Cross calculates the cross product of two vectors.
// The cross product of 2D vectors is a scalar.
func (v Vector2D) Crossed(other Vector2D) float64 {
	return v.X*other.Y - v.Y*other.X
}

// Rotate rotates a vector by a given angle in radians.
func (v Vector2D) Rotated(angle float64) Vector2D {
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	return Vector2D{
		X: v.X*cos - v.Y*sin,
		Y: v.X*sin + v.Y*cos,
	}
}

// Add adds another vector to the current vector.
func (v *Vector2D) Add(other Vector2D) {
	v.X += other.X
	v.Y += other.Y
}

// Sub subtracts another vector from the current vector.
func (v *Vector2D) Sub(other Vector2D) {
	v.X -= other.X
	v.Y -= other.Y
}

// Scale scales the current vector by a scalar.
func (v *Vector2D) Scale(s float64) {
	v.X *= s
	v.Y *= s
}

// Normalize normalizes the current vector.
func (v *Vector2D) Normalize() {
	length := v.Length()
	v.X /= length
	v.Y /= length
}

// Rotate rotates the current vector by a given angle in radians.
func (v *Vector2D) Rotate(angle float64) {
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	x := v.X*cos - v.Y*sin
	y := v.X*sin + v.Y*cos
	v.X = x
	v.Y = y
}
