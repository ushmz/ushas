package controllers

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"ushas/models"
	"ushas/views"

	"github.com/labstack/echo/v4"
)

// UserController : Struct for controll `User` resource.
type UserController struct{}

// NewUserController : Return pointer to `UserController`.
func NewUserController() *UserController {
	return new(UserController)
}

// Index : index
func (uc *UserController) Index(c echo.Context) error {
	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		"OK",
	))
}

// CreateUserRequest : Struct for request of `/signup` endpoint
type CreateUserRequest struct {
	// UID : User name/ID for label.
	UID string `json:"uid" validate:"required"`
}

// Create : Creates new user.
func (uc *UserController) Create(c echo.Context) error {
	if uc == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("UserController is nil"))
	}

	p := CreateUserRequest{}
	if err := c.Bind(&p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, models.NewInternalError(err, "Invalid request body.", p))
	}

	if err := c.Validate(&p); err != nil {
		if e, ok := err.(*models.InternalError); ok {
			return echo.NewHTTPError(http.StatusBadRequest, e)
		}
		return echo.NewHTTPError(http.StatusBadRequest, models.NewInternalError(err, "Failed to validate.", p))
	}

	u, err := models.GetUserByUID(p.UID)
	if err != nil {
		fmt.Printf("\033[1;33m[INFO]\033[0m Username `%s` is not exist. Create new user.\n", p.UID)
		s := generateOneTimeSecret(32, 5, 5, 5, 5)
		model := &models.User{UID: p.UID, Secret: s}
		if err := models.CreateUser(model); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		u = model
	}

	return c.JSON(http.StatusCreated, newResponse(
		http.StatusCreated,
		http.StatusText(http.StatusCreated),
		u,
	))
}

// GetUserByIDRequest : Request parameters for "/user/:id"
type GetUserByIDRequest struct {
	ID int `json:"id" param:"id" validate:"required,numeric"`
}

// GetByID : Gets an user by ID.
func (uc *UserController) GetByID(c echo.Context) error {
	if uc == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("UserController is nil"))
	}

	p := GetUserByIDRequest{}
	if err := c.Bind(&p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, models.NewInternalError(err, "Invalid request body.", p))
	}

	if err := c.Validate(&p); err != nil {
		if e, ok := err.(*models.InternalError); ok {
			return echo.NewHTTPError(http.StatusBadRequest, e)
		}
		return echo.NewHTTPError(http.StatusBadRequest, models.NewInternalError(err, "Failed to validate.", p))
	}

	u, err := models.GetUserByID(p.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	uv := views.NewUserView(u)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		uv,
	))
}

// GetUserByUIDRequest : Request parameters for "/user/:uid"
type GetUserByUIDRequest struct {
	UID string `json:"uid" param:"uid" validate:"required"`
}

// GetByUID : Gets an user by UID.
func (uc *UserController) GetByUID(c echo.Context) error {
	if uc == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("UserController is nil"))
	}

	p := GetUserByUIDRequest{}
	if err := c.Bind(&p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, models.NewInternalError(err, "Invalid request body.", p))
	}

	if err := c.Validate(&p); err != nil {
		if e, ok := err.(*models.InternalError); ok {
			return echo.NewHTTPError(http.StatusBadRequest, e)
		}
		return echo.NewHTTPError(http.StatusBadRequest, models.NewInternalError(err, "Failed to validate.", p))
	}

	u, err := models.GetUserByUID(p.UID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	uv := views.NewUserView(u)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		uv,
	))
}

// List : Lists all users.
func (uc *UserController) List(c echo.Context) error {
	if uc == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("UserController is nil"))
	}

	us, err := models.ListUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	uv := views.NewListUserView(us)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		uv,
	))
}

// UpdateUserRequest : Struct for request of `user` update endpoint
type UpdateUserRequest struct {
	// ID : The ID of user.
	ID int `json:"id" param:"id" validate:"required,numeric"`

	// UID : User name/ID for label.
	UID string `json:"uid" validate:"required"`
}

// Update : Updates user information. This connot update password.
func (uc *UserController) Update(c echo.Context) error {
	if uc == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("UserController is nil"))
	}

	p := UpdateUserRequest{}
	if err := c.Bind(&p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, models.NewInternalError(err, "Invalid request body.", p))
	}

	if err := c.Validate(p); err != nil {
		if e, ok := err.(*models.InternalError); ok {
			return echo.NewHTTPError(http.StatusBadRequest, e)
		}
		return echo.NewHTTPError(http.StatusBadRequest, models.NewInternalError(err, "Failed to validate.", p))
	}

	u := &models.User{ID: p.ID, UID: p.UID}
	if err := models.UpdateUser(u); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	uv := views.NewUserView(u)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		uv,
	))
}

// DeleteUserRequest : Request parameters to delete user.
type DeleteUserRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}

// Delete : Deletes an user.
func (uc *UserController) Delete(c echo.Context) error {
	if uc == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("UserController is nil"))
	}

	p := DeleteTaskRequest{}
	if err := c.Bind(&p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, models.NewInternalError(err, "Invalid request body.", p))
	}

	if err := c.Validate(&p); err != nil {
		if e, ok := err.(*models.InternalError); ok {
			return echo.NewHTTPError(http.StatusBadRequest, e)
		}
		return echo.NewHTTPError(http.StatusBadRequest, models.NewInternalError(err, "Failed to validate.", p))
	}

	if err := models.DeleteUser(p.ID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		p.ID,
	))
}

// generateOneTimeSecret : Generate password for new user.
func generateOneTimeSecret(length, lower, upper, digits, symbols int) string {
	var (
		lowerCharSet = "abcdedfghijklmnopqrst"
		upperCharSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		digitsSet    = "0123456789"
		symbolsSet   = "!@#$%&*"
		allCharSet   = lowerCharSet + upperCharSet + digitsSet + symbolsSet
	)

	var passwd strings.Builder

	for i := 0; i < lower; i++ {
		random := rand.Intn(len(lowerCharSet))
		passwd.WriteString(string(lowerCharSet[random]))
	}

	for i := 0; i < upper; i++ {
		random := rand.Intn(len(upperCharSet))
		passwd.WriteString(string(upperCharSet[random]))
	}

	for i := 0; i < digits; i++ {
		random := rand.Intn(len(digitsSet))
		passwd.WriteString(string(digitsSet[random]))
	}

	for i := 0; i < symbols; i++ {
		random := rand.Intn(len(symbolsSet))
		passwd.WriteString(string(symbolsSet[random]))
	}

	remaining := length - lower - upper - digits - symbols
	for i := 0; i < remaining; i++ {
		random := rand.Intn(len(allCharSet))
		passwd.WriteString(string(allCharSet[random]))
	}

	inRune := []rune(passwd.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})

	return string(inRune)
}
