package models

import "gorm.io/gorm"

type Status string

type Task struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      Status `json:"status"`
}

type User struct {
	gorm.Model
	Username string `json:"username"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
