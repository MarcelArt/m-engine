package entities

import (
	"github.com/MarcelArt/m-engine/examples/internal/enums"
	"github.com/MarcelArt/m-engine/pkg/engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type SafeGap struct {
	Rect         rl.Rectangle
	Velocity     rl.Vector2
	State        *FlappyGameState
	baseVelocity float32
}

func (s *SafeGap) Start(g *engine.Game) {
	s.baseVelocity = s.Velocity.X
}

func (s *SafeGap) Update(g *engine.Game) {
	s.Velocity.X = s.baseVelocity - s.State.VelocityModifier
	// log.Println("s.Velocity :>> ", s.Velocity)
}

func (s *SafeGap) Destroy(g *engine.Game) {
	s.State = nil
}

func (s *SafeGap) GetPosition() rl.Vector2 {
	return rl.NewVector2(s.Rect.X, s.Rect.Y)
}

func (s *SafeGap) GetVelocity() rl.Vector2 {
	return s.Velocity
}

func (s *SafeGap) SetPosition(pos rl.Vector2) {
	s.Rect.X = pos.X
	s.Rect.Y = pos.Y
}

func (s *SafeGap) SetVelocity(vel rl.Vector2) {
	s.Velocity = vel
}

func (s *SafeGap) IsGravityEnabled() bool {
	return false
}

func (s *SafeGap) GetColliderRect() rl.Rectangle {
	return s.Rect
}

func (s *SafeGap) SetColliderRect(rect rl.Rectangle) {
	s.Rect = rect
}

func (s *SafeGap) OnCollision(g *engine.Game, other engine.RectCollidable) {
}

func (s *SafeGap) OnCollisionExit(g *engine.Game, other engine.RectCollidable) {
	if _, ok := other.(*FlappyBird); ok && s.State.State == enums.StatePlaying {
		s.State.ScoreUp()
	}
}

func SafeGapPrefab() func(rect rl.Rectangle, velocity rl.Vector2, state *FlappyGameState) *SafeGap {
	return func(rect rl.Rectangle, velocity rl.Vector2, state *FlappyGameState) *SafeGap {
		return &SafeGap{
			Rect:     rect,
			Velocity: velocity,
			State:    state,
		}
	}
}

var _ engine.Entity = &SafeGap{}
var _ engine.RectCollidable = &SafeGap{}
var _ engine.PhysicsObject = &SafeGap{}
