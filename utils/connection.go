package utils

import (
	"log"

	"github.com/jinzhu/gorm"

	// mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// GetConnection obtiene una conexión a la base de datos
func GetConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root1234@tcp(database-1.c9az1iogvwyg.us-east-1.rds.amazonaws.com:3306)/microservicio?charset=utf8&parseTime=True&loc=Local") // no olvidar poner luego del nombre de la db : ?charset=utf8&parseTime=True&loc=Local
	if err != nil {
		log.Fatal(err)
	}
	return db
}
