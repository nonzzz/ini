package ini

import (
	"testing"

	"github.com/nonzzz/ini/internal/test"
)

func expectJson(t *testing.T, ini *Ini, expect string) {
	t.Helper()
	t.Run("expect Json", func(t *testing.T) {
		test.AssertEqual(t, string(ini.Marshal2Json()), expect)
	})
}

func TestParse(t *testing.T) {
	fixture := `
	 address = 127.0.0.1

	 [info] #login info
	 account = 123 #comment1
	 password = 456 ; comment2
	`

	ini := New().Parse(fixture)
	expect := `{"address":"127.0.0.1","info":{"account":"123","password":"456"}}`
	expectJson(t, ini, expect)
}

func TestLoadFile(t *testing.T) {
	ini := New().LoadFile("./fixture.ini")
	expect := `{"node1#123":{"a":"1","b":"2","c":"3","d":"4"},"node2#456":{}}`
	expectJson(t, ini, expect)
}
