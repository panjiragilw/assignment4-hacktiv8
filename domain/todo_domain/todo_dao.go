package todo_domain

import (
	"assignment-4/db"
	"assignment-4/utils/error_formats"
	"assignment-4/utils/error_utils"
)

const (
	queryCreateTodo = `
		INSERT INTO todos 
		(title, description, completed) 
		VALUES ($1, $2, $3)
		RETURNING id, title, description, completed
	`
	queryUpdateTodo = `
		UPDATE todos
		SET title = $2, description = $3, completed = $4
		WHERE id = $1
		RETURNING id, title, description, completed
	`
	queryGetTodoById = `
		SELECT id, title, description, completed 
		FROM todos
		WHERE id = $1
	`
	queryGetAllTodos = `
		SELECT id, title, description, completed 
		FROM todos
	`
	queryDeleteTodoById = `
		DELETE
		FROM todos
		WHERE id = $1
	`
)

var TodoDomain todoDomain = &todoRepo{}

type todoDomain interface {
	CreateTodo(*Todo) (*Todo, error_utils.MessageErr)
	UpdateTodo(*Todo) (*Todo, error_utils.MessageErr)
	GetTodoById(int64) (*Todo, error_utils.MessageErr)
	GetAllTodos() (*[]Todo, error_utils.MessageErr)
	DeleteTodoById(int64) (*map[string]interface{}, error_utils.MessageErr)
}

type todoRepo struct{}

func (m *todoRepo) CreateTodo(todoReq *Todo) (*Todo, error_utils.MessageErr) {
	db := db.GetDB()

	row := db.QueryRow(queryCreateTodo, todoReq.Title, todoReq.Description, todoReq.Completed)

	var todo Todo
	err := row.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Completed)

	if err != nil {
		return nil, error_formats.ParseError(err)
	}
	return &todo, nil
}

func (m *todoRepo) UpdateTodo(todoReq *Todo) (*Todo, error_utils.MessageErr) {
	db := db.GetDB()
	row := db.QueryRow(queryUpdateTodo, todoReq.Id, todoReq.Title, todoReq.Description, todoReq.Completed)
	//id, title, image_url, user_id
	var todo Todo
	err := row.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Completed)

	if err != nil {
		return nil, error_formats.ParseError(err)
	}

	return &todo, nil
}

func (m *todoRepo) GetTodoById(todoId int64) (*Todo, error_utils.MessageErr) {
	db := db.GetDB()
	row := db.QueryRow(queryGetTodoById, todoId)

	var todo Todo
	err := row.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Completed)
	if err != nil {
		return nil, error_formats.ParseError(err)
	}

	return &todo, nil
}

func (m *todoRepo) GetAllTodos() (*[]Todo, error_utils.MessageErr) {
	db := db.GetDB()
	row, err := db.Query(queryGetAllTodos)
	if err != nil {
		return nil, error_formats.ParseError(err)
	}

	var todos []Todo

	for row.Next() {
		var todo Todo
		err := row.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Completed)
		if err != nil {
			return nil, error_formats.ParseError(err)
		}
		todos = append(todos, todo)
	}

	return &todos, nil
}

func (m *todoRepo) DeleteTodoById(todoId int64) (*map[string]interface{}, error_utils.MessageErr) {
	db := db.GetDB()
	res, err := db.Exec(queryDeleteTodoById, todoId)
	if err != nil {
		return nil, error_formats.ParseError(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		return nil, error_formats.ParseError(err)
	}

	deleteResult := map[string]interface{}{
		"StatusDelete": "Success",
		"AffectedRow":  count,
	}

	return &deleteResult, nil
}
