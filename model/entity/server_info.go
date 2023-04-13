package entity

type ServerInfo struct {
	Machine string      `json:"machine"`
	Os      string      `json:"os"`
	Node    string      `json:"node"`
	Psql    string      `json:"psql"`
	Cpu     interface{} `json:"cpu"`
	Fs      interface{} `json:"fs"`
	Net     interface{} `json:"net"`
}
