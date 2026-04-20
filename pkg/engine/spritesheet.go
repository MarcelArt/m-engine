package engine

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Spritesheet struct {
	Texture  rl.Texture2D
	TileSize rl.Vector2
	Scale    float32
}

func NewSpritesheet(path string, tileSize rl.Vector2, scale float32) *Spritesheet {
	return &Spritesheet{
		Texture:  rl.LoadTexture(path),
		TileSize: tileSize,
		Scale:    scale,
	}
}

func (s *Spritesheet) DrawTile(tileIndex int, pos rl.Vector2, tint color.RGBA) {
	tilesPerRow := s.Texture.Width / int32(s.TileSize.X)

	srcX := (tileIndex % int(tilesPerRow)) * int(s.TileSize.X)
	srcY := (tileIndex / int(tilesPerRow)) * int(s.TileSize.Y)

	source := rl.Rectangle{
		X:      float32(srcX),
		Y:      float32(srcY),
		Width:  s.TileSize.X,
		Height: s.TileSize.Y,
	}

	dest := rl.Rectangle{
		X:      pos.X,
		Y:      pos.Y,
		Width:  s.TileSize.X * s.Scale,
		Height: s.TileSize.Y * s.Scale,
	}

	origin := rl.Vector2{X: 0, Y: 0}

	rl.DrawTexturePro(s.Texture, source, dest, origin, 0, tint)
}
