package main

import (
	"ApiBase/src/api/app"
	"ApiBase/src/api/controllers"
	"os"
)

func main(){
	env := os.Getenv("GO_ENVIRONMENT")
	if len(env) == 0{
		env = "development"
		os.Setenv("GO_ENVIRONMENT", env)
	}

	scope := os.Getenv("SCOPE")
	if len(scope) == 0{
		scope = "local"
		os.Setenv("SCOPE", scope)
	}

	router, err := app.Start()
	if err != nil{
		panic(err)
	}

	controllers.SetURLMappgins(router)

	router.Run(":8080")
}