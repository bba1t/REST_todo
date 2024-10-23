package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo"
)

// Хендлер во фреймворке gin, должен иметь в качестве параметра указатель на объект *gin.Context

func (h *Handler) signUp(c *gin.Context) {
	var input todo.User // json-данные от пользователей

	// gin.Context - BindJSON() десириализует тело http-запроса и помещает json данные в указанную структуру
	if err := c.BindJSON(&input); err != nil {
		// выдаст ошибку если json данные от пользователя придут не в соответствии с шаблоном, код 400
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// После того как распарсил и валидировал запрос, он передается ниже в бизнес логику(service), которая реализует регистрацию

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error()) // внутренняя ошибка на сервере, код 500
		return
	}

	// отправляет http ответ в формате json, причем превратит в json любой тип
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// получение логина и пароля
type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// вызываю новый метод создания токена
	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
