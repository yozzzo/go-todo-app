package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/yozzzo/go-todo-app/domain/model"
	"github.com/yozzzo/go-todo-app/usecase"
)

// UserHandler interface
type TodoHandler interface {
	GetHello(c echo.Context) error

	GetTodos(c echo.Context) error
	GetTodo(c echo.Context) error
	CreateTodo(c echo.Context) error
	UpdateTodo(c echo.Context) error
	DeleteTodo(c echo.Context) error
}

type todoHandler struct {
	TodoUseCase usecase.TodoUseCase
}

type userHandler struct {
	UserUseCase usecase.UserUseCase
}

func NewTodoHandler(u usecase.TodoUseCase) TodoHandler {
	return &todoHandler{u}
}

// NewUserHandler UserHandlerを取得します.
// func NewUserHandler(u usecase.UserUseCase) UserHandler {
// 	return &userHandler{u}
// }

func (h *todoHandler) GetHello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func (h *userHandler) GetUsers(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	users, err := h.UserUseCase.GetUsers(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Users does not exist.")
	}

	return c.JSON(http.StatusOK, users)
}

func (h *userHandler) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "User ID must be int")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	user, err := h.UserUseCase.GetUser(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User does not exist.")
	}

	return c.JSON(http.StatusOK, user)
}

func (h *userHandler) CreateUser(c echo.Context) error {
	user := &model.User{}
	if err := c.Bind(user); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	user, err := h.UserUseCase.CreateUser(ctx, user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "User can not Create.")
	}

	return c.JSON(http.StatusCreated, user)
}

func (h *userHandler) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "User ID must be int")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	user, err := h.UserUseCase.UpdateUser(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "User can not Create.")
	}

	return c.JSON(http.StatusOK, user)
}

func (h *userHandler) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "User ID must be int")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if err := h.UserUseCase.DeleteUser(ctx, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "User can not Delete.")
	}

	return c.NoContent(http.StatusNoContent)
}
