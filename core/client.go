package core

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
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
func (c *Client) SendRequest(endpoint string, body interface{}) ([]byte, error) {
	baseUrl, _ := url.Parse(c.BaseURL)
	if req, err := http.NewRequest(http.MethodPost, path.Join(baseUrl.Path, endpoint), nil); err != nil {
		return nil, err
	} else {
		if c.Token != "" {
			// リクエストヘッダーに Content-Type を設定する
			req.Header.Set("Content-Type", "application/json")

			// req.Body に SetToken で作成したリクエストボディを設定する
			if r, err := json.Marshal(SetToken(c.Token, body)); err != nil {
				return nil, err
			} else {
				req.Body = io.NopCloser(strings.NewReader(string(r)))
			}
		}
		if res, err := c.HttpClient.Do(req); err != nil {
			return nil, err
		} else {
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					fmt.Println(err)
				}
			}(res.Body)
			return io.ReadAll(res.Body)
		}
	}
}

func SetToken(token string, body interface{}) map[string]interface{} {
	var reqBody map[string]interface{}
	if body == nil {
		reqBody = make(map[string]interface{})
	} else {
		r, _ := json.Marshal(body)
		if err := json.Unmarshal(r, &reqBody); err != nil {
			fmt.Println(err)
		}
	}

	reqBody["i"] = token
	return reqBody
}
