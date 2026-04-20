package scenes

import (
	"github.com/MarcelArt/m-engine/internal/entities"
	"github.com/MarcelArt/m-engine/pkg/engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type FlappyScene struct {
	engine.Scene
}

func (s *FlappyScene) Start(g *engine.Game) {
	flappyBird := &entities.FlappyBird{
		Position:  rl.NewVector2(350, 100),
		JumpForce: 400,
		Sprite:    engine.NewSpritesheet("assets/Bird1-7.png", rl.NewVector2(16, 16), 3),
		// ColliderRect: rl.NewRectangle(350, 100, 48, 48),
	}

	obstacleSprite := engine.NewSpritesheet("assets/PipeStyle1.png", rl.NewVector2(32, 20), 3)
	spawner := &entities.ObstacleSpawner{
		Position:      rl.NewVector2(900, 0),
		SpawnRate:     2,
		FloorObstacle: entities.FloorObstaclePrefab(obstacleSprite, rl.NewVector2(900, 0)),
		CeilObstacle:  entities.CeilingObstaclePrefab(obstacleSprite, rl.NewVector2(900, 0)),
	}

	s.AddEntity(flappyBird)
	s.AddEntity(spawner)

	physics := engine.NewPhysicsSystem(rl.NewVector2(0, 1), 800)
	physics.AddEntity(flappyBird)
	g.SetPhysicsSystem(physics)

	collision := engine.NewCollisionSystem(true)
	collision.AddRectCollidable(flappyBird)
	g.SetCollisionSystem(collision)
}

func (s *FlappyScene) Update(g *engine.Game) {
	rl.ClearBackground(rl.RayWhite)

	for _, e := range s.GetEntities() {
		e.Update(g)
	}
}

func (s *FlappyScene) Destroy(g *engine.Game) {
}

var _ engine.IScene = &FlappyScene{}
