package service

import (
	"net/http"

	"github.com/hyper-dot/books-server/internal/config"
	"github.com/hyper-dot/books-server/internal/queries"
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

	var parsedUser = queries.AddUserParams{Username: user.Username, Email: user.Email, Password: user.Password, Salt: "salt", RefreshToken: "refesh token"}
	err = query.AddUser(r.Context(), parsedUser)

	if err != nil {
		utils.RespondWithError("Error writing data to the database", w, http.StatusInternalServerError, r.Context())
		return
	}

	utils.RespondWithJSON(Response{Message: "User added successfully !!"}, w, http.StatusOK)
}
