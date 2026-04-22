package engine

import (
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type SceneManager struct {
	scenes       map[string]IScene
	currentScene string
}

func NewSceneManager() *SceneManager {
	return &SceneManager{
		scenes:       make(map[string]IScene),
		currentScene: "",
	}
}

func (sm *SceneManager) Start(g *Game) {
	if sm.currentScene == "" {
		log.Fatalln("Default scene not set, please add SetDefaultScene")
	}

	sm.scenes[sm.currentScene].Start(g)

	for _, e := range sm.scenes[sm.currentScene].GetEntities() {
		e.Start(g)
	}

	for _, ui := range sm.scenes[sm.currentScene].GetUIEntities() {
		ui.Start(g)
	}
}

func (sm *SceneManager) Update(g *Game) {
	rl.BeginDrawing()
	sm.scenes[sm.currentScene].Update(g)
	for _, e := range sm.scenes[sm.currentScene].GetEntities() {
		e.Update(g)
	}

	for _, ui := range sm.scenes[sm.currentScene].GetUIEntities() {
		ui.Update(g)
	}
	rl.EndDrawing()
}

func (sm *SceneManager) LoadScene(g *Game, sceneID string) {
	g.PhysicsSystem = nil
	g.CollisionSystem = nil
	sm.GetCurrentScene().ClearEntities()
	sm.scenes[sm.currentScene].Destroy(g)
	sm.currentScene = sceneID
	sm.Start(g)
}

func (sm *SceneManager) Register(sceneID string, scene IScene) {
	sm.scenes[sceneID] = scene
}

func (sm *SceneManager) Unregister(sceneID string) {
	delete(sm.scenes, sceneID)
}

func (sm *SceneManager) SetDefaultScene(sceneID string) {
	sm.currentScene = sceneID
}

func (sm *SceneManager) GetCurrentScene() IScene {
	return sm.scenes[sm.currentScene]
}
