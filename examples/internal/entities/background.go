package entities

import (
	"github.com/MarcelArt/m-engine/pkg/engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Background struct {
	Texture rl.Texture2D
}

func (b *Background) Start(g *engine.Game) {}

func (b *Background) Update(g *engine.Game) {
	screenHeight := float32(rl.GetScreenHeight())

	width := float32(b.Texture.Width)
	height := float32(b.Texture.Height)

	scale := screenHeight / height

	source := rl.NewRectangle(0, 0, width, height)
	dest := rl.NewRectangle(0, screenHeight-(height*scale), width*scale, height*scale)

	rl.DrawTexturePro(b.Texture, source, dest, rl.Vector2Zero(), 0, rl.White)
}

func (b *Background) Destroy(g *engine.Game) {

}

var _ engine.Entity = &Background{}
