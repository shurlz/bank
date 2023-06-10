package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shurlz/bank-backend/models"
	"github.com/shurlz/bank-backend/utils"
	"gorm.io/gorm"
)

func GetMyCards(
	w http.ResponseWriter, r *http.Request, database *gorm.DB, user models.Users,
) {
	cards := []*models.CreditCards{}

	database.Where("user_id = ?", user.ID).Find(&cards)

	utils.SendJsonResponse(w, 200, cards)
}

func CreateCard(
	w http.ResponseWriter, r *http.Request, database *gorm.DB, user models.Users,
) {
	decoded := json.NewDecoder(r.Body)
	payload := models.CreditCards{}
	payload.GenerateCredentials()

	decode_err := decoded.Decode(&payload)
	if decode_err != nil {
		utils.SendErrorResponse(w, 400, fmt.Sprintf("%v", decode_err))
		return
	}

	validation_err := utils.ValidateStruct(payload)
	if validation_err != nil {
		utils.SendErrorResponse(w, 400, fmt.Sprintf("%v", validation_err))
		return
	}

	new_card := models.CreditCards{
		User_id:     user.ID,
		Account_id:  payload.Account_id,
		Card_number: payload.Card_number,
		Expiry_date: payload.Expiry_date,
		Pin:         payload.Pin,
		Cvv:         payload.Cvv,
	}
	results := database.Create(&new_card)
	if results.Error != nil {
		utils.SendErrorResponse(w, 400, fmt.Sprintf("%v", results.Error))
		return
	}
	utils.SendJsonResponse(w, 201, new_card)
}
