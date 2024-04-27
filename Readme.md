# Scale-X Service

Scale-X Service is a backend service that provides functionality for user authentication and book management. It includes endpoints for user login, fetching books, adding books, and deleting books.

## Installation

1. Clone the repository:

```bash
git clone https://github.com/KumaranElavazhagn/sclae-x-services.git

go run main.go

```

## Usage

### Endpoints
And I have attached a Postman collection in this repository.

1. **User Login**
   - **URL:** `/login`
   - **Method:** `POST`
   - **Request Body:** 
    ```json
    {
        "username": "admin",
        "password": "admin@123"
    }
    ```
   - **Response:** 
    ```json
    {
    "message": "Successfully Login"
    }
    ```

2. **Get Books**
   - **URL:** `/home`
   - **Method:** `GET`
   - **Response:** 
    ```json
    [
    {
        "name": "Book Name",
        "author": "Author Name",
        "publicationYear": 2022
    },
    {
        "name": "Another Book",
        "author": "Another Author",
        "publicationYear": 2020
    }
    ]
    ```

3. **Add Book**
   - **URL:** `/addBook`
   - **Method:** `POST`
   - **Request Body:** 
    ```json
    {
    "name": "New Book",
    "author": "New Author",
    "publicationYear": 2023
    }
    ```
   - **Response:** 
    ```json
    {
        "message": "Successfully book Added"
    }
    ```

4. **Delete Book**
   - **URL:** `/deleteBook?bookName=Book Name`
   - **Method:** `DELETE`
   - **Response:** 
    ```json
    {
    "message": "Successfully book deleted"
    }
    ```

### Dependencies
   - github.com/dgrijalva/jwt-go - JWT implementation for Go
   - github.com/gorilla/mux - HTTP router and dispatcher for Go
   - github.com/rs/cors - CORS handler for Go

### Contributing
Contributions are welcome! Please open an issue or submit a pull request for any new features or bug fixes.