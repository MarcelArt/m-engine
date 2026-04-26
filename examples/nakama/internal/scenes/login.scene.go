package scenes

import (
	"log"

	"github.com/MarcelArt/m-engine/examples/nakama/internal/entities"
	"github.com/MarcelArt/m-engine/pkg/engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type LoginScene struct {
	engine.Scene
	email    string
	password string
}

func (s *LoginScene) Start(g *engine.Game) {
	screenHeight := rl.GetScreenHeight()
	screenWidth := rl.GetScreenWidth()

	errorMsg := &entities.Label{
		Text:      "",
		FontSize:  20,
		TextColor: rl.Red,
		Position:  rl.NewVector2(10, 10),
		IsActive:  false,
	}

	emailInput := &entities.TextInput{
		Value:       "",
		Placeholder: "Email",
		IsDisabled:  false,
		IsSecret:    false,
		OnValueChange: func(value string) {
			s.email = value
		},
		BackgroundColor: rl.White,
		TextColor:       rl.Black,
		FontSize:        20,
		Rect:            rl.NewRectangle(float32(screenWidth)/2, (float32(screenHeight)/2)-100, 300, 50),
	}
	passwordInput := &entities.TextInput{
		Value:       "",
		Placeholder: "Password",
		IsDisabled:  false,
		IsSecret:    true,
		OnValueChange: func(value string) {
			s.password = value
		},
		BackgroundColor: rl.White,
		TextColor:       rl.Black,
		FontSize:        20,
		Rect:            rl.NewRectangle(float32(screenWidth)/2, float32(screenHeight)/2, 300, 50),
	}
	loginBtn := &entities.Button{
		Rect:       rl.NewRectangle(float32(screenWidth)/2, (float32(screenHeight)/2)+100, 300, 50),
		Text:       "Login",
		FontSize:   20,
		Color:      rl.DarkGray,
		HoverColor: rl.Gray,
		OnClick: func() {
			log.Printf("login")
			if err := g.Nakama.AuthenticateEmail(s.email, s.password, s.email); err != nil {
				errorMsg.Text = "Email or password is incorrect"
				errorMsg.IsActive = true
				return
			}
			errorMsg.IsActive = false
			errorMsg.Text = ""
			g.SceneManager.LoadScene(g, "lobby")
		},
	}

	s.AddUIEntity(emailInput)
	s.AddUIEntity(passwordInput)
	s.AddUIEntity(loginBtn)
	s.AddUIEntity(errorMsg)
}

func (s *LoginScene) Update(g *engine.Game) {
	rl.ClearBackground(rl.Black)
}

func (s *LoginScene) Destroy(g *engine.Game) {

}

var _ engine.IScene = &LoginScene{}
