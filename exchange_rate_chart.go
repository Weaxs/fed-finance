/*
 * Copyright(C) 2023 Weaxs
 */

package fed_finance

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"os"
	"sort"
	"time"
)

func ExchangeRateLine(exchangeRates [][]ExchangeRate, file *os.File) error {
	dateRates := make(map[string][]float64)
	var dates []string
	countries := make([]string, len(exchangeRates))

	for i, exchangeRate := range exchangeRates {
		countries[i] = exchangeRate[0].Country
		for _, everyRate := range exchangeRate {
			date := everyRate.Date.Format(ChartDate)
			rates, ok := dateRates[date]
			if !ok {
				rates = make([]float64, len(exchangeRates))
				dates = append(dates, date)
			}
			rates[i] = everyRate.Value
			dateRates[date] = rates
		}
	}

	// sort
	sort.Slice(dates, func(i, j int) bool {
		d1, _ := time.Parse(ChartDate, dates[i])
		d2, _ := time.Parse(ChartDate, dates[j])
		if d1.After(d2) {
			return false
		} else {
			return true
		}
	})

	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title: "Exchange Rate Line",
		}))

	line.SetXAxis(dates)

	for i, country := range countries {
		lines := make([]opts.LineData, 0)
		for _, date := range dates {
			lines = append(lines, opts.LineData{
				Value: dateRates[date][i],
				Name:  country,
			})
		}
		line.AddSeries(country, lines).
			SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	}

	err := line.Render(file)
	return err
}
