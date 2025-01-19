package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"backend_CMS_Golang/src/config"
	"backend_CMS_Golang/src/models"

	"github.com/gin-gonic/gin"
)

func LoginWithPassword(w http.ResponseWriter, r *http.Request) {
	if config.Client == nil {
		http.Error(w, "Firestore client is not initialized", http.StatusInternalServerError)
		return
	}

	var loginCredentials models.LoginCredentials
	if err := json.NewDecoder(r.Body).Decode(&loginCredentials); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	fmt.Println("loginCredentials.Identifier: ", loginCredentials.Username)

	// Search for user by email
	iter := config.Client.Collection("users").Where("email", "==", loginCredentials.Username).Documents(config.Ctx)
	doc, err := iter.Next()
	if err != nil {
		fmt.Println("Error fetching document: ", err)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// User found, now compare the password
	var user models.User
	doc.DataTo(&user)
	fmt.Println("Fetched user: ", user)

	if loginCredentials.Password != user.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Respond with user information
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login successful",
		"userID":  user.ID,
	})
}

type LoginRequest struct {
	Token string `json:"token"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func LoginWtihGoogle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var loginRequest LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	token, err := config.AuthClient.VerifyIDToken(config.Ctx, loginRequest.Token)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	_, err = config.Client.Collection("users").Doc(token.UID).Set(config.Ctx, map[string]interface{}{
		"name":  loginRequest.Name,
		"email": loginRequest.Email,
	})
	if err != nil {
		http.Error(w, "Failed to save user data", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Login successful", "uid": token.UID})
}

func VerifyGoogleToken(c *gin.Context) {
	var body struct {
		IDtoken string `json:"idToken"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := config.AuthClient.VerifyIDToken(config.Ctx, body.IDtoken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	user, err := config.AuthClient.GetUser(context.Background(), body.IDtoken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Token verified", "user": user, "uid": token.UID})
}
