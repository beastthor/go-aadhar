package main

import (
	"go-aadhar/controllers"
	"go-aadhar/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	initDB()
	log.Println("Starting the HTTP server on port 8095")

	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8095", router))
}

func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/create", controllers.CreateAadhar).Methods("POST")
	router.HandleFunc("/get", controllers.GetAllAadhar).Methods("GET")
	router.HandleFunc("/get/{id}", controllers.GetAadharByID).Methods("GET")
	router.HandleFunc("/update/{id}", controllers.UpdateAadharByID).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.DeletAadharByID).Methods("DELETE")
}

func initDB() {
	config :=
		database.Config{
			ServerName: "127.0.0.1:3306",
			User:       "root",
			Password:   "password",
			DB:         "aadharcard",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
}
