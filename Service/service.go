package Service

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"scale-x/dto"
	"scale-x/parser"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Service struct contains methods to handle business logic
type Service struct {
	// Define any dependencies here
}

// NewService initializes a new instance of the Service
func NewService() *Service {
	// Perform any setup needed, if any
	return &Service{}
}

// ErrUnauthorized is returned for unauthorized access
var ErrUnauthorized = errors.New("unauthorized access")

// ErrBadRequest is returned for bad requests
var ErrBadRequest = errors.New("Bad Request")

// AuthenticateUser authenticates the user and returns a JWT token
func (s *Service) AuthenticateUser(creds dto.Credentials) (string, error) {
	// Check if the provided credentials are valid
	expectedPassword, ok := dto.Users[creds.Username]
	if !ok || expectedPassword != creds.Password {
		return "", ErrUnauthorized
	}

	// Create a JWT token with claims
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Subject:   creds.Username,
	}

	// Sign the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(dto.JwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// GetBooks retrieves books based on user type
func (s *Service) GetBooks(r *http.Request) ([]dto.Book, error) {
	// Authentication
	claims, err := s.authenticateUser(r)
	if err != nil {
		return nil, err
	}

	// Fetch books based on user type
	var books []dto.Book
	if claims.Subject == "admin" {
		adminBooks, _ := parser.ReadBooksFromFile("adminUser.csv")
		books = append(books, adminBooks...)
	}

	regularBooks, _ := parser.ReadBooksFromFile("regularUser.csv")
	books = append(books, regularBooks...)

	return books, nil
}

// AddBook adds a new book to the system
func (s *Service) AddBook(r *http.Request) error {
	// Authentication
	claims, err := s.authenticateUser(r)
	if err != nil {
		return err
	}

	// Authorization: Only admins can add books
	if claims.Subject != "admin" {
		return ErrUnauthorized
	}

	// Read book details from request
	var book dto.Book
	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		return err
	}

	// Validate book details
	if book.Name == "" || book.Author == "" || book.PublicationYear <= 0 {
		return ErrBadRequest
	}

	// Append the book to regularUser.csv
	f, err := os.OpenFile("regularUser.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	// Write the new book entry to the CSV file
	err = writer.Write([]string{book.Name, book.Author, strconv.Itoa(book.PublicationYear)})
	if err != nil {
		return err
	}
	writer.Flush()

	return nil
}

// DeleteBook deletes a book from the system
func (s *Service) DeleteBook(r *http.Request) error {
	// Authentication
	claims, err := s.authenticateUser(r)
	if err != nil {
		return err
	}

	// Authorization: Only admins can delete books
	if claims.Subject != "admin" {
		return ErrUnauthorized
	}

	// Extract book name from request
	params := r.URL.Query()
	bookName := params.Get("bookName")
	if bookName == "" {
		return ErrBadRequest
	}

	// Read regularUser.csv and filter out the book to delete
	books, _ := parser.ReadBooksFromFile("regularUser.csv")
	var filteredBooks []dto.Book
	for _, b := range books {
		if b.Name != bookName {
			filteredBooks = append(filteredBooks, b)
		}
	}

	// Rewrite regularUser.csv with filtered books
	f, err := os.Create("regularUser.csv")
	if err != nil {
		return err
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	for _, b := range filteredBooks {
		err = writer.Write([]string{b.Name, b.Author, strconv.Itoa(b.PublicationYear)})
		if err != nil {
			return err
		}
	}

	return nil
}

// authenticateUser authenticates the user using the JWT token
func (s *Service) authenticateUser(r *http.Request) (*jwt.StandardClaims, error) {
	// Extract token from cookie
	cookie, err := r.Cookie("token")
	if err != nil {
		return nil, err
	}

	tokenString := cookie.Value
	claims := &jwt.StandardClaims{}

	// Validate and parse JWT token
	_, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return dto.JwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	return claims, nil
}
