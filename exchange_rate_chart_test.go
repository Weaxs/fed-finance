/*
 * Copyright(C) 2023 Weaxs
 */

package fed_finance

import (
	"os"
	"testing"
)

func TestExchangeRateLine(t *testing.T) {
	f, _ := os.Create("chart/GE-UK_rates_82-87.html")

	beginYear := 1982
	endYear := 1987
	geRates, _ := HistoricalExchangeRate(beginYear, endYear, Germany)
	ukRates, _ := HistoricalExchangeRate(beginYear, endYear, UnitedKingdom)

	var exchangeRates [][]ExchangeRate
	exchangeRates = append(exchangeRates, geRates)
	exchangeRates = append(exchangeRates, ukRates)

	_ = ExchangeRateLine(exchangeRates, f)

}
