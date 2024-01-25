package controllers

import (
	"book-store/pkg/domain"
	"book-store/pkg/models"
	"book-store/pkg/types"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// AuthorController struct to access the methods of service and repo
type AuthorController struct {
	authorSvc domain.IAuthorService
}

// NewAuthorController initializes a new AuthorController
func NewAuthorController(authorSvc domain.IAuthorService) AuthorController {
	return AuthorController{
		authorSvc: authorSvc,
	}
}

// CreateAuthor handles the creation of a new author
func (ac *AuthorController) CreateAuthor(e echo.Context) error {
	reqAuthor := &types.AuthorRequest{}
	if err := e.Bind(reqAuthor); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	if err := reqAuthor.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	author := &models.Author{
		Name:        reqAuthor.Name,
		Address:     reqAuthor.Address,
		PhoneNumber: reqAuthor.PhoneNumber,
	}

	if err := ac.authorSvc.CreateAuthor(author); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusCreated, "Author was created successfully")
}

// GetAuthor retrieves information about a specific author
func (ac *AuthorController) GetAuthor(e echo.Context) error {
	tempAuthorID := e.QueryParam("authorID")
	authorID, err := strconv.ParseUint(tempAuthorID, 0, 0)
	if err != nil && tempAuthorID != "" {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	author, err := ac.authorSvc.GetAuthor(uint(authorID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, author)
}

// UpdateAuthor handles the update of an existing author
func (ac *AuthorController) UpdateAuthor(e echo.Context) error {
	reqAuthor := &types.AuthorRequest{}
	if err := e.Bind(reqAuthor); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	if err := reqAuthor.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	tempAuthorID := e.Param("authorID")
	authorID, err := strconv.ParseUint(tempAuthorID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	existingAuthor, err := ac.authorSvc.GetAuthor(uint(authorID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	existingAuthorData := existingAuthor[0]

	updatedAuthor := &models.Author{
		ID:          uint(authorID),
		Name:        reqAuthor.Name,
		Address:     reqAuthor.Address,
		PhoneNumber: reqAuthor.PhoneNumber,
	}

	if reqAuthor.Name == "" {
		updatedAuthor.Name = existingAuthorData.Name
	}

	if reqAuthor.Address == "" {
		updatedAuthor.Address = existingAuthorData.Address
	}

	if reqAuthor.PhoneNumber == "" {
		updatedAuthor.PhoneNumber = existingAuthorData.PhoneNumber
	}

	if err := ac.authorSvc.UpdateAuthor(updatedAuthor); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusCreated, "Author was updated successfully")
}

// DeleteAuthor handles the deletion of an author
func (ac *AuthorController) DeleteAuthor(e echo.Context) error {
	tempAuthorID := e.Param("authorID")
	authorID, err := strconv.ParseUint(tempAuthorID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	_, err = ac.authorSvc.GetAuthor(uint(authorID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	if err := ac.authorSvc.DeleteAuthor(uint(authorID)); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, "Author was deleted successfully")
}
