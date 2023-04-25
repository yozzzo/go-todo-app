package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// NewMiddleware middlewareの設定を行います.
func NewMiddleware(e *echo.Echo) {

	// echoが用意するmiddleware。
	// https://echo.labstack.com/middleware/

	// http headerやquery、formのデータをログに出力できる。
	e.Use(middleware.Logger())
	// panicを起こしてもサーバーが落ちないようにし、エラーレスポンスを返すようにする。
	e.Use(middleware.Recover())
	// CORSWithConfig CORSの設定。
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
}
