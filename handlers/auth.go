package handlers

import (
	"auth-api/database"
	"auth-api/models"
	"auth-api/utils"
	"encoding/json"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		utils.ErrorResponse(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	hashedPassword, err := utils.HashPassword(creds.Password)
	if err != nil {
		utils.ErrorResponse(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	query := "INSERT INTO users(email, password) VALUES(?, ?)"
	_, err = database.Execute(query, creds.Email, hashedPassword)
	if err != nil {
		utils.ErrorResponse(w, "User already exists", http.StatusBadRequest)
		return
	}

	user := models.User{Email: creds.Email}
	utils.JSONResponse(w, user, http.StatusCreated)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		utils.ErrorResponse(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var user models.User
	query := "SELECT id, email, password FROM users WHERE email = ?"
	err := database.QueryRow(query, creds.Email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		utils.ErrorResponse(w, "User not found", http.StatusUnauthorized)
		return
	}

	err = utils.CheckPassword(user.Password, creds.Password)
	if err != nil {
		utils.ErrorResponse(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	tokenString, err := utils.GenerateToken(user.Email)
	if err != nil {
		utils.ErrorResponse(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	utils.JSONResponse(w, map[string]string{"token": tokenString}, http.StatusOK)
}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	tokenString := utils.GetTokenFromHeader(r)
	if tokenString == "" {
		utils.ErrorResponse(w, "Authorization header missing or invalid", http.StatusUnauthorized)
		return
	}

	claims, err := utils.ParseToken(tokenString)
	if err != nil {
		utils.ErrorResponse(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	newTokenString, err := utils.GenerateToken(claims.Email)
	if err != nil {
		utils.ErrorResponse(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	utils.JSONResponse(w, map[string]string{"token": newTokenString}, http.StatusOK)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	tokenString := utils.GetTokenFromHeader(r)
	if tokenString == "" {
		utils.ErrorResponse(w, "Auth token missing", http.StatusUnauthorized)
		return
	}

	utils.BlacklistToken(tokenString)
	utils.JSONResponse(w, map[string]string{"message": "Successfully logged out"}, http.StatusOK)
}

func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
	utils.JSONResponse(w, map[string]string{"message": "Hello you can access to Private Route."}, http.StatusOK)
}
