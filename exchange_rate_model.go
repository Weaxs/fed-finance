/*
 * Copyright(C) 2023 Weaxs
 */

package fed_finance

import "time"

// Exchange Rates

const (
	FedDate = "02-Jan-06"

	// Historical Data (Beginning Jan. 4, 1971 to 1989)
	// https://www.federalreserve.gov/releases/h10/hist/default1989.htm
	through1989 = "https://www.federalreserve.gov/releases/h10/hist/dat89_{country}.htm"
	// Historical Data (1990-99)
	// https://www.federalreserve.gov/releases/h10/hist/default1999.htm
	from1990to99 = "https://www.federalreserve.gov/releases/h10/hist/dat96_{country}.htm"
	// Historical Data (2000 to Present)
	// https://www.federalreserve.gov/releases/h10/hist/default.htm
	from2000toPresent = "https://www.federalreserve.gov/releases/h10/hist/dat00_{country}.htm"
	// Current Data
	// https://www.federalreserve.gov/releases/h10/current/default.htm
	current = "https://www.federalreserve.gov/releases/h10/current/default.htm"
)

type ExchangeRate struct {
	Country      string
	MonetaryUnit string
	Date         time.Time
	Value        float64
}
