/*
 * Copyright(C) 2023 Weaxs
 */

package examples

import (
	"os"
	"testing"
)

func TestJaFrom1971To1972(t *testing.T) {
	file, _ := os.Create("日元 1971汇率 (兑美元).html")
	ja1971(file)
}

func TestJaFrom1973To1976(t *testing.T) {
	file, _ := os.Create("日元 1972-1976汇率 (兑美元).html")
	jaFrom1972To1976(file)
}
