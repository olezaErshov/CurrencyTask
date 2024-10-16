package entity

type Currency struct {
	Rate float32 `json:"rate"`
	Date string  `json:"date"`
}

type CurrencyResponse struct {
	Date  string             `json:"date"`
	Rates map[string]float32 `json:"rub"`
}
