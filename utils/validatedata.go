package utils

import (
	"github.com/asaskevich/govalidator"
)

func ValidateStruct(s interface{}) error {
	_, err := govalidator.ValidateStruct(s)
	if err != nil {
		return err
	}
	return nil
}
