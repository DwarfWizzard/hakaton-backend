package service

type Response struct {
	Result any    `json:"resul"` //TODO
	Error  *Error `json:"error"`
}