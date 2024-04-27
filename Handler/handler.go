package handler

import (
	"encoding/json"
	"net/http"
	Service "scale-x/Service"
	"time"

	"scale-x/dto"
)

// Handlers struct contains methods to handle HTTP requests
type Handlers struct {
	Service Service.Service // Service layer dependency
}

// LoginHandler handles user login requests
func (h *Handlers) LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Decode login credentials from request body
	var creds dto.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}

	// Call service function to authenticate user
	token, err := h.Service.AuthenticateUser(creds)
	if err != nil {
		writeResponse(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	// Set the token in the cookie
	expirationTime := time.Now().Add(5 * time.Minute)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expirationTime,
	})

	writeResponse(w, http.StatusOK, "Successfully Login")
}

// HomeHandler handles requests for user's home page
func (h *Handlers) HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Extract token from cookie and authenticate user using service method
	books, err := h.Service.GetBooks(r)
	if err != nil {
		writeResponse(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	writeResponse(w, http.StatusOK, books)
}

// AddBookHandler handles requests to add a book
func (h *Handlers) AddBookHandler(w http.ResponseWriter, r *http.Request) {
	// Extract token from cookie and authenticate user using service method
	err := h.Service.AddBook(r)
	if err != nil {
		writeResponse(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	writeResponse(w, http.StatusOK, "Successfully book Added")
}

// DeleteBookHandler handles requests to delete a book
func (h *Handlers) DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	// Extract token from cookie and authenticate user using service method
	err := h.Service.DeleteBook(r)
	if err != nil {
		writeResponse(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	writeResponse(w, http.StatusOK, "Successfully book deleted")
}

// writeResponse writes the HTTP response with appropriate headers and status code
func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
