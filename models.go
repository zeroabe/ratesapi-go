package ratesapi_go

// Default response struct
type RatesResponse struct {
	Base  string             `json:"base"`
	Date  RatesDate          `json:"date"`
	Rates map[string]float64 `json:"rates"`
}
