package models

import (
	"github.com/google/uuid"
)

type Status string

type Task struct {
	Task_id     uuid.UUID `json:"t_id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
}

type User struct {
	UID      uuid.UUID `json:"u_id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username string    `json:"username"`
	Login    string    `json:"login"`
	Password string    `json:"-"`
}
