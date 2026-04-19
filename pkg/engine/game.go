package engine

import (
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	width         int32
	height        int32
	title         string
	targetFPS     int32
	SceneManager  *SceneManager
	physicsSystem *PhysicsSystem
}

func NewGame(width int32, height int32, title string, targetFPS int32) *Game {
	return &Game{
		width:         width,
		height:        height,
		title:         title,
		targetFPS:     targetFPS,
		SceneManager:  NewSceneManager(),
		physicsSystem: nil,
	}
}

func (g *Game) Start() {
	rl.InitWindow(g.width, g.height, g.title)
	defer rl.CloseWindow()

	rl.SetTargetFPS(g.targetFPS)
	rl.SetWindowState(rl.FlagWindowResizable)

	g.SceneManager.Start(g)
	for !rl.WindowShouldClose() {
		g.handlePhysics()
		g.SceneManager.Update(g)
	}
}

func (g *Game) SetPhysicsSystem(physicSystem *PhysicsSystem) {
	g.physicsSystem = physicSystem
}

func (g *Game) handlePhysics() {
	if g.physicsSystem != nil {
		log.Println("physics")
		g.physicsSystem.Update()
	}
}
