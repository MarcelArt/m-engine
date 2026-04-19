package entities

import (
	"github.com/MarcelArt/m-engine/pkg/engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type FlappyBird struct {
	Position  rl.Vector2
	Velocity  rl.Vector2
	JumpForce float32
}

func (f *FlappyBird) Start(g *engine.Game) {
	// Hardcoded for now
	f.JumpForce = 400
}

func (f *FlappyBird) Update(g *engine.Game) {
	if rl.IsKeyPressed(rl.KeySpace) {
		f.Velocity.Y = -f.JumpForce
	}

	rl.DrawRectangle(int32(f.Position.X), int32(f.Position.Y), 100, 100, rl.Blue)
}

func (f *FlappyBird) Destroy(g *engine.Game) {
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

var _ engine.Entity = &FlappyBird{}
