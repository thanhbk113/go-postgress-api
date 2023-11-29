package gin

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinCustomCtx struct {
	*gin.Context
}

// GinGetCustomCtx ...
func GinGetCustomCtx(c *gin.Context) *GinCustomCtx {
	return &GinCustomCtx{c}
}

// GetRequestCtx get request context
func (c *GinCustomCtx) GetRequestCtx() context.Context {
	return c.Request.Context()
}

// Response ...
type Response struct {
	HTTPCode int         `json:"-"`
	Data     interface{} `json:"data"`
	Message  string      `json:"message"`
}

func sendResponse(c *GinCustomCtx, data Response) {
	c.JSON(data.HTTPCode, data)
}

func getResponse(data interface{}, messageKey string, httpCode int) Response {
	if messageKey == "" {
		messageKey = "success"
	}

	return Response{
		HTTPCode: httpCode,
		Data:     data,
		Message:  messageKey,
	}

}

func (c *GinCustomCtx) Response200(data interface{}, msgKey string) error {

	resp := getResponse(data, msgKey, http.StatusOK)
	sendResponse(c, resp)
	return nil
}

// Response400 bad request
func (c *GinCustomCtx) Response400(data interface{}, msgKey string) error {

	resp := getResponse(data, msgKey, http.StatusBadRequest)
	sendResponse(c, resp)
	return nil
}
