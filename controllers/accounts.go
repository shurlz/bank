package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shurlz/bank-backend/models"
	"github.com/shurlz/bank-backend/utils"
	"gorm.io/gorm"
)

func GetMyAccounts(
	w http.ResponseWriter, r *http.Request, database *gorm.DB, user models.Users,
) {
	accounts := []*models.Accounts{}

	database.Where("user_id = ?", user.ID).Find(&accounts)
	fmt.Println(user.ID)

	utils.SendJsonResponse(w, 200, accounts)
}

func CreateAccount(
	w http.ResponseWriter, r *http.Request, database *gorm.DB, user models.Users,
) {
	decoded := json.NewDecoder(r.Body)
	payload := models.Accounts{}
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

	new_account := models.Accounts{
		User_id:        user.ID,
		Account_name:   payload.Account_name,
		Account_number: payload.Account_number,
		Address:        payload.Address,
		Balance:        payload.Balance,
	}
	results := database.Create(&new_account)
	if results.Error != nil {
		utils.SendErrorResponse(w, 400, fmt.Sprintf("%v", results.Error))
		return
	}
	utils.SendJsonResponse(w, 201, new_account)
}
