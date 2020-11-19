package currency

import (
	"context"

	"github.com/iamaul/game-currency/app/models"
)

// Repository represent the currency's repository
type Repository interface {
	FetchCurrency(ctx context.Context) (res []models.Currency, err error)
	GetCurrencyByID(ctx context.Context, id int64) (models.Currency, error)
	GetConversionRateByCurrencyID(ctx context.Context, rate uint64, from int64, to int64) (models.Rate, error)
	GetConversionCurrencyByCurrencyID(ctx context.Context, from int64, to int64, amount uint64, result uint64) (models.ConvertCurrency, error)
	GetCurrencyByName(ctx context.Context, name string) (models.Currency, error)
	StoreCurrency(ctx context.Context, currency *models.Currency) error
	StoreConversionRate(ctx context.Context, rate *models.Rate) error
	StoreConversionCurrency(ctx context.Context, convert *models.ConvertCurrency) error
}
