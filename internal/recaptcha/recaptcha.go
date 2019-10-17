package recaptcha

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// https://developers.google.com/recaptcha/docs/v3?authuser=3
const verifyAPI = "https://www.google.com/recaptcha/api/siteverify"

type Client struct {
	SecretKey string
}

type Response struct {
	Success     bool      `json:"success"`
	Score       float64   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes,omitempty"`
}

func (r *Response) String() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *Client) Verify(ctx context.Context, response string) (bool, error) {
	data := url.Values{
		"secret":   {r.SecretKey},
		"response": {response},
	}
	body := strings.NewReader(data.Encode())
	req, err := http.NewRequest("POST", verifyAPI, body)
	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var rr Response
	if err := json.NewDecoder(resp.Body).Decode(&rr); err != nil {
		return false, nil
	}

	log.Printf("recapture response: %s\n", rr.String())

	return rr.Success, nil
}
