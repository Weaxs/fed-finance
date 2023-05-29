/*
 * Copyright(C) 2023 Weaxs
 */

package examples

import (
	"os"
	"testing"
)

func TestGeFrom1971To1972(t *testing.T) {
	file, _ := os.Create("德国马克 1971汇率 (兑美元).html")
	ge1971(file)
}

func TestName(t *testing.T) {
	file, _ := os.Create("德国马克 1972-1976汇率 (兑美元).html")
	geFrom1972To1976(file)
}
