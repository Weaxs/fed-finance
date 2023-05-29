/*
 * Copyright(C) 2023 Weaxs
 */

package examples

import (
	"os"
	"testing"
)

func TestUkFrom1971To1972(t *testing.T) {
	file, _ := os.Create("英镑 1971汇率 (兑美元).html")
	uk1971(file)
}

func TestUkFrom1973To1976(t *testing.T) {
	file, _ := os.Create("英镑 1972-1976汇率 (兑美元).html")
	ukFrom1972To1976(file)
}
