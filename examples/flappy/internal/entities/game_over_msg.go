package entities

import (
	"fmt"

	"github.com/MarcelArt/m-engine/examples/flappy/internal/enums"
	"github.com/MarcelArt/m-engine/pkg/engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameOverMsg struct {
	State    *FlappyGameState
	position rl.Vector2
	fontSize int32
	color    rl.Color
}

func (m *GameOverMsg) Start(g *engine.Game) {
	m.position = rl.NewVector2(float32(rl.GetScreenWidth())/2, float32(rl.GetScreenHeight())/2)
	m.fontSize = 40
	m.color = rl.White
}

func (m *GameOverMsg) Update(g *engine.Game) {
	m.drawDialog()

	if m.State.State == enums.StateGameOver && rl.IsKeyPressed(rl.KeySpace) {
		g.SceneManager.LoadScene(g, "flappy")
	}
}

func (m *GameOverMsg) Destroy(g *engine.Game) {
	m.State = nil
}

func (m *GameOverMsg) drawDialog() {
	if m.State.State != enums.StateGameOver {
		return
	}

	centerX := int32(m.position.X)
	centerY := int32(m.position.Y)

	dialogWidth := int32(400)
	dialogHeight := int32(250)
	dialogX := centerX - (dialogWidth / 2)
	dialogY := centerY - (dialogHeight / 2)

	bgColor := rl.NewColor(30, 30, 40, 230)
	borderColor := rl.NewColor(150, 150, 150, 255)
	textColor := rl.White
	accentColor := rl.NewColor(255, 100, 100, 255)

	rl.DrawRectangle(dialogX, dialogY, dialogWidth, dialogHeight, bgColor)
	rl.DrawRectangleLines(dialogX, dialogY, dialogWidth, dialogHeight, borderColor)
	rl.DrawRectangleLines(dialogX-2, dialogY-2, dialogWidth+4, dialogHeight+4, borderColor)

	gameOverText := "GAME OVER"
	gameOverWidth := rl.MeasureText(gameOverText, m.fontSize)
	rl.DrawText(gameOverText, centerX-(gameOverWidth/2), dialogY+30, m.fontSize, accentColor)

	scoreText := fmt.Sprintf("Score: %d", m.State.Score)
	scoreWidth := rl.MeasureText(scoreText, m.fontSize-10)
	rl.DrawText(scoreText, centerX-(scoreWidth/2), dialogY+100, m.fontSize-10, textColor)

	highScoreText := fmt.Sprintf("High Score: %d", m.State.HighScore)
	highScoreWidth := rl.MeasureText(highScoreText, m.fontSize-10)
	rl.DrawText(highScoreText, centerX-(highScoreWidth/2), dialogY+140, m.fontSize-10, textColor)

	restartText := "Press SPACE to restart"
	restartWidth := rl.MeasureText(restartText, m.fontSize-20)
	rl.DrawText(restartText, centerX-(restartWidth/2), dialogY+200, m.fontSize-20, rl.LightGray)
}

var _ engine.Entity = &GameOverMsg{}
