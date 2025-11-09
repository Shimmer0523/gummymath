package gummymath_test

import (
	"testing"

	"github.com/Shimmer0523/gummymath"
)

func TestNewRect2D(t *testing.T) {
	t.Run("creates a new rectangle", func(t *testing.T) {
		rect := gummymath.NewRect2D(10, 20, 100, 50)
		if rect.Position.X != 10 {
			t.Errorf("expected X to be 10, got %f", rect.Position.X)
		}
		if rect.Position.Y != 20 {
			t.Errorf("expected Y to be 20, got %f", rect.Position.Y)
		}
		if rect.Width != 100 {
			t.Errorf("expected Width to be 100, got %f", rect.Width)
		}
		if rect.Height != 50 {
			t.Errorf("expected Height to be 50, got %f", rect.Height)
		}
	})
}

func TestCorners(t *testing.T) {
	t.Run("returns the four corners", func(t *testing.T) {
		rect := gummymath.NewRect2D(10, 20, 100, 50)
		corners := rect.Corners()

		// Top-Left
		if corners[0].X != 10 || corners[0].Y != 20 {
			t.Errorf("expected Top-Left to be (10, 20), got (%f, %f)", corners[0].X, corners[0].Y)
		}
		// Top-Right
		if corners[1].X != 110 || corners[1].Y != 20 {
			t.Errorf("expected Top-Right to be (110, 20), got (%f, %f)", corners[1].X, corners[1].Y)
		}
		// Bottom-Right
		if corners[2].X != 110 || corners[2].Y != 70 {
			t.Errorf("expected Bottom-Right to be (110, 70), got (%f, %f)", corners[2].X, corners[2].Y)
		}
		// Bottom-Left
		if corners[3].X != 10 || corners[3].Y != 70 {
			t.Errorf("expected Bottom-Left to be (10, 70), got (%f, %f)", corners[3].X, corners[3].Y)
		}
	})
}

func TestRect2DModification(t *testing.T) {
	t.Run("modifies a rectangle", func(t *testing.T) {
		rect := gummymath.NewRect2D(10, 20, 100, 50)

		rect.Position.X = 30
		rect.Position.Y = 40
		rect.Width = 150
		rect.Height = 75

		if rect.Position.X != 30 {
			t.Errorf("expected X to be 30, got %f", rect.Position.X)
		}
		if rect.Position.Y != 40 {
			t.Errorf("expected Y to be 40, got %f", rect.Position.Y)
		}
		if rect.Width != 150 {
			t.Errorf("expected Width to be 150, got %f", rect.Width)
		}
		if rect.Height != 75 {
			t.Errorf("expected Height to be 75, got %f", rect.Height)
		}
	})
}
