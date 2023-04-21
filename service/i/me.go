package i

import (
	"encoding/json"
	"fmt"
	"github.com/n4tkk/mi-go/model/entity"
)

// GetMe はユーザーの情報を取得します。
func (s *Service) GetMe() (entity.User, error) {
	endpoint := fmt.Sprintf("i")
	r, _ := s.Client.SendRequest(endpoint, nil)

	var me entity.User
	err := json.Unmarshal(r, &me)
	return me, err
}
