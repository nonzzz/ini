package ini

import "testing"

var txt = "[s1]\n a=0\b=2 #comment"

func MakeIniInstance() {
	New().Parse(txt)
}

func BenchmarkParse(b *testing.B) {

	for i := 0; i < b.N; i++ {
		MakeIniInstance()
	}

}
