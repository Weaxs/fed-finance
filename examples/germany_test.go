package examples

import (
	"os"
	"testing"
)

func TestGe1971(t *testing.T) {
	file, _ := os.Create("德国马克 1971汇率 (兑美元).html")
	ge1971(file)
}

func TestGeFrom1972To1976(t *testing.T) {
	file, _ := os.Create("德国马克 1972-1976汇率 (兑美元).html")
	geFrom1972To1976(file)
}

func TestGeFrom1982To1985(t *testing.T) {
	file, _ := os.Create("德国马克 1982-1985汇率 (兑美元).html")
	geFrom1982To1985(file)
}

func TestGeFrom1986To1988(t *testing.T) {
	file, _ := os.Create("德国马克 1986-1988汇率 (兑美元).html")
	geFrom1986To1988(file)
}
