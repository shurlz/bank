package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shurlz/bank-backend/models"
	"github.com/shurlz/bank-backend/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	database, db_err := utils.FetchDatabase()
	if db_err != nil {
		utils.SendErrorResponse(w, 400, fmt.Sprintf("%v", db_err))
		return
	}

	decoded := json.NewDecoder(r.Body)
	payload := models.Users{}
	payload.GenerateCredentials()
	err := decoded.Decode(&payload)
	if err != nil {
		utils.SendErrorResponse(w, 400, fmt.Sprintf("%v", err))
		return
	}

	validation_err := utils.ValidateStruct(payload)
	if validation_err != nil {
		utils.SendErrorResponse(w, 400, fmt.Sprintf("%v", validation_err))
		return
	}

	new_user := models.Users{
		Firstname: payload.Firstname,
		Lastname:  payload.Lastname,
		Email:     payload.Email,
		Apikey:    payload.Apikey,
	}
	results := database.Create(&new_user)
	if results.Error != nil {
		utils.SendErrorResponse(w, 400, fmt.Sprintf("%v", results.Error))
		return
	}
	utils.SendJsonResponse(w, 201, new_user)
}
