package router

import (
	"assignment-4/controllers/todo_controller"
	"assignment-4/db"

	"assignment-4/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
	// "github.com/swaggo/swag/example/basic/docs"
	// swagger embed files
)

var PORT = ":8080"

func init() {
	db.InitializeDB()
}

func StartRouter() {
	route := gin.Default()

	docs.SwaggerInfo.Title = "Example Swagger TODO Rest API"
	docs.SwaggerInfo.Description = "Documentation of TODO Rest API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost" + PORT
	docs.SwaggerInfo.Schemes = []string{"http"}

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	todoRoute := route.Group("/todo")
	{
		todoRoute.POST("/", todo_controller.CreateTodo)
		todoRoute.GET("/:todoId", todo_controller.GetTodoById)
		todoRoute.GET("/", todo_controller.GetAllTodos)
		todoRoute.PUT("/:todoId", todo_controller.UpdateTodo)
		todoRoute.DELETE("/:todoId", todo_controller.DeleteTodoById)
	}

	route.Run(PORT)
}
