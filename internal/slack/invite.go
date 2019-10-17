package slack

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	token     string
	inviteAPI string
}

func NewInviteClient(workspace, token string) *Client {
	return &Client{
		token:     token,
		inviteAPI: "https://" + workspace + ".slack.com/api/users.admin.invite",
	}
}

type Response struct {
	OK       bool   `json:"ok"`
	Error    string `json:"error,omitempty"`
	Needed   string `json:"needed,omitempty"`
	Provided string `json:"provided,omitempty"`
}

func (c *Client) Invite(ctx context.Context, email string) (*Response, error) {
	data := url.Values{
		"email":      {email},
		"token":      {c.token},
		"set_active": {"true"},
	}
	body := strings.NewReader(data.Encode())

	req, err := http.NewRequest("POST", c.inviteAPI, body)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var ir Response
	if err := json.NewDecoder(resp.Body).Decode(&ir); err != nil {
		return nil, err
	}
	return &ir, nil
}
