package entities

import (
	"github.com/MarcelArt/m-engine/examples/internal/enums"
	"github.com/MarcelArt/m-engine/pkg/engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type CeilingObstacle struct {
	Sprite       *engine.Spritesheet
	State        *FlappyGameState
	Position     rl.Vector2
	Velocity     rl.Vector2
	Height       uint
	baseVelocity float32
}

func (o *CeilingObstacle) Start(g *engine.Game) {
	o.baseVelocity = o.Velocity.X
}

func (o *CeilingObstacle) Update(g *engine.Game) {
	o.Velocity.X = o.baseVelocity - o.State.VelocityModifier
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
	o.Sprite = nil
}

func (o *CeilingObstacle) GetPosition() rl.Vector2 {
	return o.Position
}

func (o *CeilingObstacle) GetVelocity() rl.Vector2 {
	return o.Velocity
}

func (o *CeilingObstacle) SetPosition(pos rl.Vector2) {
	o.Position = pos
}

func (o *CeilingObstacle) SetVelocity(vel rl.Vector2) {
	o.Velocity = vel
}

func (CeilingObstacle) IsGravityEnabled() bool {
	return false
}

func (o *CeilingObstacle) GetColliderRect() rl.Rectangle {
	return rl.NewRectangle(
		o.Position.X,
		o.Position.Y,
		o.Sprite.TileSize.X*o.Sprite.Scale,
		o.Sprite.TileSize.Y*o.Sprite.Scale*float32(o.Height+1),
	)
}

func (o *CeilingObstacle) SetColliderRect(rl.Rectangle) {}

func (o *CeilingObstacle) OnCollision(g *engine.Game, other engine.RectCollidable) {}

func (o *CeilingObstacle) OnCollisionEnter(g *engine.Game, other engine.RectCollidable) {
	if _, ok := other.(*FlappyBird); ok {
		o.State.GameOver()
	}
}

func CeilingObstaclePrefab(sprite *engine.Spritesheet, pos rl.Vector2) func(height uint, velocity rl.Vector2, state *FlappyGameState) *CeilingObstacle {
	return func(height uint, velocity rl.Vector2, state *FlappyGameState) *CeilingObstacle {
		return &CeilingObstacle{
			Sprite:   sprite,
			Position: pos,
			Height:   height,
			Velocity: velocity,
			State:    state,
		}
	}
}

var _ engine.RectCollidable = &CeilingObstacle{}
