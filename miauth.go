package misskey

import (
	"github.com/n4tkk/mi-go/core"
	"github.com/n4tkk/mi-go/service/miauth"
	"net/url"
)

// GetToken は misskey_host から sessionId でトークンを取得する
func GetToken(misskeyHost string, sessionId string) (string, error) {
	u, _ := url.Parse(misskeyHost)
	u.Path = "/api/"
	client := core.NewClient(u.String(), "")

	s := miauth.NewService(client)
	return s.GetToken(sessionId)
}
