package examples

import (
	"os"
	"testing"
)

func TestUk1971(t *testing.T) {
	file, _ := os.Create("英镑 1971汇率 (兑美元).html")
	uk1971(file)
}

func TestUkFrom1972To1976(t *testing.T) {
	file, _ := os.Create("英镑 1972-1976汇率 (兑美元).html")
	ukFrom1972To1976(file)
}

func TestUkFrom1982To1985(t *testing.T) {
	file, _ := os.Create("英镑 1982-1985汇率 (兑美元).html")
	ukFrom1982To1985(file)
}
func TestUkFrom1986To1988(t *testing.T) {
	file, _ := os.Create("英镑 1986-1988汇率 (兑美元).html")
	ukFrom1986To1988(file)
}
