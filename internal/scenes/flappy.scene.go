package scenes

import (
	"github.com/MarcelArt/m-engine/internal/entities"
	"github.com/MarcelArt/m-engine/pkg/engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type FlappyScene struct {
	Entities []engine.Entity
}

func (s *FlappyScene) Start(g *engine.Game) {
	flappyBird := &entities.FlappyBird{
		Position:  rl.NewVector2(350, 100),
		JumpForce: 400,
	}

	s.Entities = append(s.Entities, flappyBird)

	physics := engine.NewPhysicsSystem(rl.NewVector2(0, 1), 800)
	physics.AddEntity(flappyBird)
	g.SetPhysicsSystem(physics)
}

func (s *FlappyScene) Update(g *engine.Game) {
	rl.ClearBackground(rl.RayWhite)

	for _, e := range s.Entities {
		e.Update(g)
	}
}

func (s *FlappyScene) Destroy(g *engine.Game) {
}

var _ engine.Scene = &FlappyScene{}
