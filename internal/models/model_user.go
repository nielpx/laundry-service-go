package models

type User struct {
	ID       uint   `gorm:"unique" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}