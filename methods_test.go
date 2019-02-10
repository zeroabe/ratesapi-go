package ratesapi_go

import "testing"

func TestLatestByBaseAndSymbols(t *testing.T) {
	t.Run("Succesfuly rated", func(t *testing.T) {
		base := "USD"
		symbolA := "EUR"
		symbolB := "RUB"
		resp, err := LatestByBaseAndSymbols(base, symbolA, symbolB)
		if err != nil {
			t.Error(err)
		}
		if resp.Rates == nil {
			t.Error("Can't get rates")
		}
	})
}
