package http

import "net/http"

type Message struct {
	Code    int         `json:"code" xml:"code" yaml:"code" `
	Message string      `json:"message" xml:"message" yaml:"message"`
	Data    interface{} `json:"data" xml:"data" yaml:"data"`
}

type Context interface {
	JSON(code int, obj interface{})
	AbortWithStatusJSON(code int, jsonObj interface{})
}

func Msg(data ...interface{}) Message {
	result := Message{
		Code:    http.StatusOK,
		Message: "ok",
		Data:    nil,
	}
	count := len(data)
	if count > 0 {
		result.Data = data[0]
	}
	if count > 1 {
		if message, ok := data[1].(string); ok && message != "" {
			result.Message = message
		}
	}
	if count > 2 {
		if code, ok := data[2].(int); ok && code > 0 {
			result.Code = code
		}
	}
	return result
}

func Ok(ctx Context, data ...interface{}) {
	var newData interface{}
	newMessage := ""
	count := len(data)
	if count > 0 {
		newData = data[0]
	}
	if count > 1 {
		if message, ok := data[1].(string); ok {
			newMessage = message
		}
	}
	ctx.JSON(http.StatusOK, Msg(newData, newMessage, http.StatusOK))
}

func Error(ctx Context, code int, err error, data ...interface{}) {
	var newData interface{}
	if len(data) > 0 {
		newData = data[0]
	}
	ctx.AbortWithStatusJSON(http.StatusOK, Msg(newData, err.Error(), code))
}

func Err(ctx Context, err error, data ...interface{}) {
	Error(ctx, http.StatusInternalServerError, err, data)
}

func ErrBadRequest(ctx Context, err error, data ...interface{}) {
	Error(ctx, http.StatusBadRequest, err, data)
}

func ErrUnauthorized(ctx Context, err error, data ...interface{}) {
	Error(ctx, http.StatusUnauthorized, err, data)
}

func ErrNotFound(ctx Context, err error, data ...interface{}) {
	Error(ctx, http.StatusNotFound, err, data)
}
