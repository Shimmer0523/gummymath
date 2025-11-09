package gummymath

// Rect2D は2Dの長方形を表します。
type Rect2D struct {
	Position Vector2D
	Width    float64
	Height   float64
}

// NewRect2D は新しい2Dの長方形を作成します。
func NewRect2D(x, y, width, height float64) Rect2D {
	return Rect2D{
		Position: NewVector2D(x, y),
		Width:    width,
		Height:   height,
	}
}

// Corners は長方形の4つの角を返します。
// 順序は左上、右上、右下、左下です。
func (r Rect2D) Corners() [4]Vector2D {
	return [4]Vector2D{
		r.Position,
		NewVector2D(r.Position.X+r.Width, r.Position.Y),
		NewVector2D(r.Position.X+r.Width, r.Position.Y+r.Height),
		NewVector2D(r.Position.X, r.Position.Y+r.Height),
	}
}
