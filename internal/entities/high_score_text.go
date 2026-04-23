package entities

import (
	"fmt"

	"github.com/MarcelArt/m-engine/internal/enums"
	"github.com/MarcelArt/m-engine/pkg/engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type HighScoreText struct {
	State    *FlappyGameState
	Position rl.Vector2
	FontSize int32
	Color    rl.Color
}

func (s *HighScoreText) Start(g *engine.Game) {
}

func (s *HighScoreText) Update(g *engine.Game) {
	if s.State.State != enums.StateMenu {
		return
	}
	score := fmt.Sprintf("High Score: %d", s.State.HighScore)
	textWidth := rl.MeasureText(score, s.FontSize)
	rl.DrawText(score, int32(s.Position.X)-(textWidth/2), int32(s.Position.Y), s.FontSize, s.Color)
}

func (s *HighScoreText) Destroy(g *engine.Game) {
}

var _ engine.Entity = &HighScoreText{}
