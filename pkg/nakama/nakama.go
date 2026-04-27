package nakama

import (
	"net/http"
)

type NakamaSession struct {
	accessToken  string
	refreshToken string
}

type Nakama struct {
	Session   NakamaSession
	client    *http.Client
	baseURL   string
	serverKey string
}

func New(baseURL string, serverKey string) *Nakama {
	return &Nakama{
		client:    &http.Client{},
		baseURL:   baseURL,
		serverKey: serverKey,
	}
}
