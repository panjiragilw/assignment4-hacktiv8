package todo_controller

import (
	"assignment-4/domain/todo_domain"
	"assignment-4/service/todo_service"
	"assignment-4/utils/error_utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateTodo godoc
// @Summary Create a todo
// @Tags todo
// @Description creating a new todo
// @ID create-todo
// @Accept json
// @Produce json
// @Param RequestBody body doc_datas.CreateTodoRequest true "request body json"
// @Success 201 {object} doc_datas.CreateTodoResponse
// @Failure 400 {object} error_utils.MessageErrData
// @Failure 500 {object} error_utils.MessageErrData
// @Router /todo [post]
func CreateTodo(c *gin.Context) {
	var todo todo_domain.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		theErr := error_utils.NewBadRequest("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}

	res, err := todo_service.TodoService.CreateTodo(&todo)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, res)
}

// UpdateTodo godoc
// @Summary Update todo
// @Tags todo
// @Description Updating a todo by ID
// @ID update-todo
// @Accept json
// @Produce json
// @Param RequestBody body doc_datas.UpdateTodoRequest true "request body json"
// @Param todoId path int true "todo's todo id"
// @Success 200 {object} doc_datas.UpdateTodoResponse
// @Failure 400 {object} error_utils.MessageErrData
// @Failure 404 {object} error_utils.MessageErrData
// @Failure 500 {object} error_utils.MessageErrData
// @Router /todo/{todoId} [put]
func UpdateTodo(c *gin.Context) {
	var todo todo_domain.Todo

	todoId, err := todo.GetTodoIdParam(c)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	if err := c.ShouldBindJSON(&todo); err != nil {
		theErr := error_utils.NewBadRequest("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}

	todo.Id = todoId

	res, err := todo_service.TodoService.UpdateTodo(&todo)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetTodoById godoc
// @Summary Get todo by ID
// @Tags todo
// @Description Getting a todo by ID
// @ID get-todo
// @Accept json
// @Produce json
// @Param todoId path int true "todo's todo id"
// @Success 200 {object} doc_datas.GetTodoResponse
// @Failure 400 {object} error_utils.MessageErrData
// @Failure 404 {object} error_utils.MessageErrData
// @Failure 500 {object} error_utils.MessageErrData
// @Router /todo/{todoId} [get]
func GetTodoById(c *gin.Context) {
	var todo todo_domain.Todo

	todoId, err := todo.GetTodoIdParam(c)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	res, err := todo_service.TodoService.GetTodoById(todoId)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetAllTodos godoc
// @Summary Get all todos
// @Tags todo
// @Description Getting all todos
// @ID get-all-todos
// @Accept json
// @Produce json
// @Param todoId path int true "todo's todo id"
// @Success 200 {object} doc_datas.GetAllTodosResponse
// @Failure 400 {object} error_utils.MessageErrData
// @Failure 404 {object} error_utils.MessageErrData
// @Failure 500 {object} error_utils.MessageErrData
// @Router /todo [get]
func GetAllTodos(c *gin.Context) {
	res, err := todo_service.TodoService.GetAllTodos()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteTodoById godoc
// @Summary Delete todo by ID
// @Tags todo
// @Description Deleting a todo by ID
// @ID delete-todo
// @Accept json
// @Produce json
// @Param todoId path int true "todo's todo id"
// @Success 200 {object} doc_datas.DeleteTodoResponse
// @Failure 400 {object} error_utils.MessageErrData
// @Failure 404 {object} error_utils.MessageErrData
// @Failure 500 {object} error_utils.MessageErrData
// @Router /todo/{todoId} [delete]
func DeleteTodoById(c *gin.Context) {
	var todo todo_domain.Todo

	todoId, err := todo.GetTodoIdParam(c)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	res, err := todo_service.TodoService.DeleteTodoById(todoId)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, res)
}
