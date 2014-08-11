package main

import "testing"

var testData = []struct {
	value1 string
	value2 string
	out    string
}{
	{"A", "B", "AB"},
	{"", "B", "B"},
	{"A", "", "A"},
	{"", "", ""},
	{"ABC", "HOGE", "ABCHOGE"},
}

func TestAppendString(t *testing.T) {
	for _, td := range testData {
		s := AppendString(td.value1, td.value2)
		if s != td.out {
			t.Errorf("Sprintf(%q, %q) = <%s> want <%s>", td.value1, td.value2, s, td.out)
		}

	}
}
