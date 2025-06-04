package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"grimoire-api/db"

	"github.com/joho/godotenv"

	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not founded")
	}

	db.Connect()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server running!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
