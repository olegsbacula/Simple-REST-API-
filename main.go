package main

import (
	"net/http"
	"errors"
	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record Video", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {

	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)

}

func toggleTodoStatus(context *gin.Context){
	id := context.Param("id")
	todo, error := getTodoById(id)
	
	if error!=nil{
		context.IndentedJSON(http.StatusNotFound, gin.H{"mesage":"Todo not found"})
		return
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}

func getTodoById(id string) (*todo,error){
	for i, t := range todos {
		if t.ID == id{
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}


func getTodo(context* gin.Context){
	id := context.Param("id")
	todo, error := getTodoById(id)
	
	if error!=nil{
		context.IndentedJSON(http.StatusNotFound, gin.H{"mesage":"Todo not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}



func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", addTodo)
	router.Run("localhost:9090")
}
