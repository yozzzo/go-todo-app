package router

import (
	"github.com/labstack/echo"
	"github.com/yozzzo/go-todo-app/presenter/http/handler"
)

// NewRouter Routerの設定を行います.
func NewRouter(e *echo.Echo, h handler.AppHandler, g handler.GreetHandler) {
	e.GET("/hello", g.GetHello)

	e.GET("/todos", h.GetTodos)
	e.GET("/todos/:id", h.GetTodo)
	e.POST("/todos", h.CreateTodo)
	e.PUT("/todos/:id", h.UpdateTodo)
	e.DELETE("/todos/:id", h.DeleteTodo)
}
