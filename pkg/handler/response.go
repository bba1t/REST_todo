package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)

	// AbortWithStatusJSON(code int, message) — прерывает цепочки обработчиков, а в ответ записывает статус код и тело сообщения в формате json
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
