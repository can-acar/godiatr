package handler

import (
	"github.com/gin-gonic/gin"
	"godiatr/library/dto"
	"godiatr/library/jsonrpc"
	"time"
)

// TimeHandler returns server time in RFC3339 format.
type TimeHandler struct{}

func (h TimeHandler) Method() string { return "time.now" }

func (h TimeHandler) Handle(c *gin.Context, req jsonrpc.Request) (interface{}, *jsonrpc.Error) {
	return dto.Time{Time: time.Now().Format(time.RFC3339)}, nil
}
