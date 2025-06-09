package jsonrpc

import "encoding/json"

const Version = "2.0"

const (
	ErrParseError     = -32700
	ErrInvalidRequest = -32600
	ErrMethodNotFound = -32601
	ErrInvalidParams  = -32602
	ErrInternal       = -32603
)

type Request struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      interface{}     `json:"id"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
}

type Response struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id"`
	Result  interface{} `json:"result,omitempty"`
	Error   *Error      `json:"error,omitempty"`
}

type Error struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResultResponse(id interface{}, result interface{}) Response {
	return Response{
		JSONRPC: Version,
		ID:      id,
		Result:  result,
	}
}

func NewErrorResponse(id interface{}, code int, msg string, data interface{}) Response {
	return Response{
		JSONRPC: Version,
		ID:      id,
		Error:   &Error{Code: code, Message: msg, Data: data},
	}
}
