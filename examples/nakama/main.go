package main

import (
	"log"

	"github.com/MarcelArt/m-engine/pkg/nakama"
)

func main() {
	n := nakama.New("https://api-nakama.bangmarcel.art", "defaultkey")

	if err := n.AuthenticateEmail("marcel@yopmail.com", "<string>", "marcel"); err != nil {
		log.Fatalln(err.Error())
	}

	account, err := n.GetAccount()
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println(account)
}
