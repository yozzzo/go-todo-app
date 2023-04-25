package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/yozzzo/go-todo-app/usecase"
)

// ハンドラーのインターフェースを定義
// なんの役割？
type GreetHandler interface {
	GetHello(c echo.Context) error
}

// ハンドラーの構造体を定義
// なんの役割？
type greetHandler struct {
	greetUseCase usecase.GreetUseCase
}

// greetHandler型にNewGreetHandler関数を定義
// usecaseからGreetUseCaseを受け取り、greetHandler型を返す
func NewGreetHandler(u usecase.GreetUseCase) GreetHandler {
	return &greetHandler{u}
}

// greetHandler型にGetHello関数を定義
func (h *greetHandler) GetHello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
