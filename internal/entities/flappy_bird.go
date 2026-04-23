package entities

import (
	"github.com/MarcelArt/m-engine/internal/enums"
	"github.com/MarcelArt/m-engine/pkg/engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type FlappyBird struct {
	Position       rl.Vector2
	Velocity       rl.Vector2
	JumpForce      float32
	Sprite         *engine.Spritesheet
	State          *FlappyGameState
	animationTimer float32
	animationIndex int
	animationSPF   float32
}

func (f *FlappyBird) Start(g *engine.Game) {
	// Hardcoded for now
	f.JumpForce = 400
	f.animationTimer = 0
	f.animationSPF = 0.2
}

func (f *FlappyBird) Update(g *engine.Game) {
	dt := rl.GetFrameTime()
	f.animationTimer += dt
	spf := f.animationSPF

	if rl.IsKeyPressed(rl.KeySpace) && f.State.State != enums.StateGameOver {
		f.StartGame(g)
		f.Velocity.Y = -f.JumpForce
	}

	if f.animationTimer >= spf {
		f.animationTimer = 0
		if f.animationIndex >= 3 {
			f.animationIndex = 0
		} else {
			f.animationIndex++
		}
	}
	f.Sprite.DrawTile(f.animationIndex, f.Position, rl.White)

	if int(f.Position.Y) > rl.GetScreenHeight() || f.Position.Y < 0 {
		f.State.GameOver()
	}
}

func (f *FlappyBird) Destroy(g *engine.Game) {
	f.Sprite = nil
	f.State = nil
}

func (f *FlappyBird) StartGame(g *engine.Game) {
	if f.State.State == enums.StateMenu {
		f.State.State = enums.StatePlaying
		g.PhysicsSystem.AddEntity(f)
	}
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
	// log.Println("Collision ENTER with", other)
}

func (f *FlappyBird) OnCollisionExit(g *engine.Game, other engine.RectCollidable) {
	// log.Println("Collision EXIT with", other)
}

func (f *FlappyBird) OnCollision(g *engine.Game, other engine.RectCollidable) {
	// log.Println("Collision with", other)
}

var _ engine.Entity = &FlappyBird{}
var _ engine.RectCollidable = &FlappyBird{}
