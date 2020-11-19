package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/iamaul/game-currency/config"
	"github.com/iamaul/game-currency/config/database"

	cdh "github.com/iamaul/game-currency/app/currency/delivery/http"
	cr "github.com/iamaul/game-currency/app/currency/repository"
	cu "github.com/iamaul/game-currency/app/currency/usecase"

	"github.com/iamaul/game-currency/app/middleware"

	"github.com/labstack/echo"
)

func main() {
	config := config.NewConfig()

	connection, err := database.ConnectDatabase(config)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	appMiddl := middleware.InitAppMiddleware(config.AppName)
	e.Use(appMiddl.CORS)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusAccepted,
			"message": "iamaul (Game Currency - BE TEST)",
		})
	})

	currencyRepo := cr.NewCurrencyRepository(connection.SQL)

	timeoutContext := time.Duration(2) * time.Second

	currencyUcase := cu.NewCurrencyUsecase(currencyRepo, timeoutContext)

	cdh.NewCurrencyHandler(e, currencyUcase)

	log.Fatal(e.Start(fmt.Sprintf(`%s`, config.AppPort)))
}
