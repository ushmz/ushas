package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"ushas/models"
	"ushas/views"

	"github.com/labstack/echo/v4"
)

type UserController struct{}

func NewUserController() *UserController {
	return new(UserController)
}

func (uc *UserController) Index(c echo.Context) error {
	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		"OK",
	))
}

// CreateUserRequest : Struct for request of `/signup` endpoint
type CreateUserRequest struct {
	// Uid : User name/ID for label.
	Uid string `json:"uid" validate:"required"`
}

func (uc *UserController) Create(c echo.Context) error {
	if uc == nil {
		new500Response(c, nil, nil)
	}

	p := CreateUserRequest{}
	if err := c.Bind(&p); err != nil {
		return c.JSON(http.StatusBadRequest, newResponse(
			http.StatusBadRequest,
			"Failed to parse request body",
			p,
		))
	}

	if err := c.Validate(p); err != nil {
		if e, ok := err.(*models.APIError); ok {
			return c.JSON(e.Code, newResponse(e.Code, e.Message, e.Result))
		}
		return c.JSON(http.StatusBadRequest, newResponse(
			http.StatusBadRequest,
			err.Error(),
			p,
		))
	}

	u, err := models.GetUserByUID(p.Uid)
	if err != nil {
		fmt.Printf("\033[1;33m[INFO]\033[0m Username `%s` is not exist. Create new user.\n", p.Uid)
		s := generateOneTimeSecret(32, 5, 5, 5, 5)
		model := &models.User{UID: p.Uid, Secret: s}
		if err := models.CreateUser(model); err != nil {
			if e, ok := err.(*models.APIError); ok {
				return c.JSON(e.Code, newResponse(e.Code, e.Message, e.Result))
			}
			return c.JSON(http.StatusInternalServerError, newResponse(
				http.StatusInternalServerError,
				http.StatusText(http.StatusInternalServerError),
				p,
			))
		}
		u = model
	}

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		u,
	))
}

func (uc *UserController) GetUserByID(c echo.Context) error {
	if uc == nil {
		new500Response(c, nil, nil)
	}

	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, newResponse(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			"User ID must be number",
		))
	}

	u, err := models.GetUserByID(id)
	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		u,
	))
}

func (uc *UserController) GetUserByUID(c echo.Context) error {
	if uc == nil {
		return newErrResponse(c, http.StatusInternalServerError, nil, nil)
	}

	uid := c.Param("uid")

	if len(uid) > 0 {
		return c.JSON(http.StatusBadRequest, newResponse(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			"UID must be string and non-zero value",
		))
	}

	u, err := models.GetUserByUID(uid)
	if err != nil {
		return newErrResponse(c, http.StatusNotFound, err, uid)
	}

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		u,
	))
}

func (uc *UserController) List(c echo.Context) error {
	if uc == nil {
		return newErrResponse(c, http.StatusInternalServerError, nil, nil)
	}

	us, err := models.ListUser()
	if err != nil {
		return newErrResponse(c, http.StatusInternalServerError, err, nil)
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
	ID int `json:"id" validate:"required"`

	// Uid : User name/ID for label.
	Uid string `json:"uid" validate:"required"`
}

// Update : Update user information. This connot update password.
func (uc *UserController) Update(c echo.Context) error {
	if uc == nil {
		return newErrResponse(c, http.StatusInternalServerError, nil, nil)
	}

	p := UpdateUserRequest{}
	if err := c.Bind(&p); err != nil {
		return newErrResponse(c, http.StatusBadRequest, err, nil)
	}

	if err := c.Validate(p); err != nil {
		return newErrResponse(c, http.StatusBadRequest, err, p)
	}

	u := &models.User{ID: p.ID, UID: p.Uid}
	if err := models.UpdateUser(u); err != nil {
		return newErrResponse(c, http.StatusInternalServerError, err, nil)
	}

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		u,
	))
}

func (uc *UserController) Delete(c echo.Context) error {
	if uc == nil {
		return newErrResponse(c, http.StatusInternalServerError, nil, nil)
	}

	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, newResponse(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			"User ID must be number",
		))
	}

	if err := models.DeleteUser(id); err != nil {
		return newErrResponse(c, http.StatusInternalServerError, err, id)
	}

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		id,
	))
}

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
