package types

type ErrorSerializer struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type Response struct {
	Code int `json:"code"`
}

