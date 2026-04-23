package entities

import (
	"fmt"

	"github.com/MarcelArt/m-engine/examples/internal/enums"
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
	textX := int32(s.Position.X) - (textWidth / 2)
	textY := int32(s.Position.Y)

	padding := int32(15)
	bgWidth := textWidth + (padding * 2)
	bgHeight := s.FontSize + (padding * 2)
	bgX := int32(s.Position.X) - (bgWidth / 2)
	bgY := textY - (padding / 2)

	bgColor := rl.NewColor(30, 30, 40, 200)
	borderColor := rl.NewColor(150, 150, 150, 255)

	rl.DrawRectangle(bgX, bgY, bgWidth, bgHeight, bgColor)
	rl.DrawRectangleLines(bgX, bgY, bgWidth, bgHeight, borderColor)

	rl.DrawText(score, textX, textY, s.FontSize, s.Color)
}

func (s *HighScoreText) Destroy(g *engine.Game) {
}

var _ engine.Entity = &HighScoreText{}
