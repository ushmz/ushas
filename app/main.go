package main

import (
	"flag"
	"ushas/config"
	"ushas/database"
	"ushas/server"
)

func main() {
	env := flag.String("e", "dev", "Environment")
	flag.Parse()

	config.Init(*env)
	database.Init(false)
	if err := server.Init(); err != nil {
		panic(err)
	}
}
