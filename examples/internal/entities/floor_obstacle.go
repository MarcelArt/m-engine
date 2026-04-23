package entities

import (
	"github.com/MarcelArt/m-engine/examples/internal/enums"
	"github.com/MarcelArt/m-engine/pkg/engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type FloorObstacle struct {
	Sprite   *engine.Spritesheet
	State    *FlappyGameState
	Position rl.Vector2
	Velocity rl.Vector2
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
	o.Sprite = nil
}

func (o *FloorObstacle) GetPosition() rl.Vector2 {
	return o.Position
}

func (o *FloorObstacle) GetVelocity() rl.Vector2 {
	return o.Velocity
}

func (o *FloorObstacle) SetPosition(pos rl.Vector2) {
	o.Position = pos
}

func (o *FloorObstacle) SetVelocity(vel rl.Vector2) {
	o.Velocity = vel
}

func (FloorObstacle) IsGravityEnabled() bool {
	return false
}

func (o *FloorObstacle) GetColliderRect() rl.Rectangle {
	actualHeight := o.Sprite.TileSize.Y * o.Sprite.Scale * float32(o.Height+1)
	screenHeight := rl.GetScreenHeight()
	tileHeight := o.Sprite.TileSize.Y
	tileScale := o.Sprite.Scale

	segmentPos := rl.NewVector2(
		o.Position.X,
		float32(screenHeight)-tileHeight*tileScale*float32(o.Height),
	)

	return rl.NewRectangle(
		segmentPos.X,
		segmentPos.Y,
		o.Sprite.TileSize.X*o.Sprite.Scale,
		actualHeight,
	)
}

func (o *FloorObstacle) SetColliderRect(rl.Rectangle) {}

func (o *FloorObstacle) OnCollision(g *engine.Game, other engine.RectCollidable) {}

func (o *FloorObstacle) OnCollisionEnter(g *engine.Game, other engine.RectCollidable) {
	if _, ok := other.(*FlappyBird); ok {
		o.State.GameOver()
	}
}

func FloorObstaclePrefab(sprite *engine.Spritesheet, pos rl.Vector2) func(height uint, velocity rl.Vector2, state *FlappyGameState) *FloorObstacle {
	return func(height uint, velocity rl.Vector2, state *FlappyGameState) *FloorObstacle {
		return &FloorObstacle{
			Sprite:   sprite,
			Position: pos,
			Height:   height,
			Velocity: velocity,
			State:    state,
		}
	}
}
