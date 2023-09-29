package main

import (
	config "cerpengolang/config"
	"cerpengolang/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// loadEnv()
	config.InitDatabase()
	e := echo.New()
	routes.InitRoutes(e)
	//e.Start(":" + os.Getenv("PORT"))
	e.Start(":2222")
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}
