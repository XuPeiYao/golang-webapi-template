package main

import (
	"fmt"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample api-server celler api-server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      127.0.0.1:8080

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	fmt.Println("Starting...")

	app := InitializeFiberApp()

	err := app.Run()

	if err != nil {
		panic(err)
	}

	fmt.Println("Stopped.")
}
