package prefix

import "testing"

func TestHasPrefix(t *testing.T) {
	var tests = []struct {
		input string
		pre   string
		want  bool
	}{
		{"","",true},
		{"","d",false},
		{"fdsf","",true},
		{"abcd","ab",true},
		{"abcd","abcd",true},
		{"abcd","av",false},
		{"abcd","Ab",false},
		{"abcd","abcde",false},
		{"abcd","a b",false},
		{"abcd"," ab",false},
		{"abcd","ab ",false},
		{"a b","ab",false},
	}

	for _,test := range tests {
		got := HasPrefix(test.input, test.pre)
		if got != test.want {
			t.Errorf("HasPrefix(%q,%q) = %v",test.input, test.pre, got)
		}
	}
}

