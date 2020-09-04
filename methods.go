package ratesapi_go

import "strings"

const latestRoute = "latest"

// Returns RatesResponse with specified base and destination currencies (symbols)
func LatestByBaseAndSymbols(base string, symb ...string) (response RatesResponse, err error) {
	c := NewRatesClient()
	symbols := strings.Join(symb, ",")
	params := map[string]string{
		"base":    base,
		"symbols": symbols,
	}
	err = c.Get(baseURL+latestRoute, params, &response)
	return
}
