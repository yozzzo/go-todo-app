package handler

// AppHandler interfase
// AppHandlerは全てのHandlerのinterfaceを満たす.※routerの実装が依存する.
type AppHandler interface {
	UserHandler
	GreetHandler
	TodoHandler
	// embed all handler interfaces
}
