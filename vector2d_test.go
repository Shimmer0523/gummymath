package gummymath_test

import (
	"math"
	"testing"

	"github.com/Shimmer0523/gummymath"
)

func TestNew(t *testing.T) {
	t.Run("creates a new vector", func(t *testing.T) {
		v := gummymath.NewVector2D(1, 2)
		if v.X != 1 {
			t.Errorf("expected X to be 1, got %f", v.X)
		}
		if v.Y != 2 {
			t.Errorf("expected Y to be 2, got %f", v.Y)
		}
	})
}

func TestAdd(t *testing.T) {
	t.Run("adds two vectors", func(t *testing.T) {
		v1 := gummymath.NewVector2D(1, 2)
		v2 := gummymath.NewVector2D(3, 4)
		v3 := v1.Add(v2)
		if v3.X != 4 {
			t.Errorf("expected X to be 4, got %f", v3.X)
		}
		if v3.Y != 6 {
			t.Errorf("expected Y to be 6, got %f", v3.Y)
		}
	})
}

func TestSub(t *testing.T) {
	t.Run("subtracts two vectors", func(t *testing.T) {
		v1 := gummymath.NewVector2D(5, 6)
		v2 := gummymath.NewVector2D(1, 2)
		v3 := v1.Sub(v2)
		if v3.X != 4 {
			t.Errorf("expected X to be 4, got %f", v3.X)
		}
		if v3.Y != 4 {
			t.Errorf("expected Y to be 4, got %f", v3.Y)
		}
	})
}

func TestScale(t *testing.T) {
	t.Run("scales a vector", func(t *testing.T) {
		v1 := gummymath.NewVector2D(2, 3)
		v2 := v1.Scale(2)
		if v2.X != 4 {
			t.Errorf("expected X to be 4, got %f", v2.X)
		}
		if v2.Y != 6 {
			t.Errorf("expected Y to be 6, got %f", v2.Y)
		}
	})
}

func TestDot(t *testing.T) {
	t.Run("calculates the dot product", func(t *testing.T) {
		v1 := gummymath.NewVector2D(1, 2)
		v2 := gummymath.NewVector2D(3, 4)
		dot := v1.Dot(v2)
		if dot != 11 {
			t.Errorf("expected dot product to be 11, got %f", dot)
		}
	})
}

func TestCross(t *testing.T) {
	t.Run("calculates the cross product", func(t *testing.T) {
		v1 := gummymath.NewVector2D(1, 2)
		v2 := gummymath.NewVector2D(3, 4)
		cross := v1.Cross(v2)
		if cross != -2 {
			t.Errorf("expected cross product to be -2, got %f", cross)
		}
	})
}

func TestRotate(t *testing.T) {
	t.Run("rotates a vector", func(t *testing.T) {
		v1 := gummymath.NewVector2D(1, 0)
		v2 := v1.Rotate(math.Pi / 2)
		if math.Abs(v2.X-0) > 1e-9 {
			t.Errorf("expected X to be 0, got %f", v2.X)
		}
		if math.Abs(v2.Y-1) > 1e-9 {
			t.Errorf("expected Y to be 1, got %f", v2.Y)
		}
	})
}

func TestLength(t *testing.T) {
	t.Run("calculates the length", func(t *testing.T) {
		v1 := gummymath.NewVector2D(3, 4)
		length := v1.Length()
		if length != 5 {
			t.Errorf("expected length to be 5, got %f", length)
		}
	})
}
