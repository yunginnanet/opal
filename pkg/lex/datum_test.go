package lex

import "testing"

func TestLookup(t *testing.T) {
	type test struct {
		i int
		r rune
	}
	tests := []test{
		{0, ' '},
		{1, '='},
		{2, '{'},
		{3, '}'},
		{4, '('},
		{5, ')'},
		{6, ';'},
		{7, ','},
		{8, '+'},
		{9, '-'},
	}
	for _, tst := range tests {
		if rune(lookup(tst.i)) != tst.r {
			t.Errorf("lookup(%d) = '%s', wanted '%s'", tst.i, string(lookup(tst.i)), string(tst.r))
			continue
		}
		t.Logf("lookup(%d) = '%s'", tst.i, string(lookup(tst.i)))
	}
}
