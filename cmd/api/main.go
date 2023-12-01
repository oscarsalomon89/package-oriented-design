package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/teamcubation/pod/cmd/api/app"
	"github.com/teamcubation/pod/internal/platform/environment"
)

func main() {
	env := environment.GetFromString(os.Getenv("GO_ENVIRONMENT"))
	if env == environment.Development {
		if err := godotenv.Load(); err != nil {
			log.Fatal("error loading env file", err)
		}
	}

	dependencies, err := app.BuildDependencies(env)
	if err != nil {
		log.Fatal("error at dependencies building", err)
	}

	app := app.Build(dependencies)
	if err := app.Run(); err != nil {
		log.Fatal("error at inventory-app startup", err)
	}
}
