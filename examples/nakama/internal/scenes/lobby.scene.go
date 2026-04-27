package scenes

import (
	"fmt"
	"log"

	"github.com/MarcelArt/m-engine/examples/nakama/internal/entities"
	"github.com/MarcelArt/m-engine/pkg/engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type LobbyScene struct {
	engine.Scene
}

func (s *LobbyScene) Start(g *engine.Game) {
	account, err := g.Nakama.GetAccount()
	if err != nil {
		g.SceneManager.LoadScene(g, "login")
		return
	}

	welcomeText := &entities.Label{
		Text:      fmt.Sprintf("Welcome, %s", account.Email),
		FontSize:  20,
		TextColor: rl.White,
		Position:  rl.NewVector2(10, 10),
		IsActive:  true,
	}

	s.AddUIEntity(welcomeText)

	matches, _ := g.Nakama.GetMatches(10, false, 0, 4)
	log.Println("matches :>> ", matches)
}

func (s *LobbyScene) Update(g *engine.Game) {
	rl.ClearBackground(rl.Black)
}

func (s *LobbyScene) Destroy(g *engine.Game) {
	s.ClearEntities()
}

var _ engine.IScene = &LobbyScene{}
