package models

// Currency is a model of SQL table currencies
type Currency struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// Rate is a model of SQL table conversion_rates
type Rate struct {
	ID             int64  `json:"id"`
	CurrencyIDFrom int64  `json:"currency_id_from"`
	CurrencyIDTo   int64  `json:"currency_id_to"`
	Rate           uint64 `json:"rate"`
}

// ConvertCurrency represent the convert currency model
type ConvertCurrency struct {
	ID             int64  `json:"id"`
	CurrencyIDFrom int64  `json:"currency_id_from"`
	CurrencyIDTo   int64  `json:"currency_id_to"`
	Amount         uint64 `json:"amount"`
	Result         uint64 `json:"result"`
}
