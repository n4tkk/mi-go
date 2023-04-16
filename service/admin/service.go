package admin

import (
	"github.com/n4tkk/mi-go/core"
)

type Service struct {
	Client *core.Client
}

func NewService(client *core.Client) *Service {
	return &Service{Client: client}
}
