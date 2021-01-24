package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pankornch/restful-api/backend/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Find All Todos
func AllTodos(c *gin.Context) {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	var todos []model.Todo

	db.Find(&todos)

	c.JSON(http.StatusOK, todos)
}

// GetTodo
func GetTodo(c *gin.Context) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	id := c.Param("id")

	var todo model.Todo
	db.First(&todo, id)

	c.JSON(http.StatusOK, todo)
}

// Add Todo
func AddTodo(c *gin.Context) {
	var json model.TodoJSON

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	todo := model.Todo{Title: json.Title, Completed: false}

	db.Create(&todo)

	c.JSON(200, todo)
}

// Update Todo
func UpdateTodo(c *gin.Context) {
	var json model.TodoJSON
	id := c.Param("id")

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var curTodo model.Todo

	db.First(&curTodo, id)

	var todo model.Todo
	db.Where("ID = ?", id).Find(&todo)

	todo.Title = json.Title
	todo.Completed = json.Completed

	db.Save(&todo)

	c.JSON(http.StatusOK, todo)

}

// Delete Todo
func DeleteTodo(c *gin.Context) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var todo model.Todo
	id := c.Param("id")

	db.Where("ID = ?", id).Find(&todo)
	db.Delete(&todo)

	c.JSON(200, gin.H{
		"message": "Delete Todo",
	})
}
