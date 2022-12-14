package model

import "time"

type User struct {
	ID         uint      `gorm:"primary_key" json:"id"`
	Name       string    `json:"name"`
	Age        string    `json:"age"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

func (User) TableName() string { return "users" }
