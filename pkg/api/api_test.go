package api

import (
	"testing"
)

func TestJson(t *testing.T) {

	txt := `
	name = kanno
	[node1]
	a=1
	b=2
	c=3
	#comment1
	[node2]
	a = 2
	b = 4
	c = 6
	d= 8

	;comment2
	`
	ini := New()
	expected := `{"name":"kanno","node1":{"a":"1","b":"2","c":"3"},"node2":{"a":"2","b":"4","c":"6","d":"8"}}`
	r := string(ini.Parse(txt).Marshl2Json())
	if r != expected {
		t.Fatalf("%s != %s", r, expected)
	}
}

func TestLoad(t *testing.T) {
	ini := New()
	r := ini.LoadFile("../../fixture.ini").Marshl2Json()
	expected := `{"node1#123":{"a":"1","b":"2","c":"3","d":"4"},"node2#456":{}}`
	if string(r) != expected {
		t.Fatalf("%s != %s", r, expected)
	}
}
