package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shurlz/bank-backend/models"
	"github.com/shurlz/bank-backend/utils"
	"gorm.io/gorm"
)

func CreateTransfer(
	w http.ResponseWriter, r *http.Request, database *gorm.DB, user models.Users,
) {
	decoded := json.NewDecoder(r.Body)
	payload := models.TransferHistory{}
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

	go func() {
		s_account, r_account := models.Accounts{}, models.Accounts{}
		database.Where("account_number = ?", payload.Sender_account_number).First(&s_account)
		database.Where("account_number = ?", payload.Receiver_account_number).First(&r_account)

		s_account.Balance = s_account.Balance - payload.Amount
		r_account.Balance = r_account.Balance + payload.Amount

		database.Save(&s_account)
		database.Save(&r_account)
	}()

	new_transfer := models.TransferHistory{
		Sender_account_number:   payload.Sender_account_number,
		Receiver_account_number: payload.Receiver_account_number,
		TransferHash:            payload.TransferHash,
		Amount:                  payload.Amount,
	}
	results := database.Create(&new_transfer)
	if results.Error != nil {
		utils.SendErrorResponse(w, 400, fmt.Sprintf("%v", results.Error))
		return
	}
	utils.SendJsonResponse(w, 201, new_transfer)
}

func GetTransfers(
	w http.ResponseWriter, r *http.Request, database *gorm.DB, user models.Users,
) {
	accountId := struct {
		Account_Id string `valid:"required"`
	}{}
	decoded := json.NewDecoder(r.Body)
	history := []*models.TransferHistory{}

	err := decoded.Decode(&accountId)
	if err != nil {
		utils.SendErrorResponse(w, 400, fmt.Sprintf("%v", err))
		return
	}

	validation_err := utils.ValidateStruct(accountId)
	if validation_err != nil {
		utils.SendErrorResponse(w, 400, fmt.Sprintf("%v", validation_err))
		return
	}

	database.Where(
		"sender_account_number = ?", accountId.Account_Id,
	).Or(
		"receiver_account_number = ?", accountId.Account_Id,
	).Find(&history)

	utils.SendJsonResponse(w, 200, history)
}
