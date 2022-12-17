package router

import (
	"todo-app/src/handler"

	"github.com/labstack/echo/v4"
)

func InitializeRouter(server *echo.Echo, handler *handler.Handler) {

	activity := server.Group("/activity-groups")
	{
		activity.PATCH("/:id", handler.UpdateActivity)
		activity.GET("", handler.GetActivityList)
		activity.GET("/:id", handler.GetActivity)
		activity.POST("", handler.CreateActivity)
		activity.DELETE("/:id", handler.DeleteActivity)
	}

	todo := server.Group("/todo-items")
	{
		todo.GET("", handler.GetTodoList)
		todo.POST("", handler.CreateTodo)
		todo.GET("/:id", handler.GetTodo)
		todo.PATCH("/:id", handler.UpdateTodo)
		todo.DELETE("/:id", handler.DeleteTodo)
	}
}
