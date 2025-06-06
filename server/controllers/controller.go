package controllers

import (
	"math/rand"
	"net/http"

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

		return c.String(http.StatusOK, randomMessage)
	})
}
