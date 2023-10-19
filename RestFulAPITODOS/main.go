package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//==================================VARIABLES==================================

// The todo data type (Not gonna create an obj too lazy 4 that)
type todo struct {
	TodoID          string `json:"id"`
	TodoContent     string `json:"content"`
	IsFinished      bool   `json:"finished"`
	ExpirationYear  string `json:"endYear"`
	ExpirationMonth string `json:"endMonth"`
	ExpirationDay   string `json:"endDay"`
}

// The list of todos
var todolist = []todo{
	{TodoID: "1", TodoContent: "test1", IsFinished: false, ExpirationYear: "2023", ExpirationMonth: "11", ExpirationDay: "30"},
	{TodoID: "2", TodoContent: "test2", IsFinished: true, ExpirationYear: "2022", ExpirationMonth: "10", ExpirationDay: "22"},
	{TodoID: "3", TodoContent: "test3", IsFinished: false, ExpirationYear: "2023", ExpirationMonth: "12", ExpirationDay: "5"},
}

//==================================ENDPOINT FUNCTIONS==================================

// Getting all the todos
func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todolist)
}

// Adding a new TODO to the TODOLIST
func newTodo(context *gin.Context) {
	var addTodo todo

	if err := context.BindJSON(&addTodo); err != nil {
		return
	}
	if addTodo.TodoContent != "" && !isTodoExpired(&addTodo) {
		addTodo.IsFinished = false
		todolist = append(todolist, addTodo)
		context.JSON(http.StatusCreated, addTodo)
		return
	}
	context.JSON(http.StatusForbidden, gin.H{"message": "TODO WITH DESCRIPTION EMPTY AND/OR WITH INVALID DATE"})
}

// The functions that finds a todo via the specific ID
func findTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "TODO NOT FOUND IN LIST"})
		return
	}

	context.JSON(http.StatusFound, todo)
}

// Putting the todoStatus into DONE
func completeTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "TODO NOT FOUND IN LIST"})
	}
	todo.IsFinished = true

	context.IndentedJSON(http.StatusOK, todo)
}

// Editing the todo's content
func editTodoContent(context *gin.Context) {
	id := context.Param("id")
	newContent := context.Param("content")
	todo, err := getTodoById(id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "TODO NOT FOUND IN LIST"})
	}
	todo.TodoContent = newContent

	context.IndentedJSON(http.StatusOK, todo)
}

// Deleting a todo via his id
func deleteTodoViaID(context *gin.Context) {
	id := context.Param("id")
	var newTodoList []todo

	for i, t := range todolist {
		if t.TodoID != id {
			newTodoList = append(newTodoList, todolist[i])
		}
	}
	todolist = newTodoList
	context.JSON(http.StatusOK, gin.H{"message": "OPERATION OK"})
}

//==================================HANDLERS==================================

// Returns true if the todo is expired
func isTodoExpired(check *todo) bool {
	y, _ := strconv.Atoi(check.ExpirationYear)
	m, _ := strconv.Atoi(check.ExpirationMonth)
	d, _ := strconv.Atoi(check.ExpirationDay)
	t := time.Now()

	if check.ExpirationDay == "" || check.ExpirationMonth == "" || check.ExpirationYear == "" {
		return true
	}
	if y == t.Year() {
		fmt.Println("Ycheck Passed")
		if m >= int(t.Month()) {
			if d >= t.Day() && m <= int(t.Month()) || m >= int(t.Month()) {
				return false
			}
		}
	}

	return true
}

// Finding the todo with a specific todo via its ID
func getTodoById(id string) (*todo, error) {
	for i, t := range todolist {
		if t.TodoID == id {
			return &todolist[i], nil
		}
	}
	return nil, errors.New("TODO NOT FOUND IN LIST")
}

// Creating the "/todos" endpoint for get req. and starting server on port 2555
func main() {
	app := gin.Default()
	//Returns all the TODO
	app.GET("/todos", getTodos)
	//Finds a TODO by it's id
	app.GET("/todos/:id", findTodo)
	//Adds a new TODO
	app.POST("/todos", newTodo)
	//Edits TODO's desctiption
	app.PATCH("/todos/:id/", editTodoContent)
	//Marks the TODO as complete
	app.PATCH("/todos/:id/:content", completeTodo)
	//Deletes a TODO via it's ID
	app.DELETE("/todos/:id/", deleteTodoViaID)
	app.Run("localhost:2555")
}
