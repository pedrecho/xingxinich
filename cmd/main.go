package main

import (
	"flag"
	"xingxinich/internal/app"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "", "config path")
	flag.Parse()

	if err := app.Run(configPath); err != nil {
		panic(err)
	}
}
