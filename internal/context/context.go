package context

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Context represents context of a request.
type Context struct {
	Resp http.ResponseWriter
	Req  *http.Request
}

// NewContext creates new context.
func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Resp: w,
		Req:  req,
	}
}

// Text renders content as Text.
func (ctx *Context) Text(status int, response string) {
	ctx.Resp.WriteHeader(status)
	ctx.Resp.Write([]byte(response))
}

// JSON renders content as JSON.
func (ctx *Context) JSON(status int, content interface{}) {
	ctx.Resp.Header().Set("Content-Type", "application/json;charset=utf-8")
	ctx.Resp.WriteHeader(status)

	if err := json.NewEncoder(ctx.Resp).Encode(content); err != nil {
		ctx.Error(http.StatusInternalServerError, fmt.Errorf("render JSON failed: %v", err))
	}
}

// APIError is error format response
// swagger:model APIError
type APIError struct {
	Message string `json:"message"`
}

// APIValidationError is error format response related to input validation.
// swagger:model APIValidationError
type APIValidationError struct {
	Message string `json:"message"`
}

// APINotFound is a not found response
// swagger:model APINotFound
type APINotFound struct {
	Message string `json:"message"`
}

// Error responds with an error message to client with given obj as the message.
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
