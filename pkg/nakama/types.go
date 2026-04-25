package nakama

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type AuthEmailAccount struct {
	Email    string         `json:"email"`
	Password string         `json:"password"`
	Vars     map[string]any `json:"vars"`
}

type AuthEmail struct {
	Account  AuthEmailAccount `json:"account"`
	Create   bool             `json:"create"`
	Username string           `json:"username"`
}

func (a AuthEmailAccount) IntoReader() (*bytes.Reader, error) {
	b, err := json.Marshal(a)
	if err != nil {
		return nil, fmt.Errorf("AuthEmail.IntoReader: %w", err)
	}

	return bytes.NewReader(b), nil
}

type AuthResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatarUrl"`
}

type Account struct {
	Email    string `json:"email"`
	CustomID string `json:"customId"`
	User     User   `json:"user"`
}
