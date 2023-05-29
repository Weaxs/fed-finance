package examples

import (
	"os"
	"testing"
)

func TestJa1971(t *testing.T) {
	file, _ := os.Create("日元 1971汇率 (兑美元).html")
	ja1971(file)
}

func TestJaFrom1972To1976(t *testing.T) {
	file, _ := os.Create("日元 1972-1976汇率 (兑美元).html")
	jaFrom1972To1976(file)
}

func TestJsFrom1982To1985(t *testing.T) {
	file, _ := os.Create("日元 1982-1985汇率 (兑美元).html")
	jaFrom1982To1985(file)
}

func TestJsFrom1986To1988(t *testing.T) {
	file, _ := os.Create("日元 1986-1988汇率 (兑美元).html")
	jaFrom1986To1988(file)
}
