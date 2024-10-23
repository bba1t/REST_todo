package handler

import (
	"github.com/gin-gonic/gin"
	"todo/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

/*
Маршрутизатор (router) - это экземпляр, который принимает на вход адрес запроса (`http://localhost:8080/pizzas`)
и вызывает исполнителя для этого запроса. Исполнитель (handler) - это метод, который вызывается маршрутизатором.
*/

// Будет инициализировать все эндпоинты

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// Объявляю группы маршрутизаторов

	// Группа по авторизации, то есть с этим url будет работать не авторизированный пользователь
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	// И уже после авторизации пользователь получит адрес "/api", где каждый раз будет вызываться проверка токена
	api := router.Group("/api", h.userIdentity) // группа по работе с задачами
	{
		lists := api.Group("/lists") //группа по работе со списками
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group(":id/items") // группа для задач списков
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET("/:item_id", h.getItemById)
				items.PUT("/:item_id", h.updateItem)
				items.DELETE("/:item_id", h.deleteItem)
			}
		}
	}
	return router
}
