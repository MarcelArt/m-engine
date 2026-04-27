package nakama

import (
	"fmt"
	"net/http"
)

func (n *Nakama) GetMatches(limit uint, authoritative bool, minSize uint, maxSize uint) ([]Match, error) {
	funcStackTrace := "Nakama.GetMatch"

	url := fmt.Sprintf("%s/v2/match?limit=%d&authoritative=%t&minSize=%d&maxSize=%d", n.baseURL, limit, authoritative, minSize, maxSize)

	res, err := Fetch[GetMatchResponse](n.client, http.MethodGet, url, nil, map[string]string{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", n.Session.accessToken),
	})
	if err != nil {
		return nil, fmt.Errorf("%s: failed fetching matches: %w", funcStackTrace, err)
	}

	return res.Matches, nil
}
