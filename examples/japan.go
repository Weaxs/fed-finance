package examples

import (
	fed_finance "github.com/Weaxs/fed-finance"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"os"
	"strings"
)

const jaTheme = types.ThemeWalden

func ja1971(f *os.File) {
	country := fed_finance.Japan
	rates, _ := fed_finance.HistoricalExchangeRate(1971, 1971, country)
	var dates []string
	var lines []opts.LineData
	for _, rate := range rates {
		date := rate.Date.Format(fed_finance.ChartDate)
		dates = append(dates, date)

		line := opts.LineData{
			Value: rate.Value,
			Name:  date,
		}
		lines = append(lines, line)
	}

	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: jaTheme}),
		charts.WithXAxisOpts(opts.XAxis{
			Name:      "时间",
			Show:      false,
			AxisLabel: &opts.AxisLabel{Show: false, Interval: "6", ShowMinLabel: true, ShowMaxLabel: true},
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name:        strings.ToLower(fed_finance.MonetaryUnit[country]) + "/usd",
			Show:        true,
			Max:         360,
			Min:         240,
			SplitNumber: 8,
			SplitLine:   &opts.SplitLine{Show: true},
		}),
		//charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemePurplePassion}),
		charts.WithTitleOpts(opts.Title{
			Title: strings.TrimSuffix(f.Name(), ".html"),
		}))

	line.SetXAxis(dates)
	line.AddSeries(rates[0].Country, lines,
		charts.WithMarkLineNameXAxisItemOpts(opts.MarkLineNameXAxisItem{Name: "停止美元兑付", XAxis: "1971-08-13"}),
		charts.WithMarkLineNameXAxisItemOpts(opts.MarkLineNameXAxisItem{Name: "史密森协议", XAxis: "1971-12-17"}),
	).SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{
		//Smooth: true,
	}))

	_ = line.Render(f)
}

func jaFrom1972To1976(f *os.File) {
	country := fed_finance.Japan
	rates, _ := fed_finance.HistoricalExchangeRate(1972, 1976, country)
	var dates []string
	var lines []opts.LineData
	for _, rate := range rates {
		date := rate.Date.Format(fed_finance.ChartDate)
		dates = append(dates, date)

		line := opts.LineData{
			Value: rate.Value,
			Name:  date,
		}
		lines = append(lines, line)
	}

	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: jaTheme}),
		charts.WithXAxisOpts(opts.XAxis{
			Name:      "时间",
			Show:      false,
			AxisLabel: &opts.AxisLabel{Show: false, Interval: "15", ShowMinLabel: true, ShowMaxLabel: true},
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name:        strings.ToLower(fed_finance.MonetaryUnit[country]) + "/usd",
			Show:        true,
			Max:         360,
			Min:         240,
			SplitNumber: 8,
			SplitLine:   &opts.SplitLine{Show: true},
		}),
		//charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemePurplePassion}),
		charts.WithTitleOpts(opts.Title{
			Title: strings.TrimSuffix(f.Name(), ".html"),
		}))

	line.SetXAxis(dates)
	line.AddSeries(rates[0].Country, lines,
		charts.WithMarkLineNameXAxisItemOpts(opts.MarkLineNameXAxisItem{Name: "日元浮动", XAxis: "1973-02-13"}),
		charts.WithMarkLineNameXAxisItemOpts(opts.MarkLineNameXAxisItem{Name: "第一次石油危机开始", XAxis: "1973-10-16"}),
		charts.WithMarkLineNameXAxisItemOpts(opts.MarkLineNameXAxisItem{Name: "第一次石油危机结束", XAxis: "1974-03-18"}),
		charts.WithMarkLineNameXAxisItemOpts(opts.MarkLineNameXAxisItem{Name: "IMF确定浮动汇率制", XAxis: "1976-01-08"}),
	).SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{
		//Smooth: true,
	}))

	_ = line.Render(f)
}

func jaFrom1982To1985(f *os.File) {
	country := fed_finance.Japan
	rates, _ := fed_finance.HistoricalExchangeRate(1982, 1985, country)
	var dates []string
	var lines []opts.LineData
	for _, rate := range rates {
		date := rate.Date.Format(fed_finance.ChartDate)
		dates = append(dates, date)

		line := opts.LineData{
			Value: rate.Value,
			Name:  date,
		}
		lines = append(lines, line)
	}

	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: jaTheme}),
		charts.WithXAxisOpts(opts.XAxis{
			Name:      "时间",
			Show:      false,
			AxisLabel: &opts.AxisLabel{Show: false, Interval: "12", ShowMinLabel: true, ShowMaxLabel: true},
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name:        strings.ToLower(fed_finance.MonetaryUnit[country]) + "/usd",
			Show:        true,
			Max:         300,
			Min:         150,
			SplitNumber: 8,
			SplitLine:   &opts.SplitLine{Show: true},
		}),
		//charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemePurplePassion}),
		charts.WithTitleOpts(opts.Title{
			Title: strings.TrimSuffix(f.Name(), ".html"),
		}))

	line.SetXAxis(dates)
	line.AddSeries(rates[0].Country, lines,
		charts.WithMarkLineNameXAxisItemOpts(opts.MarkLineNameXAxisItem{Name: "广场协议", XAxis: "1985-09-23"}),
	).SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{
		//Smooth: true,
	}))

	_ = line.Render(f)
}

func jaFrom1986To1988(f *os.File) {
	country := fed_finance.Japan
	rates, _ := fed_finance.HistoricalExchangeRate(1986, 1988, country)
	var dates []string
	var lines []opts.LineData
	for _, rate := range rates {
		date := rate.Date.Format(fed_finance.ChartDate)
		dates = append(dates, date)

		line := opts.LineData{
			Value: rate.Value,
			Name:  date,
		}
		lines = append(lines, line)
	}

	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: jaTheme}),
		charts.WithXAxisOpts(opts.XAxis{
			Name:      "时间",
			Show:      false,
			AxisLabel: &opts.AxisLabel{Show: false, Interval: "12", ShowMinLabel: true, ShowMaxLabel: true},
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name:        strings.ToLower(fed_finance.MonetaryUnit[country]) + "/usd",
			Show:        true,
			Max:         200,
			Min:         100,
			SplitNumber: 8,
			SplitLine:   &opts.SplitLine{Show: true},
		}),
		//charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemePurplePassion}),
		charts.WithTitleOpts(opts.Title{
			Title: strings.TrimSuffix(f.Name(), ".html"),
		}))

	line.SetXAxis(dates)
	line.AddSeries(rates[0].Country, lines,
		charts.WithMarkLineNameXAxisItemOpts(opts.MarkLineNameXAxisItem{Name: "卢浮宫协议", XAxis: "1987-02-23"}),
		charts.WithMarkLineNameXAxisItemOpts(opts.MarkLineNameXAxisItem{Name: "圣诞公报", XAxis: "1987-12-23"}),
	).SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{
		//Smooth: true,
	}))

	_ = line.Render(f)
}
