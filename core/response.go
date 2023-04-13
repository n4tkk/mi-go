package core

type Response struct {
	Ok    bool        `json:"ok"`
	Error interface{} `json:"error"`
}
