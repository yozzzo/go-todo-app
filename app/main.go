package main

import (
	"flag"
	"fmt"

	_ "github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/yozzzo/go-todo-app/conf"
	"github.com/yozzzo/go-todo-app/interactor"
	"github.com/yozzzo/go-todo-app/presenter/http/middleware"
	"github.com/yozzzo/go-todo-app/presenter/http/router"
)

// Dockerコンテナで実行する時(production.confをもとに起動するとき)は起動時に-serverを指定
var runServer = flag.Bool("server", false, "production is -server option require")

func main() {
	flag.Parse()
	conf.NewConfig(*runServer)

	e := echo.New()
	conn := conf.NewDBConnection()
	defer func() {
		if err := conn.Close(); err != nil {
			e.Logger.Fatal(fmt.Sprintf("Failed to close: %v", err))
		}
	}()
	i := interactor.NewInteractor(conn)
	h := i.NewAppHandler()

	router.NewRouter(e, h)
	middleware.NewMiddleware(e)
	if err := e.Start(fmt.Sprintf(":%d", conf.Current.Server.Port)); err != nil {
		e.Logger.Fatal(fmt.Sprintf("Failed to start: %v", err))
	}
}
