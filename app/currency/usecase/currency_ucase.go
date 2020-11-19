package usecase

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/iamaul/game-currency/app/currency"
	"github.com/iamaul/game-currency/app/models"
)

type currencyUsecase struct {
	currencyRepo currency.Repository
	ctxTimeout   time.Duration
}

// NewCurrencyUsecase is a type of currencyUsecase instance
func NewCurrencyUsecase(cr currency.Repository, timeout time.Duration) currency.Usecase {
	return &currencyUsecase{
		currencyRepo: cr,
		ctxTimeout:   timeout,
	}
}

func (cu *currencyUsecase) FetchCurrency(c context.Context) (res []models.Currency, err error) {
	ctx, cancel := context.WithTimeout(c, cu.ctxTimeout)
	defer cancel()

	res, err = cu.currencyRepo.FetchCurrency(ctx)
	if err != nil {
		return nil, err
	}

	return
}

func (cu *currencyUsecase) GetCurrencyByID(c context.Context, id int64) (res models.Currency, err error) {
	ctx, cancel := context.WithTimeout(c, cu.ctxTimeout)
	defer cancel()

	res, err = cu.currencyRepo.GetCurrencyByID(ctx, id)
	if err != nil {
		return
	}

	return
}

func (cu *currencyUsecase) GetConversionRateByCurrencyID(c context.Context, rate uint64, from int64, to int64) (res models.Rate, err error) {
	ctx, cancel := context.WithTimeout(c, cu.ctxTimeout)
	defer cancel()

	res, err = cu.currencyRepo.GetConversionRateByCurrencyID(ctx, rate, from, to)
	if err != nil {
		return
	}

	return
}

func (cu *currencyUsecase) GetConversionCurrencyByCurrencyID(c context.Context, from int64, to int64, amount uint64, result uint64) (res models.ConvertCurrency, err error) {
	ctx, cancel := context.WithTimeout(c, cu.ctxTimeout)
	defer cancel()

	res, err = cu.currencyRepo.GetConversionCurrencyByCurrencyID(ctx, from, to, amount, result)
	if err != nil {
		return
	}

	return
}

func (cu *currencyUsecase) GetCurrencyByName(c context.Context, name string) (res models.Currency, err error) {
	ctx, cancel := context.WithTimeout(c, cu.ctxTimeout)
	defer cancel()

	res, err = cu.currencyRepo.GetCurrencyByName(ctx, name)
	if err != nil {
		return
	}

	return
}

func (cu *currencyUsecase) StoreCurrency(c context.Context, currency *models.Currency) (err error) {
	ctx, cancel := context.WithTimeout(c, cu.ctxTimeout)
	defer cancel()

	exist, _ := cu.GetCurrencyByName(ctx, currency.Name)
	if exist != (models.Currency{}) {
		return errors.New("That currency is already exist")
	}

	err = cu.currencyRepo.StoreCurrency(ctx, currency)

	return
}

func (cu *currencyUsecase) StoreConversionRate(c context.Context, rate *models.Rate) (err error) {
	ctx, cancel := context.WithTimeout(c, cu.ctxTimeout)
	defer cancel()

	from, _ := cu.GetCurrencyByID(ctx, rate.CurrencyIDFrom)
	if from == (models.Currency{}) {
		return errors.New("The currency_id_from field that you've entered does not exist")
	}

	to, _ := cu.GetCurrencyByID(ctx, rate.CurrencyIDTo)
	if to == (models.Currency{}) {
		return errors.New("The currency_id_to field that you've entered does not exist")
	}

	exist, _ := cu.GetConversionRateByCurrencyID(ctx, rate.Rate, rate.CurrencyIDFrom, rate.CurrencyIDTo)
	if exist != (models.Rate{}) {
		log.Println(exist)
		return errors.New("That conversion rate is already exists")
	}

	return cu.currencyRepo.StoreConversionRate(ctx, rate)
}

func (cu *currencyUsecase) StoreConversionCurrency(c context.Context, convert *models.ConvertCurrency) (err error) {
	ctx, cancel := context.WithTimeout(c, cu.ctxTimeout)
	defer cancel()

	from, _ := cu.GetCurrencyByID(ctx, convert.CurrencyIDFrom)
	if from == (models.Currency{}) {
		return errors.New("The currency_id_from field that you've entered does not exist")
	}

	to, _ := cu.GetCurrencyByID(ctx, convert.CurrencyIDTo)
	if to == (models.Currency{}) {
		return errors.New("The currency_id_to field that you've entered does not exist")
	}

	switch {
	case from.Name == "Knut" && to.Name == "Sickle":
		convert.Result = convert.Amount / 29
	case from.Name == "Sickle" && to.Name == "Knut":
		convert.Result = convert.Amount * 29
	case from.Name == "Knut" && to.Name == "Galleon":
		convert.Result = convert.Amount / 493
	case from.Name == "Galleon" && to.Name == "Sickle":
		convert.Result = convert.Amount * 17
	case from.Name == "Sickle" && to.Name == "Galleon":
		convert.Result = convert.Amount / 17
	case from.Name == "Galleon" && to.Name == "Knut":
		convert.Result = convert.Amount * 493
	default:
		return errors.New("No result")
	}

	exist, _ := cu.GetConversionCurrencyByCurrencyID(ctx, convert.CurrencyIDFrom, convert.CurrencyIDTo, convert.Amount, convert.Result)
	if exist != (models.ConvertCurrency{}) {
		return
	}

	return cu.currencyRepo.StoreConversionCurrency(ctx, convert)
}
