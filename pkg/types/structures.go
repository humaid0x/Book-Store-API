package types

import validation "github.com/go-ozzo/ozzo-validation"

// BookRequest represents the request structure for book-related operations.
type BookRequest struct {
	ID          uint   `json:"bookID"`
	BookName    string `json:"bookName"`
	AuthorId    uint   `json:"authorID"`
	Description string `json:"description,omitempty"`
}

func (book BookRequest) Validate() error {
	return validation.ValidateStruct(&book,
		validation.Field(&book.BookName,
			validation.Required.Error("Book name cannot be empty"),
			validation.Length(1, 50)),
		validation.Field(&book.AuthorId,
			validation.Required.Error("Author Id cannot be empty")))
}



// AuthorRequest represents the request structure for author-related operations.
type AuthorRequest struct {
	ID          uint   `json:"authorID"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
}

func (author AuthorRequest) Validate() error {
	return validation.ValidateStruct(&author,
		validation.Field(&author.Name,
			validation.Required.Error("Author name cannot be empty"),
			validation.Length(1, 50)),
		validation.Field(&author.Address,
			validation.Required.Error("Address cannot be empty"),
			validation.Length(1, 100)))
}



// UserRequest represents the request structure for user-related operations.
type UserRequest struct{
	UserName string  `json:"username"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Phone    string  `json:"phone,omitempty"`
	Address  string  `json:"address,omitempty"`
}

func (user UserRequest) Validate() error{
	return validation.ValidateStruct(&user,
		validation.Field(&user.UserName, validation.Required.Error("User name is required"),
		validation.Length(1, 50)),
		validation.Field(&user.Password, validation.Required.Error("Password is required")),
		validation.Field(&user.Email, validation.Required.Error("Email is required")))
}


// Credentials represents user credentials for authentication
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (cred Credentials) Validate() error{
	return validation.ValidateStruct(&cred,
		validation.Field(&cred.Username, validation.Required.Error("User name is required"),
			validation.Length(1, 50)),
		validation.Field(&cred.Password, validation.Required.Error("Password is required")))
}