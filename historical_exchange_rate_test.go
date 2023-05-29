package fed_finance

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestHistoricalExchangeRate(t *testing.T) {
	rates, err := HistoricalExchangeRate(1985, 1995, Japan)
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, 1985, rates[1].Date.Year())
	assert.Equal(t, 1995, rates[len(rates)-1].Date.Year())
}

func TestCovertRateInHistoricalDay(t *testing.T) {
	historyDay, _ := time.Parse("2006-01-02", "1994-01-04")
	value := 100.0
	yuan, err := ConvertRateInHistoricalDay(value, Japan, China, historyDay)
	if err != nil {
		return
	}
	assert.Equal(t, value/112.75*8.7217, yuan)
	fmt.Println(fmt.Sprintf("in 1994-01-04, 100 %v can covert %.4f %v ",
		MonetaryUnit[Japan], yuan, MonetaryUnit[China]))
}
