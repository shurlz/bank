package controllers

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"github.com/shurlz/bank-backend/models"
// 	"github.com/shurlz/bank-backend/utils"
// 	"gorm.io/gorm"
// )

// func Deposit(w http.ResponseWriter, r *http.Request, database *gorm.DB, user models.Users) {
// 	decoded := json.NewDecoder(r.Body)
// 	payload := models.TransferHistory{}
// 	payload.GenerateCredentials()

// 	decode_err := decoded.Decode(&payload)
// 	if decode_err != nil {
// 		utils.SendErrorResponse(w, 400, fmt.Sprintf("%v", decode_err))
// 		return
// 	}

// 	validation_err := utils.ValidateStruct(payload)
// 	if validation_err != nil {
// 		utils.SendErrorResponse(w, 400, fmt.Sprintf("%v", validation_err))
// 		return
// 	}

// 	new_transfer := models.TransferHistory{
// 		Sender_account_id:   payload.Sender_account_id,
// 		Receiver_account_id: payload.Receiver_account_id,
// 		TransferHash:        payload.TransferHash,
// 		Amount:              payload.Amount,
// 	}
// 	results := database.Create(&new_transfer)
// 	if results.Error != nil {
// 		utils.SendErrorResponse(w, 400, fmt.Sprintf("%v", results.Error))
// 		return
// 	}
// 	utils.SendJsonResponse(w, 201, new_transfer)
// }

// func Withdraw(w http.ResponseWriter, r *http.Request, database *gorm.DB, user models.Users) {
// 	decoded := json.NewDecoder(r.Body)
// 	payload := models.TransferHistory{}
// 	payload.GenerateCredentials()

// 	decode_err := decoded.Decode(&payload)
// 	if decode_err != nil {
// 		utils.SendErrorResponse(w, 400, fmt.Sprintf("%v", decode_err))
// 		return
// 	}

// 	validation_err := utils.ValidateStruct(payload)
// 	if validation_err != nil {
// 		utils.SendErrorResponse(w, 400, fmt.Sprintf("%v", validation_err))
// 		return
// 	}

// 	new_transfer := models.TransferHistory{
// 		Sender_account_id:   payload.Sender_account_id,
// 		Receiver_account_id: payload.Receiver_account_id,
// 		TransferHash:        payload.TransferHash,
// 		Amount:              payload.Amount,
// 	}
// 	results := database.Create(&new_transfer)
// 	if results.Error != nil {
// 		utils.SendErrorResponse(w, 400, fmt.Sprintf("%v", results.Error))
// 		return
// 	}
// 	utils.SendJsonResponse(w, 201, new_transfer)
// }
