package models

import "gorm.io/gorm"

type User struct {
	gorm.Model //para que el ORM me cree automaticamnete la tabla

	FirtsName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"not null;unique_index"`
	Task      []Task `gorm:"foreignKey:UserID"`
}
