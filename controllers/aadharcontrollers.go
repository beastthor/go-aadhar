package controllers

import (
	"encoding/json"
	"go-aadhar/database"
	"go-aadhar/entity"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetAllAadhar get all aadhar data
func GetAllAadhar(w http.ResponseWriter, r *http.Request) {
	var aadhar []entity.Aadhar
	database.Connector.Find(&aadhar)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(aadhar)
}

// GetAadharByID returns person with specific ID
func GetAadharByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var aadhar entity.Aadhar
	database.Connector.First(&aadhar, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(aadhar)

}

// CreateAadhar creates aadhar
func CreateAadhar(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var aadhar entity.Aadhar
	json.Unmarshal(requestBody, &aadhar)

	database.Connector.Create(aadhar)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(aadhar)
}

// UpdateAadharByID updates person with respective ID
func UpdateAadharByID(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var aadhar entity.Aadhar
	json.Unmarshal(requestBody, &aadhar)
	database.Connector.Save(&aadhar)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(aadhar)
}

// DeleteAadharByID delete's person with specific ID
func DeletAadharByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var aadhar entity.Aadhar
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&aadhar)
	w.WriteHeader(http.StatusNoContent)
}
