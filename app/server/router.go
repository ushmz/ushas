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

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		results := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			msg := fmt.Sprintf("Field `%v` failed on '%s' restriction", e.Field(), e.Tag())
			results = append(results, msg)
		}

		return models.RaiseBadRequestError(err, "", results)

	}
	return nil
}

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
	v.PUT("/answer", answer.Update)
	v.DELETE("/answer/:id", answer.Delete)

	v.GET("/log", log.Index)
	v.POST("/log", log.Index)
	v.PUT("/log", log.Index)
	v.DELETE("/log", log.Index)

	v.GET("/task", task.Index)
	v.POST("/task", task.Index)
	v.PUT("/task", task.Index)
	v.DELETE("/task", task.Index)

	v.GET("/user", user.Index)
	v.POST("/user", user.Index)
	v.PUT("/user", user.Index)
	v.DELETE("/user", user.Index)

	return e, nil
}
