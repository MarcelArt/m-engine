package entities

import (
	"github.com/MarcelArt/m-engine/pkg/engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type CeilingObstacle struct {
	Sprite   *engine.Spritesheet
	Position rl.Vector2
	Height   uint
}

func (c *CeilingObstacle) Start(g *engine.Game) {
}

func (c *CeilingObstacle) Update(g *engine.Game) {
}

func (c *CeilingObstacle) Destroy(g *engine.Game) {
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
