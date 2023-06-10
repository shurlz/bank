package models

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type credentials interface {
	GenerateCredentials()
}

func (u *Users) GenerateCredentials() {
	u.Apikey = fmt.Sprintf("%v", uuid.New().String())
}

func (c *CreditCards) GenerateCredentials() {
	c.Card_number = rand.Intn(100000) + 899999
	c.Cvv = rand.Intn(899) + 100
	c.Expiry_date = time.Now().AddDate(4, 0, 0)
}

func (a *Accounts) GenerateCredentials() {
	a.Account_number = fmt.Sprintf("%v", rand.Intn(1000000)+9999999)
	a.Balance = 0
}

func (t *TransferHistory) GenerateCredentials() {
	t.TransferHash = fmt.Sprintf("%v", uuid.New().String())
}
