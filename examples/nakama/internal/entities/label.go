package entities

import (
	"github.com/MarcelArt/m-engine/pkg/engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Label struct {
	Text      string
	FontSize  int32
	TextColor rl.Color
	Position  rl.Vector2
	IsActive  bool
}

func (l *Label) Start(g *engine.Game) {
}

func (l *Label) Update(g *engine.Game) {
	if l.IsActive {
		rl.DrawText(l.Text, int32(l.Position.X), int32(l.Position.Y), l.FontSize, l.TextColor)
	}
}

func (l *Label) Destroy(g *engine.Game) {
	l.IsActive = false
}

var _ engine.Entity = &Label{}
