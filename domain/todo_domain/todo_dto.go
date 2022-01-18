package todo_domain

import (
	"assignment-4/utils/error_utils"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type Todo struct {
	Id          int64  `json:"id"`
	Title       string `json:"title" valid:"required~title is required"`
	Description string `json:"description" valid:"required~description is required"`
	Completed   bool   `json:"completed"`
}

func (t *Todo) Validate() error_utils.MessageErr {
	_, err := govalidator.ValidateStruct(t)

	if err != nil {
		return error_utils.NewBadRequest(err.Error())
	}

	return nil
}

func (t *Todo) GetTodoIdParam(c *gin.Context) (int64, error_utils.MessageErr) {
	paramId := c.Param("todoId")
	todoId, err := strconv.Atoi(paramId)

	if err != nil {
		return 0, error_utils.NewBadRequest("invalid todo id params")
	}

	return int64(todoId), nil
}
