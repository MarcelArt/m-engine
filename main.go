package main

import (
	"github.com/MarcelArt/m-engine/internal/scenes"
	"github.com/MarcelArt/m-engine/pkg/engine"
)

func main() {
	game := engine.NewGame(800, 450, "M-Engine", 60)

	game.SceneManager.Register("menu", &scenes.MenuScene{})
	game.SceneManager.Register("flappy", &scenes.FlappyScene{})
	game.SceneManager.SetDefaultScene("flappy")
	game.Start()
}
