package api

import (
	"fmt"
	"net/http"
	"strings"
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

		return models.NewInternalError(err, strings.Join(results, ";"), i)
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

	e.HTTPErrorHandler = HTTPErrorHandler

	e.Validator = &Validator{validator: validator.New()}

	v := e.Group("/" + c.GetString("server.version"))

	answer := controllers.NewAnswerController()
	log := controllers.NewLogController()
	task := controllers.NewTaskController()
	user := controllers.NewUserController()

	v.GET("/answers", answer.List)
	v.GET("/answers/:id", answer.Get)
	v.POST("/answers", answer.Create)
	v.PUT("/answers/:id", answer.Update)
	v.DELETE("/answers/:id", answer.Delete)

	v.GET("/logs/dwell", log.Index)
	v.GET("/logs/view", log.Index)
	v.GET("/logs", log.Index)
	// The usage that getting each log doesn't expected.
	// v.GET("/log/:id", log.Index)
	v.POST("/logs/serp/dwell", log.UpsertSerpDwellTimeLog)
	v.POST("/logs/serp/event", log.CreateSerpEventLog)
	v.POST("/logs/page/dwell", log.Index)
	v.PUT("/logs/:id", log.Index)
	v.DELETE("/logs/:id", log.Index)

	v.GET("/tasks", task.List)
	v.GET("/tasks/:id", task.GetByID)
	v.POST("/tasks", task.Create)
	v.PUT("/tasks/:id", task.Update)
	v.DELETE("/tasks/:id", task.Delete)

	v.GET("/users", user.List)
	v.GET("/users/:id", user.GetByID)
	v.GET("/:uid", user.GetByUID)
	v.POST("/users", user.Create)
	v.PUT("/users/:id", user.Update)
	v.DELETE("/users/:id", user.Delete)

	return e, nil
}
