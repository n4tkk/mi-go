package miauth

import (
	"encoding/json"
	"fmt"
	"github.com/n4tkk/mi-go/model/entity"
)

// GetToken は指定したセッションIDに対応するトークンを取得します。
func (s *Service) GetToken(sessionId string) (string, error) {
	endpoint := fmt.Sprintf("miauth/%s/check", sessionId)
	r, _ := s.Client.SendRequest(endpoint)

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
