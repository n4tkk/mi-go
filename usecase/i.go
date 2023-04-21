package usecase

import (
	"github.com/n4tkk/mi-go/core"
	"github.com/n4tkk/mi-go/model/entity"
	"github.com/n4tkk/mi-go/service/i"
	"net/url"
)

// GetMe is a usecase.
func GetMe(misskeyHost string, token string) (entity.User, error) {
	u, _ := url.Parse(misskeyHost)
	u.Path = "/api/"
	client := core.NewClient(u.String(), token)

	s := i.NewService(client)

	return s.GetMe()
}
