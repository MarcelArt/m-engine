package entities

import (
	"github.com/MarcelArt/m-engine/internal/enums"
	"github.com/MarcelArt/m-engine/pkg/engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type CeilingObstacle struct {
	Sprite   *engine.Spritesheet
	Position rl.Vector2
	Height   uint
}

func (o *CeilingObstacle) Start(g *engine.Game) {
}

func (o *CeilingObstacle) Update(g *engine.Game) {
	segmentPos := o.Position
	tileHeight := o.Sprite.TileSize.Y
	tileScale := o.Sprite.Scale

	for i := 0; i < int(o.Height); i++ {
		o.Sprite.DrawTile(enums.MiddlePipe, segmentPos, rl.White)
		segmentPos.Y += tileHeight * tileScale
	}
	o.Sprite.DrawTile(enums.BottomPipe, segmentPos, rl.White)
}

func (o *CeilingObstacle) Destroy(g *engine.Game) {
}

func CeilingObstaclePrefab(sprite *engine.Spritesheet, pos rl.Vector2, height uint) func() *CeilingObstacle {
	return func() *CeilingObstacle {
		return &CeilingObstacle{
			Sprite:   sprite,
			Position: pos,
			Height:   height,
		}
	}
}
