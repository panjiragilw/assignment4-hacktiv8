package todo_controller

import (
	"assignment-4/domain/todo_domain"
	"assignment-4/service/todo_service"
	"assignment-4/utils/error_utils"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	createTodo     func(todo *todo_domain.Todo) (*todo_domain.Todo, error_utils.MessageErr)
	updateTodo     func(todo *todo_domain.Todo) (*todo_domain.Todo, error_utils.MessageErr)
	getTodoById    func(todoId int64) (*todo_domain.Todo, error_utils.MessageErr)
	getAllTodos    func() (*[]todo_domain.Todo, error_utils.MessageErr)
	deleteTodoById func(todoId int64) (*map[string]interface{}, error_utils.MessageErr)
)

type todoServiceMock struct{}

func (t *todoServiceMock) CreateTodo(todo *todo_domain.Todo) (*todo_domain.Todo, error_utils.MessageErr) {
	return createTodo(todo)
}

func (t *todoServiceMock) UpdateTodo(todo *todo_domain.Todo) (*todo_domain.Todo, error_utils.MessageErr) {
	return updateTodo(todo)
}

func (t *todoServiceMock) GetTodoById(todoId int64) (*todo_domain.Todo, error_utils.MessageErr) {
	return getTodoById(todoId)
}

func (t *todoServiceMock) GetAllTodos() (*[]todo_domain.Todo, error_utils.MessageErr) {
	return getAllTodos()
}

func (t *todoServiceMock) DeleteTodoById(todoId int64) (*map[string]interface{}, error_utils.MessageErr) {
	return deleteTodoById(todoId)
}

// ----------------
// Test Create Todo

func TestTodoController_CreateTodo_Success(t *testing.T) {
	todo_service.TodoService = &todoServiceMock{}

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

	requestJsonData, _ := json.Marshal(requestBody)

	createTodo = func(todoReq *todo_domain.Todo) (*todo_domain.Todo, error_utils.MessageErr) {
		return expectedVal, nil
	}

	r := gin.Default()

	req, _ := http.NewRequest(http.MethodPost, "/todo", bytes.NewBufferString(string(requestJsonData)))
	rr := httptest.NewRecorder()

	r.POST("/todo", CreateTodo)

	r.ServeHTTP(rr, req)

	result := rr.Result()

	data, _ := ioutil.ReadAll(result.Body)
	defer result.Body.Close()

	var todo todo_domain.Todo

	err := json.Unmarshal(data, &todo)

	assert.NotNil(t, todo)
	assert.Nil(t, err)
}

func TestTodoService_CreateTodo_ServerError(t *testing.T) {
	todo_service.TodoService = &todoServiceMock{}

	requestBody := &todo_domain.Todo{
		Title:       "Homework",
		Description: "Deadline: January 19, 2022",
		Completed:   false,
	}

	requestJsonData, _ := json.Marshal(requestBody)

	createTodo = func(todo *todo_domain.Todo) (*todo_domain.Todo, error_utils.MessageErr) {
		return nil, error_utils.NewInternalServerError("something went wrong")
	}

	r := gin.Default()

	req, _ := http.NewRequest(http.MethodPost, "/todo", bytes.NewBufferString(string(requestJsonData)))
	rr := httptest.NewRecorder()

	r.POST("/todo", CreateTodo)

	r.ServeHTTP(rr, req)

	result := rr.Result()

	data, _ := ioutil.ReadAll(result.Body)
	defer result.Body.Close()

	var errData error_utils.MessageErrData
	var errDataInterface error_utils.MessageErr = &errData

	err := json.Unmarshal(data, &errData)

	require.Nil(t, err)
	assert.NotNil(t, errDataInterface)
	assert.EqualValues(t, http.StatusInternalServerError, errDataInterface.Status())
	assert.EqualValues(t, "server-error", errDataInterface.Error())
	assert.EqualValues(t, "something went wrong", errDataInterface.Message())
}

func TestTodoService_CreateTodo_BadRequest(t *testing.T) {
	todo_service.TodoService = &todoServiceMock{}

	requestBody := []struct {
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

	requestJsonData, _ := json.Marshal(requestBody)

	r := gin.Default()
	req, _ := http.NewRequest(http.MethodPost, "/todo", bytes.NewBufferString(string(requestJsonData)))
	rr := httptest.NewRecorder()
	r.POST("/todo", CreateTodo)
	r.ServeHTTP(rr, req)
	result := rr.Result()
	data, _ := ioutil.ReadAll(result.Body)
	defer result.Body.Close()

	for _, tt := range requestBody {

		var errData error_utils.MessageErrData
		var errDataInterface error_utils.MessageErr = &errData
		t.Run(tt.name, func(t *testing.T) {
			err := json.Unmarshal(data, &errData)

			require.Nil(t, err)
			assert.NotNil(t, errDataInterface)
			assert.EqualValues(t, http.StatusBadRequest, errDataInterface.Status())
			assert.EqualValues(t, "bad_request", errDataInterface.Error())
			assert.EqualValues(t, "invalid json body", errDataInterface.Message())
		})
	}
}

// Test Update Todo

func TestTodoService_UpdateTodo_Success(t *testing.T) {
	todo_service.TodoService = &todoServiceMock{}

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

	requestJsonData, _ := json.Marshal(requestBody)

	updateTodo = func(todo *todo_domain.Todo) (*todo_domain.Todo, error_utils.MessageErr) {
		return expectedVal, nil
	}

	r := gin.Default()

	todoId := "1"

	req, _ := http.NewRequest(http.MethodPut, "/todo/"+todoId, bytes.NewBufferString(string(requestJsonData)))
	rr := httptest.NewRecorder()

	r.PUT("/todo/:todoId", UpdateTodo)

	r.ServeHTTP(rr, req)

	result := rr.Result()

	data, _ := ioutil.ReadAll(result.Body)
	defer result.Body.Close()

	var todo todo_domain.Todo

	err := json.Unmarshal(data, &todo)

	assert.NotNil(t, todo)
	assert.Nil(t, err)
}

func TestTodoService_UpdateTodo_ServerError(t *testing.T) {
	todo_service.TodoService = &todoServiceMock{}

	requestBody := &todo_domain.Todo{
		Title:       "Homework",
		Description: "Deadline: January 19, 2022",
		Completed:   false,
	}

	requestJsonData, _ := json.Marshal(requestBody)

	updateTodo = func(todo *todo_domain.Todo) (*todo_domain.Todo, error_utils.MessageErr) {
		return nil, error_utils.NewInternalServerError("something went wrong")
	}

	r := gin.Default()
	todoId := "1"

	req, _ := http.NewRequest(http.MethodPut, "/todo/"+todoId, bytes.NewBufferString(string(requestJsonData)))
	rr := httptest.NewRecorder()

	r.PUT("/todo/:todoId", UpdateTodo)

	r.ServeHTTP(rr, req)

	result := rr.Result()

	data, _ := ioutil.ReadAll(result.Body)
	defer result.Body.Close()

	var errData error_utils.MessageErrData
	var errDataInterface error_utils.MessageErr = &errData

	err := json.Unmarshal(data, &errData)

	require.Nil(t, err)
	assert.NotNil(t, errDataInterface)
	assert.EqualValues(t, http.StatusInternalServerError, errDataInterface.Status())
	assert.EqualValues(t, "server-error", errDataInterface.Error())
	assert.EqualValues(t, "something went wrong", errDataInterface.Message())
}

func TestTodoService_UpdateTodo_BadRequest(t *testing.T) {
	todo_service.TodoService = &todoServiceMock{}

	requestBody := []struct {
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

	requestJsonData, _ := json.Marshal(requestBody)

	r := gin.Default()

	todoId := "1"

	req, _ := http.NewRequest(http.MethodPut, "/todo/"+todoId, bytes.NewBufferString(string(requestJsonData)))
	rr := httptest.NewRecorder()

	r.PUT("/todo/:todoId", UpdateTodo)

	r.ServeHTTP(rr, req)

	result := rr.Result()
	data, _ := ioutil.ReadAll(result.Body)
	defer result.Body.Close()

	for _, tt := range requestBody {

		var errData error_utils.MessageErrData
		var errDataInterface error_utils.MessageErr = &errData
		t.Run(tt.name, func(t *testing.T) {
			err := json.Unmarshal(data, &errData)

			require.Nil(t, err)
			assert.NotNil(t, errDataInterface)
			assert.EqualValues(t, http.StatusBadRequest, errDataInterface.Status())
			assert.EqualValues(t, "bad_request", errDataInterface.Error())
			assert.EqualValues(t, "invalid json body", errDataInterface.Message())
		})
	}
}

// ----------------
// Test Get Todo By ID

func TestTodoService_GetTodoById_Success(t *testing.T) {
	todo_service.TodoService = &todoServiceMock{}

	expectedVal := &todo_domain.Todo{
		Id:          1,
		Title:       "Homework",
		Description: "Deadline: January 19, 2022",
		Completed:   false,
	}

	getTodoById = func(todoId int64) (*todo_domain.Todo, error_utils.MessageErr) {
		return expectedVal, nil
	}

	r := gin.Default()

	todoId := "1"

	req, _ := http.NewRequest(http.MethodGet, "/todo/"+todoId, nil)
	rr := httptest.NewRecorder()

	r.GET("/todo/:todoId", GetTodoById)

	r.ServeHTTP(rr, req)

	result := rr.Result()

	data, _ := ioutil.ReadAll(result.Body)
	defer result.Body.Close()

	var todo todo_domain.Todo

	err := json.Unmarshal(data, &todo)

	assert.NotNil(t, todo)
	assert.Nil(t, err)

	assert.EqualValues(t, *expectedVal, todo)

}

func TestTodoService_GetTodoById_NotFoundError(t *testing.T) {
	todo_service.TodoService = &todoServiceMock{}

	getTodoById = func(todoId int64) (*todo_domain.Todo, error_utils.MessageErr) {
		return nil, error_utils.NewNotFoundError("data not found")
	}

	r := gin.Default()

	todoId := "999"

	req, _ := http.NewRequest(http.MethodGet, "/todo/"+todoId, nil)
	rr := httptest.NewRecorder()

	r.GET("/todo/:todoId", GetTodoById)

	r.ServeHTTP(rr, req)

	result := rr.Result()

	data, _ := ioutil.ReadAll(result.Body)
	defer result.Body.Close()

	var errData error_utils.MessageErrData
	var errDataInterface error_utils.MessageErr = &errData

	err := json.Unmarshal(data, &errData)

	require.Nil(t, err)
	assert.NotNil(t, errDataInterface)
	assert.EqualValues(t, http.StatusNotFound, errDataInterface.Status())
	assert.EqualValues(t, "not_found", errDataInterface.Error())
	assert.EqualValues(t, "data not found", errDataInterface.Message())
}

// ----------------
// Test Get All Todos

func TestTodoService_GetAllTodos_Success(t *testing.T) {
	todo_service.TodoService = &todoServiceMock{}

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

	r := gin.Default()

	req, _ := http.NewRequest(http.MethodGet, "/todo", nil)
	rr := httptest.NewRecorder()

	r.GET("/todo", GetAllTodos)

	r.ServeHTTP(rr, req)

	result := rr.Result()

	data, _ := ioutil.ReadAll(result.Body)
	defer result.Body.Close()

	var todo []todo_domain.Todo

	err := json.Unmarshal(data, &todo)

	assert.NotNil(t, todo)
	assert.Nil(t, err)

	assert.EqualValues(t, *expectedVal, todo)
}

// ----------------
// Test Delete Todo By ID

func TestTodoService_DeleteTodoById_Success(t *testing.T) {
	todo_service.TodoService = &todoServiceMock{}

	expectedVal := &map[string]interface{}{
		"StatusDelete": "Success",
		"AffectedRow":  1,
	}

	deleteTodoById = func(todoId int64) (*map[string]interface{}, error_utils.MessageErr) {
		return expectedVal, nil
	}

	r := gin.Default()

	todoId := "1"

	req, _ := http.NewRequest(http.MethodDelete, "/todo/"+todoId, nil)
	rr := httptest.NewRecorder()

	r.DELETE("/todo/:todoId", DeleteTodoById)

	r.ServeHTTP(rr, req)

	result := rr.Result()

	data, _ := ioutil.ReadAll(result.Body)
	defer result.Body.Close()

	var todo map[string]interface{}
	err := json.Unmarshal(data, &todo)
	todo["AffectedRow"] = int(todo["AffectedRow"].(float64))

	assert.NotNil(t, todo)
	assert.Nil(t, err)

	assert.EqualValues(t, *expectedVal, todo)
}

func TestTodoService_DeleteTodoById_NotFoundError(t *testing.T) {

	deleteTodoById = func(todoId int64) (*map[string]interface{}, error_utils.MessageErr) {
		return nil, error_utils.NewNotFoundError("data not found")
	}

	r := gin.Default()

	todoId := "999"

	req, _ := http.NewRequest(http.MethodDelete, "/todo/"+todoId, nil)
	rr := httptest.NewRecorder()

	r.DELETE("/todo/:todoId", DeleteTodoById)

	r.ServeHTTP(rr, req)

	result := rr.Result()

	data, _ := ioutil.ReadAll(result.Body)
	defer result.Body.Close()

	var errData error_utils.MessageErrData
	var errDataInterface error_utils.MessageErr = &errData

	err := json.Unmarshal(data, &errData)

	require.Nil(t, err)
	assert.NotNil(t, errDataInterface)
	assert.EqualValues(t, http.StatusNotFound, errDataInterface.Status())
	assert.EqualValues(t, "not_found", errDataInterface.Error())
	assert.EqualValues(t, "data not found", errDataInterface.Message())
}
