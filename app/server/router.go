package server

import (
	"fmt"
	"net/http"
	"ushas/config"
	"ushas/controllers"
	"ushas/models"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Validator : Struct for validation.
type Validator struct {
	validator *validator.Validate
}

// Validate : Validate request body.
func (v *Validator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		results := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			msg := fmt.Sprintf("Field `%v` failed on '%s' restriction", e.Field(), e.Tag())
			results = append(results, msg)
		}
		return models.RaiseBadRequestError(err, "Requested body is invalid", results)
	}
	return nil
}

// NewRouter : Return pointer to router struct.
func NewRouter() (*echo.Echo, error) {
	c := config.GetConfig()
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: c.GetStringSlice("server.cors"),
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
		},
	}))

	e.Validator = &Validator{validator: validator.New()}

	v := e.Group("/" + c.GetString("server.version"))

	answer := controllers.NewAnswerController()
	log := controllers.NewLogController()
	task := controllers.NewTaskController()
	user := controllers.NewUserController()

	v.GET("/answer", answer.List)
	v.GET("/answer/:id", answer.Get)
	v.POST("/answer", answer.Create)
	v.PUT("/answer/:id", answer.Update)
	v.DELETE("/answer/:id", answer.Delete)

	v.GET("/log/dwell", log.Index)
	v.GET("/log/view", log.Index)
	v.GET("/log", log.Index)
	// The usage that getting each log doesn't expected.
	// v.GET("/log/:id", log.Index)
	v.POST("/log", log.Index)
	v.PUT("/log/:id", log.Index)
	v.DELETE("/log/:id", log.Index)

	v.GET("/task", task.Index)
	v.GET("/task/:id", task.Index)
	v.POST("/task", task.Index)
	v.PUT("/task/:id", task.Index)
	v.DELETE("/task", task.Index)

	v.GET("/user", user.List)
	v.GET("/user/:id", user.GetUserByID)
	v.POST("/user", user.Create)
	v.PUT("/user/:id", user.Update)
	v.DELETE("/user", user.Delete)

	return e, nil
}
