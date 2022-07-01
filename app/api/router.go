package api

import (
	"net/http"
	"os"
	"ushas/config"
	"ushas/controllers"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// NewRouter : Return pointer to router struct.
func NewRouter() (*echo.Echo, error) {
	c := config.GetConfig()

	e := echo.New()
	e.HideBanner = true
	// If you would like to print access log by using this backend API,
	// you can use this logger as middleware.
	// (It should be printed by web server like nginx.)
	// e.Use(accessLogger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: c.GetStringSlice("server.cors"),
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
		},
	}))

	// e.Logger.SetOutput(os.Stderr)
	if env := c.GetString("env"); env == "dev" {
		e.Logger.SetHeader("[${level}]${message}")
		e.Logger.SetOutput(os.Stdout)
	}

	e.HTTPErrorHandler = httpErrorHandler

	e.Validator = &Validator{validator: validator.New()}

	answer := controllers.NewAnswerController()
	log := controllers.NewLogController()
	serp := controllers.NewSERPController()
	task := controllers.NewTaskController()
	user := controllers.NewUserController()

	api := e.Group("/api")
	api.POST("/users", user.Create)

	v := api.Group("/" + c.GetString("server.version"))

	v.GET("/users", user.List)
	v.GET("/users/:id", user.GetByID)
	v.GET("/:uid", user.GetByUID)
	v.PUT("/users/:id", user.Update)
	v.DELETE("/users/:id", user.Delete)

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

	v.GET("/serp/:id", serp.ListSERP)
	v.GET("/serp/icon/:id", serp.ListSERPWithIcon)
	v.GET("/serp/ratio/:id", serp.ListSERPWithRatio)

	v.GET("/tasks", task.List)
	v.GET("/tasks/:id", task.GetByID)
	v.POST("/tasks", task.Create)
	v.PUT("/tasks/:id", task.Update)
	v.DELETE("/tasks/:id", task.Delete)

	return e, nil
}
