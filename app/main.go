package main

import (
	"flag"
	"ushas/api"
	"ushas/config"
	"ushas/database"
)

func main() {
	env := flag.String("e", "dev", "Environment")
	flag.Parse()

	config.Init(*env)
	database.Init(false)
	if err := api.Init(); err != nil {
		panic(err)
	}
}
