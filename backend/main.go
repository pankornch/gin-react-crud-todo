package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pankornch/restful-api/backend/handler"
	"github.com/pankornch/restful-api/backend/model"
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDatabase() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.Todo{})

}

func main() {
	r := gin.Default()

	initDatabase()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}

	r.Use(cors.New(config))

	r.GET("/api/todos", handler.AllTodos)
	r.GET("/api/todos/:id", handler.GetTodo)
	r.POST("/api/todos", handler.AddTodo)
	r.PATCH("/api/todos/:id", handler.UpdateTodo)
	r.DELETE("/api/todos/:id", handler.DeleteTodo)

	r.Run()

}
