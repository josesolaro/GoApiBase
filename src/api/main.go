package main

import (
	"ApiBase/src/api/app"
	"ApiBase/src/api/controllers"
	"ApiBase/src/api/utils"
	"os"
)

func main(){
	log := utils.Logger{}
	log.Init()
	log.SetLogLevel("info")

	env := os.Getenv("GO_ENVIRONMENT")
	if len(env) == 0{
		env = "development"
		os.Setenv("GO_ENVIRONMENT", env)
	}
	log.Info("Environment set to: %v",env)

	scope := os.Getenv("SCOPE")
	if len(scope) == 0{
		scope = "local"
		os.Setenv("SCOPE", scope)
	}
	log.Info("Scope set to: %v",scope)

	router, err := app.Start()
	if err != nil{
		log.Error("Fatal initialization error", err)
		panic(err)
	}

	controllers.SetURLMappgins(router)

	router.Run(":8080")
}