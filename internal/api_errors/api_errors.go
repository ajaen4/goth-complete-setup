package api_errors

import "net/http"

type ClientErr struct {
	HttpCode int               `json:"code"`
	Message  string            `json:"message"`
	LogMess  string            `json:"-"`
	Errors   map[string]string `json:"errors,omitempty"`
}

type InternalErr struct {
	HttpCode int    `json:"code"`
	Message  string `json:"message"`
}

func (err *ClientErr) Error() string {
	if err.LogMess != "" {
		return err.LogMess
	}
	return err.Message
}

var UnauthErr = ClientErr{
	HttpCode: http.StatusUnauthorized,
	Message:  "Unauthorized",
}
