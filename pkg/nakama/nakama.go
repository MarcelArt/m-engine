package nakama

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
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

func (n *Nakama) AuthenticateEmail(email string, password string, username string) error {
	funcStackTrace := "Nakama.AuthenticateEmail"

	url := fmt.Sprintf("%s/v2/account/authenticate/email?create=true&username=%s", n.baseURL, username)

	body := AuthEmailAccount{
		Email:    email,
		Password: password,
	}

	reqBody, err := body.IntoReader()
	if err != nil {
		return fmt.Errorf("%s: failed parsing into reader: %w", funcStackTrace, err)
	}

	req, err := http.NewRequest(http.MethodPost, url, reqBody)
	if err != nil {
		return fmt.Errorf("%s: failed creating new request: %w", funcStackTrace, err)
	}

	req.Header.Set("Content-Type", "application/json")
	basicAuth := base64.StdEncoding.EncodeToString(fmt.Appendf(nil, "%s:", n.serverKey))
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", basicAuth))

	res, err := n.client.Do(req)
	if err != nil {
		return fmt.Errorf("%s: failed sending request: %w", funcStackTrace, err)
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		resBody, _ := io.ReadAll(res.Body)
		return fmt.Errorf("%s: unexpected status code: %d: %s", funcStackTrace, res.StatusCode, string(resBody))
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("%s: failed reading body: %w", funcStackTrace, err)
	}

	var authRes AuthResponse
	if err := json.Unmarshal(resBody, &authRes); err != nil {
		return fmt.Errorf("%s: failed unmarshal response: %w", funcStackTrace, err)
	}

	n.Session.accessToken = authRes.Token
	n.Session.refreshToken = authRes.RefreshToken

	return nil
}

func (n *Nakama) GetAccount() (Account, error) {
	var account Account
	funcStackTrace := "Nakama.GetAccount"

	url := fmt.Sprintf("%s/v2/account", n.baseURL)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return account, fmt.Errorf("%s: failed creating new request: %w", funcStackTrace, err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", n.Session.accessToken))

	res, err := n.client.Do(req)
	if err != nil {
		return account, fmt.Errorf("%s: failed sending request: %w", funcStackTrace, err)
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		resBody, _ := io.ReadAll(res.Body)
		return account, fmt.Errorf("%s: unexpected status code: %d: %s", funcStackTrace, res.StatusCode, string(resBody))
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return account, fmt.Errorf("%s: failed reading body: %w", funcStackTrace, err)
	}

	if err := json.Unmarshal(resBody, &account); err != nil {
		return account, fmt.Errorf("%s: failed unmarshal response: %w", funcStackTrace, err)
	}

	return account, nil
}
