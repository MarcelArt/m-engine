package entities

import "github.com/MarcelArt/m-engine/pkg/engine"

type FlappyGameState struct {
	Score      uint
	IsGameOver bool
}

func (f *FlappyGameState) Start(g *engine.Game) {
	f.Score = 0
	f.IsGameOver = false
}

func (f *FlappyGameState) Update(g *engine.Game) {
	if f.IsGameOver {
		g.SceneManager.LoadScene(g, "menu")
		return
	}

}

func (f *FlappyGameState) Destroy(g *engine.Game) {
	f.Score = 0
	f.IsGameOver = false
}

func (f *FlappyGameState) ScoreUp() {
	f.Score++
}

func (f *FlappyGameState) GameOver() {
	f.IsGameOver = true
}

var _ engine.Entity = &FlappyGameState{}
