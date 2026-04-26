package main

import (
	"github.com/MarcelArt/m-engine/examples/nakama/internal/scenes"
	"github.com/MarcelArt/m-engine/pkg/engine"
	"github.com/MarcelArt/m-engine/pkg/nakama"
)

// func main() {
// 	n := nakama.New("https://api-nakama.bangmarcel.art", "defaultkey")

// 	if err := n.AuthenticateEmail("marcel@yopmail.com", "<string>", "marcel"); err != nil {
// 		log.Fatalln(err.Error())
// 	}

// 	account, err := n.GetAccount()
// 	if err != nil {
// 		log.Fatalln(err.Error())
// 	}

// 	log.Println(account)
// }

func main() {
	game := engine.NewGame(671, 1030, "Nakama", 60)

	game.SceneManager.Register("login", &scenes.LoginScene{})
	game.SceneManager.Register("lobby", &scenes.LobbyScene{})
	game.SceneManager.SetDefaultScene("login")

	game.SetNakama(nakama.New("https://api-nakama.bangmarcel.art", "defaultkey"))

	game.Start()
}
