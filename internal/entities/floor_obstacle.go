package entities

import (
	"github.com/MarcelArt/m-engine/internal/enums"
	"github.com/MarcelArt/m-engine/pkg/engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type FloorObstacle struct {
	Sprite   *engine.Spritesheet
	Position rl.Vector2
	Height   uint
}

func (o *FloorObstacle) Start(g *engine.Game) {
}

func (o *FloorObstacle) Update(g *engine.Game) {
	screenHeight := rl.GetScreenHeight()
	segmentPos := rl.NewVector2(o.Position.X, float32(screenHeight))
	tileHeight := o.Sprite.TileSize.Y
	tileScale := o.Sprite.Scale

	for i := 0; i < int(o.Height); i++ {
		o.Sprite.DrawTile(enums.MiddlePipe, segmentPos, rl.White)
		segmentPos.Y -= tileHeight * tileScale
	}
	o.Sprite.DrawTile(enums.TopPipe, segmentPos, rl.White)
}

func (o *FloorObstacle) Destroy(g *engine.Game) {
}

func FloorObstaclePrefab(sprite *engine.Spritesheet, pos rl.Vector2, height uint) func() *FloorObstacle {
	return func() *FloorObstacle {
		return &FloorObstacle{
			Sprite:   sprite,
			Position: pos,
			Height:   height,
		}
	}
}
