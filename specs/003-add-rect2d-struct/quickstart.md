# Rect2D クイックスタート

このガイドは、`Rect2D`構造体の基本的な使用方法を示します。

## Goでの使用例

以下は、`Rect2D`を作成し、そのプロパティにアクセスし、角の座標を取得する方法の例です。

```go
package main

import (
	"fmt"
	"github.com/Shimmer0523/gummymath"
)

func main() {
	// 1. 新しいRect2Dを作成
	// 位置(10, 20)、サイズ(幅100、高さ50)
	rect := gummymath.NewRect2D(10, 20, 100, 50)
	fmt.Printf("Initial Rectangle: %+v\n", rect)

	// 2. プロパティにアクセス
	fmt.Printf("Position: %+v, Width: %f, Height: %f\n", rect.Position, rect.Width, rect.Height)

	// 3. プロパティを変更
	rect.Position = gummymath.NewVector2D(30, 40)
	rect.Width = 150
	fmt.Printf("Modified Rectangle: %+v\n", rect)

	// 4. 4つの角の座標を取得
	corners := rect.Corners()
	fmt.Println("Corner points:")
	fmt.Printf("  Top-Left: %+v\n", corners[0])
	fmt.Printf("  Top-Right: %+v\n", corners[1])
	fmt.Printf("  Bottom-Right: %+v\n", corners[2])
	fmt.Printf("  Bottom-Left: %+v\n", corners[3])
}
```

