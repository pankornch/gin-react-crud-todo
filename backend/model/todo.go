package model

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title     string
	Completed bool
}

type TodoJSON struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
