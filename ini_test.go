package ini

import (
	"testing"

	"github.com/nonzzz/ini/internal/test"
)

func TestMarshal2Json(t *testing.T) {
	i, _ := New().LoadFile("./case/str.ini")
	s := string(i.Marshal2Json())
	test.AssertEqual(t, s, `{"p":"0","s1":{"a":"1","b":"2"},"s2":{"c":"4","d":"5"}}`)
}
