package dto

// User struct represents a user with username and userType
type User struct {
	Username string `json:"username"`
	UserType string `json:"userType"`
}

// Credentials struct represents the login credentials
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Book struct represents a book with name, author, and publication year
type Book struct {
	Name            string `json:"name"`
	Author          string `json:"author"`
	PublicationYear int    `json:"publicationYear"`
}

// JwtKey is the key used for JWT token signing
var JwtKey = []byte("secret_key")

// Sample users with their passwords
var Users = map[string]string{
	"admin":   "admin@123",   // Admin user with password "admin@123"
	"regular": "regular@123", // Regular user with password "regular@123"
}
