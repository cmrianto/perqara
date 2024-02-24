package main

import (
	"log"
	"os"
	"perqara/cmd"

	"perqara/config"
)

// @title           Perqara Test API Docs
// @version         1.0
// @description     Backend API Documentation for Perqara Test
// @termsOfService  http://swagger.io/terms/

// @contact.name   Chris Muladi Rianto
// @contact.url    http://www.swagger.io/support
// @contact.email  chrisrianto84@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath
func main() {
	cfg, err := config.Setup()
	if err != nil {
		log.Fatal("Cannot load config ", err.Error())
	}

	if cmd.Cli(cfg).Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
