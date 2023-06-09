package ini

import (
	"os"
	"testing"
)

var input, _ = os.ReadFile("./str.ini")

func BenchmarkIniParse(b *testing.B) {

	s := string(input)

	for i := 0; i < b.N; i++ {
		New().Parse(s)
	}

}

func BenchmarkIniPrinter(b *testing.B) {
	s := string(input)

	for i := 0; i < b.N; i++ {
		New().Parse(s).Printer() // nolint
	}
}
