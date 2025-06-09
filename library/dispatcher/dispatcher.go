package dispatcher

import (
	"github.com/gin-gonic/gin"
	"godiatr/library/jsonrpc"
)

// Handler defines JSON-RPC handler
// implementations should return a result object or a jsonrpc.Error
// when something goes wrong.
type Handler interface {
	Method() string
	Handle(*gin.Context, jsonrpc.Request) (interface{}, *jsonrpc.Error)
}

// Dispatcher routes JSON-RPC requests to registered handlers.
type Dispatcher struct {
	engine   *gin.Engine
	handlers map[string]Handler
}

// New creates a dispatcher using the provided gin engine.
func New(engine *gin.Engine) *Dispatcher {
	return &Dispatcher{engine: engine, handlers: make(map[string]Handler)}
}

// Register adds a handler to dispatcher. Method name is taken from handler.Method().
func (d *Dispatcher) Register(h Handler) { d.handlers[h.Method()] = h }

// Bind attaches the dispatcher to given route path as POST handler.
func (d *Dispatcher) Bind(path string) { d.engine.POST(path, d.dispatch) }

// dispatch parses JSON-RPC request and executes registered handler.
func (d *Dispatcher) dispatch(c *gin.Context) {
	var req jsonrpc.Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, jsonrpc.NewErrorResponse(nil, jsonrpc.ErrInvalidRequest, "invalid request", nil))
		return
	}

	if req.JSONRPC != jsonrpc.Version {
		c.JSON(200, jsonrpc.NewErrorResponse(req.ID, jsonrpc.ErrInvalidRequest, "invalid jsonrpc version", nil))
		return
	}

	h, ok := d.handlers[req.Method]
	if !ok {
		c.JSON(200, jsonrpc.NewErrorResponse(req.ID, jsonrpc.ErrMethodNotFound, "method not found", nil))
		return
	}

	result, jerr := h.Handle(c, req)
	if jerr != nil {
		c.JSON(200, jsonrpc.NewErrorResponse(req.ID, jerr.Code, jerr.Message, jerr.Data))
		return
	}

	c.JSON(200, jsonrpc.NewResultResponse(req.ID, result))
}
