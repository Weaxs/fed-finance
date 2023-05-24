/*
 * Copyright(C) 2023 Weaxs
 */

package fed_finance

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// HistoricalExchangeRate get rates data based on usa dollar from {beginYear} to {end}
func HistoricalExchangeRate(beginYear, endYear int, country string) ([]*ExchangeRate, error) {
	if endYear < beginYear {
		return nil, errors.New("end year can not less than begin year")
	}
	if endYear < 1971 {
		return nil, errors.New("can not found data before Jan,4 1971")
	}
	if beginYear > time.Now().Year() {
		return nil, errors.New("can not found data after now")
	}

	var urls []string
	if beginYear < 1990 {
		urls = append(urls, strings.Replace(through1989, "{country}", country, 1))
	}
	if beginYear < 2000 && endYear > 1989 {
		urls = append(urls, strings.Replace(from1990to99, "{country}", country, 1))
	}
	if endYear > 1999 {
		urls = append(urls, strings.Replace(from2000toPresent, "{country}", country, 1))
	}

	var exchangeRate []*ExchangeRate
	for _, url := range urls {
		dates, rates, err := historicalTable(url)
		if err != nil {
			return nil, err
		}
		for i := 0; i < len(dates); i++ {
			if dates[i].Year() < beginYear {
				continue
			}
			if dates[i].Year() > endYear {
				break
			}

			exchangeRate = append(exchangeRate, &ExchangeRate{
				Country:      country,
				MonetaryUnit: MonetaryUnit[country],
				Date:         dates[i],
				rate:         rates[i],
			})
		}
	}
	return exchangeRate, nil
}

// ConvertRateInHistoricalDay covert rate in historical day
func ConvertRateInHistoricalDay(value float64, from string, to string, historyDay time.Time) (float64, error) {
	fromRate := 1.0
	if from != UnitedStates {
		fromRates, err := HistoricalExchangeRate(historyDay.Year(), historyDay.Year(), from)
		if err != nil {
			return 0, err
		}
		for _, rate := range fromRates {
			if rate.Date.Equal(historyDay) {
				fromRate = rate.rate
				break
			} else if rate.Date.After(historyDay) {
				return 0, fmt.Errorf("can not found contruny[%v] exchage rate in %v", from, historyDay)
			}
		}
	}

	toRate := 1.0
	if to != UnitedStates {
		toRates, err := HistoricalExchangeRate(historyDay.Year(), historyDay.Year(), to)
		if err != nil {
			return 0, err
		}
		for _, rate := range toRates {
			if rate.Date.Equal(historyDay) {
				toRate = rate.rate
				break
			} else if rate.Date.After(historyDay) {
				return 0, fmt.Errorf("can not found contruny[%v] exchage rate in %v", to, historyDay)
			}
		}
	}

	return value / fromRate * toRate, nil
}

func historicalTable(url string) (dates []time.Time, rates []float64, err error) {
	response, err := http.Get(url)
	if err != nil {
		err = fmt.Errorf("request failed: %w", err)
		return
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	htmlBytes, err := io.ReadAll(response.Body)
	if err != nil {
		err = fmt.Errorf("read response body error: %w", err)
		return
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(htmlBytes)))
	if err != nil {
		err = fmt.Errorf("parse body to html error: %w", err)
		return
	}

	table := doc.Find("table[summary]")
	if table.Length() == 0 {
		table = doc.Find("table[title]").Find("tbody")
	}
	if table.Length() == 0 {
		err = fmt.Errorf("can not found data in %v", url)
		return
	}

	table.Find("tr").Each(func(rowIdx int, row *goquery.Selection) {
		th := row.Find("th")
		row.Find("td").Each(func(colIdx int, td *goquery.Selection) {
			tdText := strings.TrimSpace(td.Text())
			thText := strings.TrimSpace(th.Text())

			rate, ignoredErr := strconv.ParseFloat(tdText, 4)
			if ignoredErr != nil {
				// no data (ND)
				return
			}

			if len(thText) == 8 {
				thText = "0" + thText
			}
			date, ignoredErr := time.Parse(FedDate, thText)
			if ignoredErr != nil {
				return
			}
			rates = append(rates, rate)
			dates = append(dates, date)
		})
	})

	return
}
