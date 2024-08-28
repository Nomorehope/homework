package models

type Status string

const (
	New        Status = "Новая"
	InProgress Status = "В процессе"
	Completed  Status = "Завершена"
)

type Task struct {
	Task_id     int    `json:"t_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type User struct {
	UID      int    `json:"u_id"`
	Username string `json:"username"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

var Tasks []Task
var Users []User
