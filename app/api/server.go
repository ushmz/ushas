package api

import (
	"ushas/config"
)

// Init : Gets configurations and runs API server.
func Init() error {
	c := config.GetConfig()
	r, err := NewRouter()
	if err != nil {
		return err
	}
	r.Logger.Fatal(r.Start(":" + c.GetString("server.port")))
	return nil
}
