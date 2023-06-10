package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/shurlz/bank-backend/models"
	"github.com/shurlz/bank-backend/utils"
	"gorm.io/gorm"
)

type customHandlerFunc func(http.ResponseWriter, *http.Request, *gorm.DB, models.Users)

func ChechAuthStatus(handler customHandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// logic for getting auth stuffs

		database, db_err := utils.FetchDatabase()
		if db_err != nil {
			utils.SendErrorResponse(w, 400, fmt.Sprintf("%v", db_err))
			return
		}

		authorization := r.Header.Get("Authorization")
		split_auth := strings.Split(authorization, " ")
		if len(split_auth) != 2 || split_auth[0] != "Bearer" {
			utils.SendErrorResponse(w, 403, "Corrupted Authorization")
			return
		}

		user := models.Users{}
		result := database.First(&user, "apikey = ?", split_auth[1])
		if result.Error != nil {
			utils.SendErrorResponse(w, 401, "Unrecognized token, Unauthorized")
			return
		}

		handler(w, r, database, user)
	}
}
