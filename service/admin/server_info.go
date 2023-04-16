package admin

import (
	"encoding/json"
	"fmt"
	"github.com/n4tkk/mi-go/model/entity"
)

// GetServerInfo はサーバーの情報を取得します。
func (s *Service) GetServerInfo() (entity.ServerInfo, error) {
	endpoint := fmt.Sprintf("admin/server-info")
	r, _ := s.Client.SendRequest(endpoint, nil)

	var serverInfo entity.ServerInfo
	err := json.Unmarshal(r, &serverInfo)
	return serverInfo, err
}
