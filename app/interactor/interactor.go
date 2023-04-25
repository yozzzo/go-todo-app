package interactor

import (
	"github.com/jinzhu/gorm"
	"github.com/yozzzo/go-todo-app/domain/repository"
	"github.com/yozzzo/go-todo-app/domain/service"
	"github.com/yozzzo/go-todo-app/infrastructure/persistence/datastore"
	"github.com/yozzzo/go-todo-app/presenter/http/handler"
	"github.com/yozzzo/go-todo-app/usecase"
)

// Interactor interfase Intractorは安易DIコンテナとしての役割を持つ.
type Interactor interface {

	// User関係を作成
	NewUserRepository() repository.UserRepository
	NewUserService() service.UserService
	NewUserUseCase() usecase.UserUseCase
	NewUserHandler() handler.UserHandler

	// Greet関係を作成
	NewGreetRepository() repository.GreetRepository
	NewGreetService() service.GreetService
	NewGreetUseCase() usecase.GreetUseCase
	NewGreetHandler() handler.GreetHandler

	// Todo関係を作成
	NewTodoRepository() repository.TodoRepository
	NewTodoService() service.TodoService
	NewTodoUseCase() usecase.TodoUseCase
	NewTodoHandler() handler.TodoHandler

	// AppHandler関係を作成
	NewAppHandler() handler.AppHandler
}

// これまだわからん
type interactor struct {
	Conn *gorm.DB
}

// NewInteractor intractorを取得します.
func NewInteractor(Conn *gorm.DB) Interactor {
	return &interactor{Conn}
}

// handler/app_handler.goに書かないのはなんで？
type appHandler struct {
	handler.UserHandler
	handler.GreetHandler
	handler.TodoHandler
	// embed all handler interfaces
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	appHandler := &appHandler{}
	appHandler.UserHandler = i.NewUserHandler()
	return appHandler
}

// Userに関する依存関係
func (i *interactor) NewUserRepository() repository.UserRepository {
	return datastore.NewUserRepository(i.Conn)
}

func (i *interactor) NewUserService() service.UserService {
	return service.NewUserService(i.NewUserRepository())
}

func (i *interactor) NewUserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(i.NewUserRepository())
}

func (i *interactor) NewUserHandler() handler.UserHandler {
	return handler.NewUserHandler(i.NewUserUseCase())
}

// Greetに関する依存関係
func (i *interactor) NewGreetRepository() repository.GreetRepository {
	return datastore.NewGreetRepository(i.Conn)
}

func (i *interactor) NewGreetService() service.GreetService {
	// serviceはrepositoryに依存する
	return service.NewGreetService(i.NewGreetRepository())
}

func (i *interactor) NewGreetUseCase() usecase.GreetUseCase {
	// usecaseはrepositoryに依存する
	return usecase.NewGreetUseCase(i.NewGreetRepository())
}

func (i *interactor) NewGreetHandler() handler.GreetHandler {
	// handlerはusecaseに依存する
	return handler.NewGreetHandler(i.NewGreetUseCase())
}

// Todoに関する依存関係
func (i *interactor) NewTodoRepository() repository.TodoRepository {
	return datastore.NewTodoRepository(i.Conn)
}

func (i *interactor) NewTodoService() service.TodoService {
	return service.NewTodoService(i.NewTodoRepository())
}

func (i *interactor) NewTodoUseCase() usecase.TodoUseCase {
	return usecase.NewTodoUseCase(i.NewTodoRepository())
}

func (i *interactor) NewTodoHandler() handler.TodoHandler {
	return handler.NewTodoHandler(i.NewTodoUseCase())
}
