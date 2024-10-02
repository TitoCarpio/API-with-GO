package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

// variable de conexion a la base de datos
var DSN = "host=localhost user=postgres password=carpio2001 dbname=go_api port=5432"

// variables de las posibles conexiones a la base de datos
var DB *gorm.DB
var err error

// funcion para conectarse a la base de datos
func DBconnection() {
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Conexion exitosa")
	}
}
