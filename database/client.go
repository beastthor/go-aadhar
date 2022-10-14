package database

import (
	"log"

	"github.com/jinzhu/gorm"
)

//Connector variable used for crud operations

var Connector *gorm.DB

// Connect creates MySql connection
func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	log.Println("connection was succesful!!")
	return nil
}
