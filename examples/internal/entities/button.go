package entities

import (
	"image/color"

	"github.com/MarcelArt/m-engine/pkg/engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Button struct {
	Rect       rl.Rectangle
	Text       string
	FontSize   int32
	Color      color.RGBA
	HoverColor color.RGBA
	OnClick    func()
}

func (b *Button) Start(g *engine.Game) {
}

func (b *Button) Update(g *engine.Game) {
	mouse := rl.GetMousePosition()
	isHovered := rl.CheckCollisionPointRec(mouse, b.Rect)

	color := b.Color
	if isHovered {
		color = b.HoverColor
	}

	rl.DrawRectangleRec(b.Rect, color)

	textWidth := rl.MeasureText(b.Text, b.FontSize)
	textX := int32(b.Rect.X + (b.Rect.Width-float32(textWidth))/2)
	textY := int32(b.Rect.Y + (b.Rect.Height / 2) - 10)

	rl.DrawText(b.Text, textX, textY, b.FontSize, rl.White)

	if isHovered && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		b.OnClick()
	}
}

func (b *Button) Destroy(g *engine.Game) {
}

var _ engine.Entity = &Button{}
