package main

import (
	"os"

	"github.com/System-Glitch/goyave-blog-example/http/route"

	_ "github.com/System-Glitch/goyave-blog-example/http/validation"

	"github.com/System-Glitch/goyave/v3"
	_ "github.com/System-Glitch/goyave/v3/database/dialect/mysql"
)

func main() {
	// This is the entry point of your application.
	if err := goyave.Start(route.Register); err != nil {
		os.Exit(err.(*goyave.Error).ExitCode)
	}
}