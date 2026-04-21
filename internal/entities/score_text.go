package entities

import (
	"strconv"

	"github.com/MarcelArt/m-engine/pkg/engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ScoreText struct {
	State    *FlappyGameState
	Position rl.Vector2
	FontSize int32
	Color    rl.Color
}

func (s *ScoreText) Start(g *engine.Game) {
}

func (s *ScoreText) Update(g *engine.Game) {
	score := strconv.Itoa(int(s.State.Score))
	rl.DrawText(score, int32(s.Position.X), int32(s.Position.Y), s.FontSize, s.Color)
}

func (s *ScoreText) Destroy(g *engine.Game) {
	s.State = nil
}

var _ engine.Entity = &ScoreText{}
