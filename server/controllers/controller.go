package controllers

import (
	"math/rand"
	"net/http"

	"github.com/Sahil2k07/React-Go/configs"
	"github.com/Sahil2k07/React-Go/models"
	"github.com/labstack/echo/v4"
)

func MapControllers(e *echo.Echo) {
	e.GET("/api/example", func(c echo.Context) error {

		messages := []string{
			"Hello, world!",
			"Echo is awesome!",
			"Go makes concurrency easy.",
			"You're doing great!",
			"Keep pushing forward!",
			"Random message here.",
			"Stay curious.",
			"Code and coffee!",
			"Never stop learning.",
			"One step at a time.",
		}

		randomIndex := rand.Intn(len(messages))
		randomMessage := messages[randomIndex]

		// Insert into database
		msg := models.Message{Content: randomMessage}
		configs.DB.Create(&msg)

		return c.String(http.StatusOK, randomMessage)
	})
}
