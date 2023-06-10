package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Firstname string    `json:"firstname" valid:"required"`
	Lastname  string    `json:"lastname" valid:"required"`
	Email     string    `gorm:"unique" json:"email" binding:"required" valid:"required"`
	Apikey    string    `json:"api_key" valid:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Accounts struct {
	gorm.Model
	User_id        uint      `json:"user_id"`
	Account_name   string    `json:"account_name" gorm:"unique" valid:"required"`
	Account_number string    `json:"account_number" gorm:"unique"`
	Balance        int       `json:"balance"`
	Address        string    `gorm:"not null,size:256" json:"address" valid:"required"`
	CreatedAt      time.Time `json:"created_at"`
}

type CreditCards struct {
	gorm.Model
	User_id     uint      `json:"user_id"`
	Account_id  int       `json:"account_id" valid:"required"`
	Card_number int       `json:"card_number"`
	Expiry_date time.Time `json:"expiry_date"`
	Pin         uint      `json:"card_pin" valid:"required"`
	Cvv         int       `json:"cvv"`
	CreatedAt   time.Time `json:"created_at"`
}

type TransferHistory struct {
	gorm.Model
	Sender_account_number   string    `json:"sender_account_id" valid:"required"`
	Receiver_account_number string    `json:"receiver_account_id" valid:"required"`
	TransferHash            string    `json:"transfer_hash"`
	Amount                  int       `json:"amount" valid:"required"`
	CreatedAt               time.Time `json:"created_at"`
}

type TransactionsHistory struct {
	gorm.Model
	Account_id       uint      `json:"account_id"`
	Transaction_type string    `json:"transaction_type"`
	CreatedAt        time.Time `json:"created_at"`
}
