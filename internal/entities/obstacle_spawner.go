package entities

import (
	"github.com/MarcelArt/m-engine/pkg/engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ObstacleSpawner struct {
	Position      rl.Vector2
	SpawnRate     float32 // in seconds
	timer         float32
	State         *FlappyGameState
	FloorObstacle func(height uint, velocity rl.Vector2, state *FlappyGameState) *FloorObstacle
	CeilObstacle  func(height uint, velocity rl.Vector2, state *FlappyGameState) *CeilingObstacle
	SafeGap       func(rect rl.Rectangle, velocity rl.Vector2, state *FlappyGameState) *SafeGap
}

func (o *ObstacleSpawner) Start(g *engine.Game) {
}

func (o *ObstacleSpawner) Update(g *engine.Game) {
	o.timer += rl.GetFrameTime()
	rng := rl.GetRandomValue(1, 12)
	if o.timer > o.SpawnRate {
		o.timer = 0
		g.SceneManager.GetCurrentScene()

		ceilObstacle := o.CeilObstacle(uint(rng), rl.NewVector2(-200, 0), o.State)
		floorObstacle := o.FloorObstacle(uint(13-rng), rl.NewVector2(-200, 0), o.State)

		ceilRect := ceilObstacle.GetColliderRect()
		floorRect := floorObstacle.GetColliderRect()
		safeGap := o.SafeGap(rl.NewRectangle(o.Position.X, ceilRect.Height, ceilRect.Width, floorRect.Y-ceilRect.Height), rl.NewVector2(-200, 0), o.State)

		g.PhysicsSystem.AddEntity(ceilObstacle)
		g.PhysicsSystem.AddEntity(floorObstacle)
		g.PhysicsSystem.AddEntity(safeGap)

		g.CollisionSystem.AddRectCollidable(ceilObstacle)
		g.CollisionSystem.AddRectCollidable(floorObstacle)
		g.CollisionSystem.AddRectCollidable(safeGap)

		g.SceneManager.GetCurrentScene().AddEntity(ceilObstacle)
		g.SceneManager.GetCurrentScene().AddEntity(floorObstacle)
		g.SceneManager.GetCurrentScene().AddEntity(safeGap)
	}
}

func (o *ObstacleSpawner) Destroy(g *engine.Game) {
	o.State = nil
	o.CeilObstacle = nil
	o.FloorObstacle = nil
	o.SafeGap = nil
}

var _ engine.Entity = &ObstacleSpawner{}
