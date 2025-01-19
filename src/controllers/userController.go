package controllers

import (
	"backend_CMS_Golang/src/config"
	"backend_CMS_Golang/src/models"
	"encoding/json"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gorilla/mux"
)

// getUsers retrieves all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	iter := config.Client.Collection("users").Documents(config.Ctx)
	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}
		var user models.User
		doc.DataTo(&user)
		user.ID = doc.Ref.ID
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// getUser retrieves a single user by ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	doc, err := config.Client.Collection("users").Doc(id).Get(config.Ctx)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	var user models.User
	doc.DataTo(&user)
	user.ID = doc.Ref.ID

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// createUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	_, _, err = config.Client.Collection("users").Add(config.Ctx, map[string]interface{}{
		"name":  user.Name,
		"email": user.Email,
	})
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}

// updateUser updates an existing user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	_, err = config.Client.Collection("users").Doc(id).Set(config.Ctx, map[string]interface{}{
		"name":  user.Name,
		"email": user.Email,
	}, firestore.MergeAll)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User updated successfully"))
}

// deleteUser deletes a user by ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_, err := config.Client.Collection("users").Doc(id).Delete(config.Ctx)
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))
}
