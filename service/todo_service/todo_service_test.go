package todo_service

import (
	"assignment-4/domain/todo_domain"
	"assignment-4/utils/error_utils"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	createTodo     func(todo *todo_domain.Todo) (*todo_domain.Todo, error_utils.MessageErr)
	updateTodo     func(todo *todo_domain.Todo) (*todo_domain.Todo, error_utils.MessageErr)
	getTodoById    func(todoId int64) (*todo_domain.Todo, error_utils.MessageErr)
	getAllTodos    func() (*[]todo_domain.Todo, error_utils.MessageErr)
	deleteTodoById func(todoId int64) (*map[string]interface{}, error_utils.MessageErr)
)

type todoDomainMock struct{}

func (t *todoDomainMock) CreateTodo(todo *todo_domain.Todo) (*todo_domain.Todo, error_utils.MessageErr) {
	return createTodo(todo)
}

func (t *todoDomainMock) UpdateTodo(todo *todo_domain.Todo) (*todo_domain.Todo, error_utils.MessageErr) {
	return updateTodo(todo)
}

func (t *todoDomainMock) GetTodoById(todoId int64) (*todo_domain.Todo, error_utils.MessageErr) {
	return getTodoById(todoId)
}

func (t *todoDomainMock) GetAllTodos() (*[]todo_domain.Todo, error_utils.MessageErr) {
	return getAllTodos()
}

func (t *todoDomainMock) DeleteTodoById(todoId int64) (*map[string]interface{}, error_utils.MessageErr) {
	return deleteTodoById(todoId)
}

// ----------------
// Test Create Todo

func TestTodoService_CreateTodo_Success(t *testing.T) {
	todo_domain.TodoDomain = &todoDomainMock{}

	requestBody := &todo_domain.Todo{
		Title:       "Homework",
		Description: "Deadline: January 19, 2022",
		Completed:   false,
	}

	expectedVal := &todo_domain.Todo{
		Id:          1,
		Title:       "Homework",
		Description: "Deadline: January 19, 2022",
		Completed:   false,
	}

	createTodo = func(todo *todo_domain.Todo) (*todo_domain.Todo, error_utils.MessageErr) {
		return expectedVal, nil
	}

	todo, err := TodoService.CreateTodo(requestBody)

	assert.NotNil(t, todo)
	assert.Nil(t, err)
}

func TestTodoService_CreateTodo_ServerError(t *testing.T) {
	todo_domain.TodoDomain = &todoDomainMock{}

	createTodo = func(todo *todo_domain.Todo) (*todo_domain.Todo, error_utils.MessageErr) {
		return nil, error_utils.NewInternalServerError("something went wrong")
	}

	requestBody := &todo_domain.Todo{
		Title:       "Homework",
		Description: "Deadline: January 19, 2022",
		Completed:   false,
	}

	todo, err := TodoService.CreateTodo(requestBody)

	assert.NotNil(t, err)
	assert.Nil(t, todo)
	assert.EqualValues(t, "server-error", err.Error())
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "something went wrong", err.Message())
}

func TestTodoService_CreateTodo_BadRequest(t *testing.T) {
	todo_domain.TodoDomain = &todoDomainMock{}

	tests := []struct {
		name    string
		todoReq *todo_domain.Todo
		errMsg  string
		status  int
		err     string
	}{
		{
			name: "empty title error",
			todoReq: &todo_domain.Todo{
				Title:       "",
				Description: "Deadline: January 19, 2022",
			},
			errMsg: "title is required",
			status: http.StatusBadRequest,
			err:    "bad_request",
		},
		{
			name: "empty image_url error",
			todoReq: &todo_domain.Todo{
				Title:       "Homework",
				Description: "",
			},
			errMsg: "description is required",
			status: http.StatusBadRequest,
			err:    "bad_request",
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			todo, err := TodoService.CreateTodo(tt.todoReq)

			assert.NotNil(t, err)
			assert.Nil(t, todo)
			assert.EqualValues(t, tt.err, err.Error())
			assert.EqualValues(t, tt.errMsg, err.Message())
			assert.EqualValues(t, tt.status, err.Status())
		})
	}
}

// ----------------
// Test Update Todo

func TestTodoService_UpdateTodo_Success(t *testing.T) {
	todo_domain.TodoDomain = &todoDomainMock{}

	requestBody := &todo_domain.Todo{
		Id:          1,
		Title:       "Homework",
		Description: "Deadline: January 19, 2022",
		Completed:   false,
	}

	expectedVal := &todo_domain.Todo{
		Id:          1,
		Title:       "Homework",
		Description: "Deadline: January 19, 2022",
		Completed:   false,
	}

	updateTodo = func(todo *todo_domain.Todo) (*todo_domain.Todo, error_utils.MessageErr) {
		return expectedVal, nil
	}

	todo, err := TodoService.UpdateTodo(requestBody)

	assert.NotNil(t, todo)
	assert.Nil(t, err)
}

func TestTodoService_UpdateTodo_ServerError(t *testing.T) {
	todo_domain.TodoDomain = &todoDomainMock{}

	updateTodo = func(todo *todo_domain.Todo) (*todo_domain.Todo, error_utils.MessageErr) {
		return nil, error_utils.NewInternalServerError("something went wrong")
	}

	requestBody := &todo_domain.Todo{
		Title:       "Homework",
		Description: "Deadline: January 19, 2022",
		Completed:   false,
	}

	todo, err := TodoService.UpdateTodo(requestBody)

	assert.NotNil(t, err)
	assert.Nil(t, todo)
	assert.EqualValues(t, "server-error", err.Error())
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "something went wrong", err.Message())
}

func TestTodoService_UpdateTodo_BadRequest(t *testing.T) {
	todo_domain.TodoDomain = &todoDomainMock{}

	tests := []struct {
		name    string
		todoReq *todo_domain.Todo
		errMsg  string
		status  int
		err     string
	}{
		{
			name: "empty title error",
			todoReq: &todo_domain.Todo{
				Title:       "",
				Description: "Deadline: January 19, 2022",
			},
			errMsg: "title is required",
			status: http.StatusBadRequest,
			err:    "bad_request",
		},
		{
			name: "empty image_url error",
			todoReq: &todo_domain.Todo{
				Title:       "Homework",
				Description: "",
			},
			errMsg: "description is required",
			status: http.StatusBadRequest,
			err:    "bad_request",
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			todo, err := TodoService.UpdateTodo(tt.todoReq)

			assert.NotNil(t, err)
			assert.Nil(t, todo)
			assert.EqualValues(t, tt.err, err.Error())
			assert.EqualValues(t, tt.errMsg, err.Message())
			assert.EqualValues(t, tt.status, err.Status())
		})
	}
}

// ----------------
// Test Get Todo By ID

func TestTodoService_GetTodoById_Success(t *testing.T) {
	todo_domain.TodoDomain = &todoDomainMock{}

	expectedVal := &todo_domain.Todo{
		Id:          1,
		Title:       "Homework",
		Description: "Deadline: January 19, 2022",
		Completed:   false,
	}

	getTodoById = func(todoId int64) (*todo_domain.Todo, error_utils.MessageErr) {
		return expectedVal, nil
	}

	todo, err := TodoService.GetTodoById(1)

	assert.Nil(t, err)
	assert.NotNil(t, todo)

	assert.EqualValues(t, expectedVal, todo)
}

func TestTodoService_GetTodoById_NotFoundError(t *testing.T) {
	todo_domain.TodoDomain = &todoDomainMock{}

	getTodoById = func(todoId int64) (*todo_domain.Todo, error_utils.MessageErr) {
		return nil, error_utils.NewNotFoundError("data not found")
	}

	todo, err := TodoService.GetTodoById(1)

	assert.NotNil(t, err)
	assert.Nil(t, todo)

	assert.EqualValues(t, error_utils.NewNotFoundError("data not found"), err)
}

// ----------------
// Test Get All Todos

func TestTodoService_GetAllTodos_Success(t *testing.T) {
	todo_domain.TodoDomain = &todoDomainMock{}

	expectedVal := &[]todo_domain.Todo{
		{
			Id:          1,
			Title:       "Homework",
			Description: "Deadline: January 19, 2022",
			Completed:   false,
		},
		{
			Id:          2,
			Title:       "Hacktiv8 Course Last Session",
			Description: "Using Gmeet at Wednesday, January 12 2022",
			Completed:   true,
		},
	}

	getAllTodos = func() (*[]todo_domain.Todo, error_utils.MessageErr) {
		return expectedVal, nil
	}

	todo, err := TodoService.GetAllTodos()

	assert.Nil(t, err)
	assert.NotNil(t, todo)

	assert.EqualValues(t, expectedVal, todo)
}

// ----------------
// Test Delete Todo By ID

func TestTodoService_DeleteTodoById_Success(t *testing.T) {
	todo_domain.TodoDomain = &todoDomainMock{}

	expectedVal := &map[string]interface{}{
		"StatusDelete": "Success",
		"AffectedRow":  1,
	}

	deleteTodoById = func(todoId int64) (*map[string]interface{}, error_utils.MessageErr) {
		return expectedVal, nil
	}

	todo, err := TodoService.DeleteTodoById(1)

	assert.Nil(t, err)
	assert.NotNil(t, todo)

	assert.EqualValues(t, expectedVal, todo)
}

func TestTodoService_DeleteTodoById_NotFoundError(t *testing.T) {

	deleteTodoById = func(todoId int64) (*map[string]interface{}, error_utils.MessageErr) {
		return nil, error_utils.NewNotFoundError("data not found")
	}

	todo, err := TodoService.DeleteTodoById(1)

	assert.NotNil(t, err)
	assert.Nil(t, todo)

	assert.EqualValues(t, error_utils.NewNotFoundError("data not found"), err)
}
