package http

import (
	"context"
	"net/http"

	"github.com/iamaul/game-currency/app/currency"
	"github.com/iamaul/game-currency/app/models"

	"github.com/labstack/echo"
)

type response struct {
	Code    int         `json:"code,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
}

// CurrencyHandler HTTP handler for currency
type CurrencyHandler struct {
	currencyUsecase currency.Usecase
}

// NewCurrencyHandler is an endpoint resources for currency
func NewCurrencyHandler(e *echo.Echo, cu currency.Usecase) {
	handler := &CurrencyHandler{
		currencyUsecase: cu,
	}

	api := e.Group("/api")
	api.GET("/currencies", handler.GetAllCurrency)
	api.POST("/currencies", handler.InsertCurrency)
	api.POST("/conversion/rates", handler.InsertConversionRate)
	api.POST("/conversion/currencies", handler.ConvertCurrency)
}

// GetAllCurrency will returns the list of all currencies
func (ch *CurrencyHandler) GetAllCurrency(c echo.Context) error {
	ctx := c.Request().Context()

	data, err := ch.currencyUsecase.FetchCurrency(ctx)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &response{
			Code:  http.StatusBadRequest,
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &response{
		Code:    http.StatusOK,
		Data:    data,
		Message: "Successfully fetched currencies",
	})
}

// InsertCurrency will store the currency data by given request raw body
func (ch *CurrencyHandler) InsertCurrency(c echo.Context) error {
	var currency models.Currency

	err := c.Bind(&currency)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &response{
			Code:  http.StatusUnprocessableEntity,
			Error: err.Error(),
		})
	}

	if currency.Name == "" {
		return c.JSON(http.StatusBadRequest, &response{
			Code:  http.StatusBadRequest,
			Error: "The field name is required",
		})
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = ch.currencyUsecase.StoreCurrency(ctx, &currency)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &response{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, &response{
		Code:    http.StatusCreated,
		Data:    currency,
		Message: "Successfully created currency",
	})
}

// InsertConversionRate will store the conversion by given a currency rate
func (ch *CurrencyHandler) InsertConversionRate(c echo.Context) error {
	var rate models.Rate

	err := c.Bind(&rate)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &response{
			Code:  http.StatusUnprocessableEntity,
			Error: err.Error(),
		})
	}

	if rate.Rate == uint64(0) {
		return c.JSON(http.StatusBadRequest, &response{
			Code:  http.StatusBadRequest,
			Error: "The field rate cannot be zero",
		})
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = ch.currencyUsecase.StoreConversionRate(ctx, &rate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &response{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, &response{
		Code:    http.StatusCreated,
		Data:    rate,
		Message: "Successfully added conversion rate",
	})
}

// ConvertCurrency will convert the value of selected currencies
func (ch *CurrencyHandler) ConvertCurrency(c echo.Context) error {
	var currency models.ConvertCurrency

	err := c.Bind(&currency)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &response{
			Code:  http.StatusUnprocessableEntity,
			Error: err.Error(),
		})
	}

	if currency.Amount == uint64(0) {
		return c.JSON(http.StatusBadRequest, &response{
			Code:  http.StatusBadRequest,
			Error: "The field amount cannot be zero",
		})
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = ch.currencyUsecase.StoreConversionCurrency(ctx, &currency)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &response{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, &response{
		Code:    http.StatusCreated,
		Data:    currency,
		Message: "Successfully converted currency",
	})
}
