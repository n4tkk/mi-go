package miauth

import (
	"encoding/json"
	"fmt"
	"github.com/n4tkk/mi-go/model/entity"
	"net/url"
	"strings"
)

type App struct {
	Scheme     string
	Host       string
	Name       string
	Callback   string
	Permission Permission
}
type Permission []string

// String はPermissionを文字列に変換します。
func (p *Permission) String() string {
	return strings.Join(*p, ",")
}

// GenerateRequestUrl は認証リクエストURLを生成します。
func (a *App) GenerateRequestUrl(sessionId string) string {
	u := url.URL{
		Scheme: a.Scheme,
		Host:   a.Host,
		Path:   fmt.Sprintf("miauth/%s", sessionId),
	}
	q := u.Query()
	if a.Name != "" {
		q.Set("name", a.Name)
	}
	if a.Callback != "" {
		q.Set("callback", a.Callback)
	}
	if a.Permission != nil {
		q.Set("permission", a.Permission.String())
	}
	u.RawQuery = q.Encode()
	return u.String()
}

// GetToken は指定したセッションIDに対応するトークンを取得します。
func (s *Service) GetToken(sessionId string) (string, error) {
	endpoint := fmt.Sprintf("miauth/%s/check", sessionId)
	r, _ := s.Client.SendRequest(endpoint, nil)

	miauth := struct {
		Ok    bool        `json:"ok"`
		Token string      `json:"token"`
		User  entity.User `json:"user"`
		Error interface{} `json:"error"`
	}{}
	if err := json.Unmarshal(r, &miauth); err != nil {
		return "", err
	} else {
		if miauth.Ok {
			return miauth.Token, nil
		} else {
			return "", fmt.Errorf("miauth error: %v", miauth.Error)
		}
	}
}
