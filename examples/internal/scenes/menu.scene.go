package scenes

import (
	"log"

	"github.com/MarcelArt/m-engine/examples/internal/entities"
	"github.com/MarcelArt/m-engine/pkg/engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type MenuScene struct {
	engine.Scene
}

func (m *MenuScene) Start(g *engine.Game) {
	newGameBtn := &entities.Button{
		Rect:       rl.NewRectangle(300, 200, 200, 50),
		Text:       "New Game",
		FontSize:   20,
		Color:      rl.DarkGray,
		HoverColor: rl.Gray,
		OnClick: func() {
			log.Println("Clicked new game")
			g.SceneManager.LoadScene(g, "flappy")
		},
	}
	settingsBtn := &entities.Button{
		Rect:       rl.NewRectangle(300, 270, 200, 50),
		Text:       "Settings",
		FontSize:   20,
		Color:      rl.DarkGray,
		HoverColor: rl.Gray,
		OnClick: func() {
			log.Println("Settings clicked")
		},
	}
	exitBtn := &entities.Button{
		Rect:       rl.NewRectangle(300, 340, 200, 50),
		Text:       "Exit",
		FontSize:   20,
		Color:      rl.DarkGray,
		HoverColor: rl.Gray,
		OnClick: func() {
			log.Println("Exit clicked")
			g.ExitGame()
		},
	}

	m.AddEntity(newGameBtn)
	m.AddEntity(settingsBtn)
	m.AddEntity(exitBtn)
}

func (m *MenuScene) Update(g *engine.Game) {
	rl.ClearBackground(rl.Black)
	rl.DrawText("Main Menu", 320, 100, 30, rl.White)
}

func (m *MenuScene) Destroy(g *engine.Game) {
	m.ClearEntities()
}

var _ engine.IScene = &MenuScene{}
