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
	// rl.DrawTexture(b.Texture, 0, 0, rl.White)

	screenHeight := float32(rl.GetScreenHeight())

	width := float32(b.Texture.Width)
	height := float32(b.Texture.Height)
	source := rl.NewRectangle(0, 0, width, height)
	dest := rl.NewRectangle(0, screenHeight-(height*3), width*3, height*3)

	rl.DrawTexturePro(b.Texture, source, dest, rl.Vector2Zero(), 0, rl.White)
}

func (b *Background) Destroy(g *engine.Game) {

}

var _ engine.Entity = &Background{}
