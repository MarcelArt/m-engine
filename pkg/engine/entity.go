package engine

import rl "github.com/gen2brain/raylib-go/raylib"

type Entity interface {
	Start(g *Game)
	Update(g *Game)
	Destroy(g *Game)
}

type PhysicsObject interface {
	GetPosition() rl.Vector2
	GetVelocity() rl.Vector2
	SetPosition(rl.Vector2)
	SetVelocity(rl.Vector2)
	IsGravityEnabled() bool
}
