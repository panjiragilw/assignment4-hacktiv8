package todo_service

import (
	"assignment-4/domain/todo_domain"
	"assignment-4/utils/error_utils"
)

var TodoService todoServiceInterface = &todoService{}

type todoServiceInterface interface {
	CreateTodo(*todo_domain.Todo) (*todo_domain.Todo, error_utils.MessageErr)
	UpdateTodo(*todo_domain.Todo) (*todo_domain.Todo, error_utils.MessageErr)
	GetTodoById(int64) (*todo_domain.Todo, error_utils.MessageErr)
	GetAllTodos() (*[]todo_domain.Todo, error_utils.MessageErr)
	DeleteTodoById(int64) (*map[string]interface{}, error_utils.MessageErr)
}

type todoService struct{}

func (t *todoService) CreateTodo(todoReq *todo_domain.Todo) (*todo_domain.Todo, error_utils.MessageErr) {
	err := todoReq.Validate()

	if err != nil {
		return nil, err
	}

	res, err := todo_domain.TodoDomain.CreateTodo(todoReq)

	if err != nil {
		return nil, err
	}
	return res, err
}

func (t *todoService) UpdateTodo(todoReq *todo_domain.Todo) (*todo_domain.Todo, error_utils.MessageErr) {
	err := todoReq.Validate()

	if err != nil {
		return nil, err
	}
	res, err := todo_domain.TodoDomain.UpdateTodo(todoReq)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (t *todoService) GetTodoById(todoId int64) (*todo_domain.Todo, error_utils.MessageErr) {
	res, err := todo_domain.TodoDomain.GetTodoById(todoId)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (t *todoService) GetAllTodos() (*[]todo_domain.Todo, error_utils.MessageErr) {
	res, err := todo_domain.TodoDomain.GetAllTodos()

	if err != nil {
		return nil, err
	}

	return res, err
}

func (t *todoService) DeleteTodoById(todoId int64) (*map[string]interface{}, error_utils.MessageErr) {
	res, err := todo_domain.TodoDomain.DeleteTodoById(todoId)

	if err != nil {
		return nil, err
	}

	return res, err
}
