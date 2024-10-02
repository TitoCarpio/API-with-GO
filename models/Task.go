package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `gorm:"not null"` //asi digo que el campo no puede ser nulo
	Description string
	Done        bool `gorm:"default:false"` //valor por defecto
	UserID      uint
}
