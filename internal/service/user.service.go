package service

import (
	"net/http"
	"strings"

	"github.com/hyper-dot/books-server/internal/config"
	"github.com/hyper-dot/books-server/internal/generated"
	"github.com/hyper-dot/books-server/internal/schema"
	"github.com/hyper-dot/books-server/internal/utils"
)

type Response struct {
	Message string `json:"message"`
}

func AddNewUser(w http.ResponseWriter, r *http.Request) {
	query, db := config.Query()
	defer db.Close()

	user, err := schema.ValidateUser(r)
	if err != nil {
		utils.RespondWithError("Invalid user data", w, http.StatusBadRequest, r.Context())
		return
	}

	/* CHECK IF USER EXISTS WITH THE EMAIL */
	_, err = query.GetUserByEmail(r.Context(), user.Email)
	if err == nil {
		utils.RespondWithError("User with the email already exists !!", w, http.StatusBadRequest, r.Context())
		return
	}

	/* GET USERNAME FROM EMAIL */
	username := strings.Split(user.Email, "@")[0]

	/* GENERATE HASH AND PASSWORD */
	hash, salt, err := utils.HashPassword(user.Password)

	var parsedUser = queries.AddUserParams{Username: username, Email: user.Email, Password: hash, Salt: salt}
	err = query.AddUser(r.Context(), parsedUser)

	if err != nil {
		utils.RespondWithError("Error writing data to the database", w, http.StatusInternalServerError, r.Context())
		return
	}

	utils.RespondWithJSON(Response{Message: "User added successfully !!"}, w, http.StatusOK)
}
