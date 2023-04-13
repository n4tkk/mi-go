package admin

import (
	"encoding/json"
	"fmt"
	"github.com/n4tkk/mi-go/model/entity"
)

// GetServerInfo はサーバーの情報を取得します。
func (s *Service) GetServerInfo() (*entity.ServerInfo, error) {
	endpoint := fmt.Sprintf("admin/server-info")
	r, _ := s.Client.SendRequest(endpoint)

	serverInfo := struct {
		Ok         bool `json:"ok"`
		ServerInfo entity.ServerInfo
		Error      interface{} `json:"error"`
	}{}
	if err := json.Unmarshal(r, &serverInfo); err != nil {
		return nil, err
	} else {
		if serverInfo.Ok {
			return &serverInfo.ServerInfo, nil
		} else {
			return nil, fmt.Errorf("server-info error: %v", serverInfo.Error)
		}
	}
}
