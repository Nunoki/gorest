package gorest

import "time"

type Response struct {
	Payload interface{}  `json:"payload"`
	Meta    ResponseMeta `json:"meta"`
}

type ResponseMeta struct {
	ModifiedAt customTime `json:"modifiedAt"`
}

type RespError struct {
	ErrorMessage string `json:"error_message"`
}

type customTime struct {
	time.Time
}
