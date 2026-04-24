package entities

import (
	"log"

	"github.com/MarcelArt/m-engine/examples/flappy/internal/enums"
	"github.com/MarcelArt/m-engine/pkg/engine"
)

type FlappyGameState struct {
	Score            uint
	HighScore        uint
	gameSave         engine.GameSave
	State            int
	VelocityModifier float32
}

func (f *FlappyGameState) Start(g *engine.Game) {
	f.Score = 0
	f.State = enums.StateMenu
	f.VelocityModifier = 0

	f.LoadSaveFile(g)
}

func (f *FlappyGameState) Update(g *engine.Game) {
	// if f.State == enums.StateGameOver {
	// 	g.SceneManager.LoadScene(g, "menu")
	// 	return
	// }
}

func (f *FlappyGameState) Destroy(g *engine.Game) {
	f.Score = 0
	f.State = enums.StateMenu
}

func (f *FlappyGameState) ScoreUp() {
	log.Println("f.VelocityModifier :>> ", f.VelocityModifier)
	f.Score++
	f.VelocityModifier += 10
	log.Println("f.VelocityModifier :>> ", f.VelocityModifier)
}

func (f *FlappyGameState) GameOver() {
	f.State = enums.StateGameOver

	if f.Score > f.HighScore {
		f.HighScore = f.Score
		f.gameSave.Set("highScore", float64(f.HighScore))
		f.gameSave.Save()
	}
}

func (f *FlappyGameState) LoadSaveFile(g *engine.Game) {
	f.gameSave = g.GameSave
	f.gameSave.Load()

	highScore, ok := f.gameSave.Get("highScore").(float64)
	if !ok {
		f.HighScore = 0
		return
	}

	f.HighScore = uint(highScore)
}

var _ engine.Entity = &FlappyGameState{}
