package currency

import (
	"context"

	"github.com/iamaul/game-currency/app/models"
)

// Usecase represent the currency's usecases
type Usecase interface {
	FetchCurrency(ctx context.Context) ([]models.Currency, error)
	GetCurrencyByID(ctx context.Context, id int64) (models.Currency, error)
	GetConversionRateByCurrencyID(ctx context.Context, rate uint64, from int64, to int64) (models.Rate, error)
	GetConversionCurrencyByCurrencyID(ctx context.Context, from int64, to int64, amount uint64, result uint64) (models.ConvertCurrency, error)
	GetCurrencyByName(ctx context.Context, name string) (models.Currency, error)
	StoreCurrency(context.Context, *models.Currency) error
	StoreConversionRate(context.Context, *models.Rate) error
	StoreConversionCurrency(context.Context, *models.ConvertCurrency) error
}
