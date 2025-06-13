// @title           Todo API
// @version         1.0
// @description     Simple example of a Gin-based todo REST service.
// @host            localhost:9090
// @BasePath        /

package main

import (
	"errors"
	_ "example/docs"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type todo struct {
	ID        string `json:"id"        example:"4"`
	Item      string `json:"item"      example:"Example Task"`
	Completed bool   `json:"completed" example:"false"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record Video", Completed: false},
}

// getTodos godoc
// @Summary      Get a list of todos
// @Description  Get a list of all todos
// @Tags         todos
// @Produce      json
// @Success      200  {object}  nil
// @Failure      404  {object}  map[string]string  "Todo not found"
// @Router       /todos [get]
func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

// addTodo godoc
// @Summary      Add a todo
// @Description  Create a new todo
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        todo  body      Todo               true  "New todo payload"
// @Success      201   {object}  Todo
// @Failure      400   {object}  map[string]string  "Invalid payload"
// @Router       /todos [post]
func addTodo(c *gin.Context) {
    var newTodo todo
    if err := c.ShouldBindJSON(&newTodo); err != nil {
        c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid payload"})
        return
    }
    todos = append(todos, newTodo)
    c.JSON(http.StatusCreated, newTodo)
}
// toggleTodoStatus godoc
// @Summary      Toggle a todo status by it's ID
// @Description  Toggle a todo status by it's ID
// @Tags         todos
// @Produce      json
// @Param        id   path      string  true  "Todo ID"
// @Success      200  {object}  nil
// @Failure      404  {object}  map[string]string  "Todo not found"
// @Router       /todos/{id} [patch]
func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, error := getTodoById(id)

	if error != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"mesage": "Todo not found"})
		return
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}

func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

// getTodo godoc
// @Summary      Get a todo
// @Description  Get a todo by ID
// @Tags         todos
// @Produce      json
// @Param        id   path      string  true  "Todo ID"
// @Success      204  {object}  nil
// @Failure      404  {object}  map[string]string  "Todo not found"
// @Router       /todos/{id} [get]
func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, error := getTodoById(id)

	if error != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"mesage": "Todo not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

// deleteTodo godoc
// @Summary      Delete a todo
// @Description  Remove a todo by ID
// @Tags         todos
// @Param        id    path      string  true  "Todo ID"
// @Success      204   {object}  nil
// @Failure      404   {object}  map[string]string  "Todo not found"
// @Router       /todos/{id} [delete]
func deleteTodo(context *gin.Context) {
	id := context.Param("id")
	_, error := getTodoById(id)

	if error != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			break
		}
	}
	context.IndentedJSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", addTodo)
	router.DELETE("/todos/:id", deleteTodo)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run("localhost:9090")
}
