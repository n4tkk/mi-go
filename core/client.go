package core

import (
	"io"
	"net/http"
)

type Client struct {
	BaseURL    string
	Token      string
	HttpClient *http.Client
}

func NewClient(baseUrl string, token string) *Client {
	return &Client{
		BaseURL:    baseUrl,
		Token:      token,
		HttpClient: new(http.Client),
	}
}

// SendRequest は指定したエンドポイントに対してリクエストを送信します。
func (c *Client) SendRequest(endpoint string) ([]byte, error) {
	if req, err := http.NewRequest(http.MethodPost, c.BaseURL+endpoint, nil); err != nil {
		return nil, err
	} else {

		req.Header.Set("Content-Type", "application/json")

		if res, err := c.HttpClient.Do(req); err != nil {
			return nil, err
		} else {
			defer res.Body.Close()
			return io.ReadAll(res.Body)
		}
	}
}
