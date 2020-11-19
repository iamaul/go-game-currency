package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/iamaul/game-currency/app/currency"
	"github.com/iamaul/game-currency/app/models"
	"github.com/sirupsen/logrus"
)

type currencyRepository struct {
	Database *sql.DB
}

// NewCurrencyRepository is a type of currencyRepository instance
func NewCurrencyRepository(connection *sql.DB) currency.Repository {
	return &currencyRepository{connection}
}

func (cr *currencyRepository) fetchCurrency(ctx context.Context, query string, args ...interface{}) (payload []models.Currency, err error) {
	rows, err := cr.Database.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer rows.Close()

	payload = make([]models.Currency, 0)
	for rows.Next() {
		data := models.Currency{}
		err = rows.Scan(
			&data.ID,
			&data.Name,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		payload = append(payload, data)
	}

	return payload, nil
}

func (cr *currencyRepository) fetchConversionRate(ctx context.Context, query string, args ...interface{}) (payload []models.Rate, err error) {
	rows, err := cr.Database.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer rows.Close()

	payload = make([]models.Rate, 0)
	for rows.Next() {
		data := models.Rate{}
		err = rows.Scan(
			&data.ID,
			&data.CurrencyIDFrom,
			&data.CurrencyIDTo,
			&data.Rate,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		payload = append(payload, data)
	}

	return payload, nil
}

func (cr *currencyRepository) fetchConversionCurrency(ctx context.Context, query string, args ...interface{}) (payload []models.ConvertCurrency, err error) {
	rows, err := cr.Database.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer rows.Close()

	payload = make([]models.ConvertCurrency, 0)
	for rows.Next() {
		data := models.ConvertCurrency{}
		err = rows.Scan(
			&data.ID,
			&data.CurrencyIDFrom,
			&data.CurrencyIDTo,
			&data.Amount,
			&data.Result,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		payload = append(payload, data)
	}

	return payload, nil
}

func (cr *currencyRepository) FetchCurrency(ctx context.Context) (res []models.Currency, err error) {
	query := `SELECT id, name FROM currencies ORDER BY created_at DESC`

	res, err = cr.fetchCurrency(ctx, query)
	if err != nil {
		return nil, err
	}

	if len(res) > 0 {
		return
	}

	return nil, errors.New("Data is empty")
}

func (cr *currencyRepository) GetCurrencyByID(ctx context.Context, id int64) (res models.Currency, err error) {
	query := `SELECT id, name FROM currencies WHERE id=?`

	result, err := cr.fetchCurrency(ctx, query, id)
	if err != nil {
		return models.Currency{}, err
	}

	if len(result) > 0 {
		res = result[0]
	} else {
		return res, errors.New("Data not found")
	}

	return
}

func (cr *currencyRepository) GetConversionRateByCurrencyID(ctx context.Context, rate uint64, from int64, to int64) (res models.Rate, err error) {
	query := `SELECT id, currency_id_from, currency_id_to, rate 
			FROM conversion_rates WHERE rate=? AND 
			currency_id_from=? AND currency_id_to=?`

	result, err := cr.fetchConversionRate(ctx, query, rate, from, to)
	if err != nil {
		return models.Rate{}, err
	}

	if len(result) > 0 {
		res = result[0]
	} else {
		return res, errors.New("Data not found")
	}

	return
}

func (cr *currencyRepository) GetConversionCurrencyByCurrencyID(ctx context.Context, from int64, to int64, amount uint64, result uint64) (res models.ConvertCurrency, err error) {
	query := `SELECT id, currency_id_from, currency_id_to, amount, result 
			FROM conversion_currencies WHERE currency_id_from=? AND currency_id_to=?
			AND amount=? AND result=?`

	results, err := cr.fetchConversionCurrency(ctx, query, from, to, amount, result)
	if err != nil {
		return models.ConvertCurrency{}, err
	}

	if len(results) > 0 {
		res = results[0]
	} else {
		return res, errors.New("Data not found")
	}

	return
}

func (cr *currencyRepository) GetCurrencyByName(ctx context.Context, name string) (res models.Currency, err error) {
	query := `SELECT id, name FROM currencies WHERE name=?`

	result, err := cr.fetchCurrency(ctx, query, name)
	if err != nil {
		return
	}

	if len(result) > 0 {
		res = result[0]
	} else {
		return res, errors.New("Data not found")
	}

	return
}

func (cr *currencyRepository) StoreCurrency(ctx context.Context, currency *models.Currency) (err error) {
	query := `INSERT currencies SET name=?`
	stmt, err := cr.Database.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, currency.Name)
	if err != nil {
		return
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}

	currency.ID = lastID

	return
}

func (cr *currencyRepository) StoreConversionRate(ctx context.Context, rate *models.Rate) (err error) {
	query := `INSERT conversion_rates SET currency_id_from=?, currency_id_to=?, rate=?`
	stmt, err := cr.Database.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, rate.CurrencyIDFrom, rate.CurrencyIDTo, rate.Rate)
	if err != nil {
		return
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}

	rate.ID = lastID

	return
}

func (cr *currencyRepository) StoreConversionCurrency(ctx context.Context, convert *models.ConvertCurrency) (err error) {
	query := `INSERT conversion_currencies SET currency_id_from=?, currency_id_to=?, amount=?, result=?`
	stmt, err := cr.Database.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, convert.CurrencyIDFrom, convert.CurrencyIDTo, convert.Amount, convert.Result)
	if err != nil {
		return
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}

	convert.ID = lastID

	return
}
