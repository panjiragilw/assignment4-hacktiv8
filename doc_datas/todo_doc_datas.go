package doc_datas

// Create ToDo

type CreateTodoResponse struct {
	Id          int64  `json:"id" example:"1"`
	Title       string `json:"title" example:"Make Dinner"`
	Description string `json:"description" example:"Cook fried rice with egg and chicken"`
	Completed   bool   `json:"completed" example:"false"`
}

type CreateTodoRequest struct {
	Title       string `json:"title" example:"Make Dinner"`
	Description string `json:"description" example:"Cook fried rice with egg and chicken"`
	Completed   bool   `json:"completed" example:"false"`
}

// Update ToDo

type UpdateTodoResponse struct {
	Id          int64  `json:"id" example:"1"`
	Title       string `json:"title" example:"Make Delicious Dinner"`
	Description string `json:"description" example:"Cook fried chicken with spicy sauce"`
	Completed   bool   `json:"completed" example:"false"`
}

type UpdateTodoRequest struct {
	Title       string `json:"title" example:"Make Delicious Dinner"`
	Description string `json:"description" example:"Cook fried chicken with spicy sauce"`
	Completed   bool   `json:"completed" example:"false"`
}

// Get ToDo By ID

type GetTodoResponse struct {
	Id          int64  `json:"id" example:"1"`
	Title       string `json:"title" example:"Make Delicious Dinner"`
	Description string `json:"description" example:"Cook fried chicken with spicy sauce"`
	Completed   bool   `json:"completed" example:"false"`
}

// Get All ToDo

type GetAllTodosResponse struct {
	Todos []GetTodoResponse
}

// Delete ToDo

type DeleteTodoResponse struct {
	StatusDelete string `json:"status_delete" example:"Success"`
	AffectedRow  int64  `json:"affected_row" example:"1"`
}
