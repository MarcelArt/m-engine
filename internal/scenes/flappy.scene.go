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
	state := &entities.FlappyGameState{
		Score:      0,
		IsGameOver: false,
	}

	screenWidth := rl.GetScreenWidth()
	scoreTxt := &entities.ScoreText{
		State:    state,
		Position: rl.NewVector2(float32(screenWidth/2), 50),
		FontSize: 48,
		Color:    rl.Black,
	}

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
		SafeGap:       entities.SafeGapPrefab(),
		State:         state,
	}

	background := &entities.Background{
		Texture: rl.LoadTexture("assets/Background2.png"),
	}

	s.AddUIEntity(scoreTxt)

	s.AddEntity(background)
	s.AddEntity(flappyBird)
	s.AddEntity(spawner)
	s.AddEntity(state)

	physics := engine.NewPhysicsSystem(rl.NewVector2(0, 1), 800)
	physics.AddEntity(flappyBird)
	g.SetPhysicsSystem(physics)

	collision := engine.NewCollisionSystem(true)
	collision.AddRectCollidable(flappyBird)
	g.SetCollisionSystem(collision)

	for _, e := range s.GetEntities() {
		e.Start(g)
	}

	for _, ui := range s.GetUIEntities() {
		ui.Start(g)
	}
}

func (s *FlappyScene) Update(g *engine.Game) {
	rl.ClearBackground(rl.RayWhite)

	for _, e := range s.GetEntities() {
		e.Update(g)
	}

	for _, ui := range s.GetUIEntities() {
		ui.Update(g)
	}
}

func (s *FlappyScene) Destroy(g *engine.Game) {
	// g.PhysicsSystem = nil
	// g.CollisionSystem = nil
	// g.SceneManager.GetCurrentScene().ClearEntities()
}

var _ engine.IScene = &FlappyScene{}
