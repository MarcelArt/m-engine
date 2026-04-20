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
		Sprite:    engine.NewSpritesheet("assets/Bird1-7.png", rl.NewVector2(16, 16), 3),
	}

	ceilObstacle := &entities.CeilingObstacle{
		Sprite:   engine.NewSpritesheet("assets/PipeStyle1.png", rl.NewVector2(32, 20), 3),
		Position: rl.NewVector2(500, 0),
		Velocity: rl.NewVector2(-200, 0),
		Height:   5,
	}
	floorObstacle := &entities.FloorObstacle{
		Sprite:   engine.NewSpritesheet("assets/PipeStyle1.png", rl.NewVector2(32, 20), 3),
		Position: rl.NewVector2(500, 0),
		Velocity: rl.NewVector2(-200, 0),
		Height:   8,
	}

	s.Entities = append(s.Entities, flappyBird)
	s.Entities = append(s.Entities, ceilObstacle)
	s.Entities = append(s.Entities, floorObstacle)

	physics := engine.NewPhysicsSystem(rl.NewVector2(0, 1), 800)
	physics.AddEntity(flappyBird)
	physics.AddEntity(ceilObstacle)
	physics.AddEntity(floorObstacle)
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
