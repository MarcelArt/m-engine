package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type PhysicsSystem struct {
	Gravity  rl.Vector2
	GravMult float32
	Entities []Movable
}

func NewPhysicsSystem(gravity rl.Vector2, gravMult float32) *PhysicsSystem {
	return &PhysicsSystem{
		Gravity:  gravity,
		GravMult: gravMult,
		Entities: make([]Movable, 0),
	}
}

func (p *PhysicsSystem) Update() {
	dt := rl.GetFrameTime()

	for _, e := range p.Entities {
		p.handleGravity(e, dt)
	}
}

func (p *PhysicsSystem) AddEntity(e Movable) {
	p.Entities = append(p.Entities, e)
}

func (p *PhysicsSystem) handleGravity(e Movable, dt float32) {
	pos := e.GetPosition()
	vel := e.GetVelocity()

	vel.X += p.GravMult * p.Gravity.X * dt
	pos.X += vel.X * dt

	vel.Y += p.GravMult * p.Gravity.Y * dt
	pos.Y += vel.Y * dt

	e.SetPosition(pos)
	e.SetVelocity(vel)
}
