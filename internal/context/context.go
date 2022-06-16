package context

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct {
	Resp http.ResponseWriter
	Req  *http.Request
}

func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Resp: w,
		Req:  req,
	}
}

// JSON render content as JSON
func (ctx *Context) Text(status int, response string) {
	ctx.Resp.WriteHeader(status)
	ctx.Resp.Write([]byte(response))
}

// JSON render content as JSON
func (ctx *Context) JSON(status int, content interface{}) {
	ctx.Resp.Header().Set("Content-Type", "application/json;charset=utf-8")
	ctx.Resp.WriteHeader(status)
	if err := json.NewEncoder(ctx.Resp).Encode(content); err != nil {
		ctx.ServerError("Render JSON failed", err)
	}
}

// APIError is error format response
// swagger:model APIError
type APIError struct {
	Message string `json:"message"`
}

// APIValidationError is error format response related to input validation
// swagger:model APIValidationError
type APIValidationError struct {
	Message string `json:"message"`
}

// APINotFound is a not found empty response
// swagger:model APINotFound
type APINotFound struct {
	Message string `json:"message"`
}

// Error responds with an error message to client with given obj as the message.
// If status is 500, also it prints error to log.
func (ctx *Context) Error(status int, obj interface{}) {
	var message string
	if err, ok := obj.(error); ok {
		message = err.Error()
	} else {
		message = fmt.Sprintf("%s", obj)
	}

	ctx.JSON(status, APIError{
		Message: message,
	})
}

// ServerError displays a 500 (Internal Server Error) page and prints the given error, if any.
func (ctx *Context) ServerError(logMsg string, logErr error) {
	ctx.serverErrorInternal(logMsg, logErr)
}

func (ctx *Context) serverErrorInternal(logMsg string, logErr error) {
	ctx.Text(http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: %s -> %s", logMsg, logErr))
}
