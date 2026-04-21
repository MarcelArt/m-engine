package entities

import (
	"log"

	"github.com/MarcelArt/m-engine/pkg/engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type FlappyBird struct {
	Position  rl.Vector2
	Velocity  rl.Vector2
	JumpForce float32
	Sprite    *engine.Spritesheet
}

func (f *FlappyBird) Start(g *engine.Game) {
	// Hardcoded for now
	f.JumpForce = 400
}

func (f *FlappyBird) Update(g *engine.Game) {
	if rl.IsKeyPressed(rl.KeySpace) {
		f.Velocity.Y = -f.JumpForce
	}

	// rl.DrawRectangle(int32(f.Position.X), int32(f.Position.Y), 100, 100, rl.Blue)
	f.Sprite.DrawTile(0, f.Position, rl.White)
}

func (f *FlappyBird) Destroy(g *engine.Game) {
	f.Sprite = nil
}

func (f *FlappyBird) GetPosition() rl.Vector2 {
	return f.Position
}

func (f *FlappyBird) GetVelocity() rl.Vector2 {
	return f.Velocity
}

func (f *FlappyBird) SetPosition(pos rl.Vector2) {
	f.Position = pos
}

func (f *FlappyBird) SetVelocity(vel rl.Vector2) {
	f.Velocity = vel
}

func (f *FlappyBird) IsGravityEnabled() bool {
	return true
}

func (f *FlappyBird) GetColliderRect() rl.Rectangle {
	return rl.NewRectangle(f.Position.X, f.Position.Y, f.Sprite.TileSize.X*f.Sprite.Scale, f.Sprite.TileSize.Y*f.Sprite.Scale)
}

func (f *FlappyBird) SetColliderRect(rect rl.Rectangle) {
	// f.ColliderRect = rect
}

func (f *FlappyBird) OnCollisionEnter(g *engine.Game, other engine.RectCollidable) {
	log.Println("Collision ENTER with", other)
}

func (f *FlappyBird) OnCollisionExit(g *engine.Game, other engine.RectCollidable) {
	// log.Println("Collision EXIT with", other)
}

func (f *FlappyBird) OnCollision(g *engine.Game, other engine.RectCollidable) {
	// log.Println("Collision with", other)
}

var _ engine.Entity = &FlappyBird{}
var _ engine.RectCollidable = &FlappyBird{}
