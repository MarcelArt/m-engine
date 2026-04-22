package engine

import (
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	width           int32
	height          int32
	title           string
	targetFPS       int32
	SceneManager    *SceneManager
	PhysicsSystem   *PhysicsSystem
	CollisionSystem *CollisionSystem
	GameSave        GameSave
}

func NewGame(width int32, height int32, title string, targetFPS int32) *Game {
	return &Game{
		width:           width,
		height:          height,
		title:           title,
		targetFPS:       targetFPS,
		SceneManager:    NewSceneManager(),
		PhysicsSystem:   nil,
		CollisionSystem: nil,
		GameSave:        nil,
	}
}

func (g *Game) Start() {
	rl.InitWindow(g.width, g.height, g.title)
	defer rl.CloseWindow()

	rl.SetTargetFPS(g.targetFPS)
	rl.SetWindowState(rl.FlagWindowResizable)

	g.SceneManager.Start(g)
	for !rl.WindowShouldClose() {
		g.SceneManager.Update(g)
		g.handlePhysics()
		g.handleCollision()
	}
}

func (g *Game) SetPhysicsSystem(physicSystem *PhysicsSystem) {
	g.PhysicsSystem = physicSystem
}

func (g *Game) SetCollisionSystem(collisionSystem *CollisionSystem) {
	g.CollisionSystem = collisionSystem
}

func (g *Game) SetGameSave(gameSave GameSave) {
	g.GameSave = gameSave
}

func (g *Game) ExitGame() {
	rl.CloseWindow()
	os.Exit(0)
}

func (g *Game) handlePhysics() {
	if g.PhysicsSystem != nil {
		g.PhysicsSystem.Update()
	}
}

func (g *Game) handleCollision() {
	if g.CollisionSystem != nil {
		g.CollisionSystem.Update(g)
	}
}
