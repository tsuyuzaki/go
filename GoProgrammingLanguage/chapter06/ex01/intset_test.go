/**
 * > test run ./
 */
package intset

import (
	"testing"
)

func addTest(s *IntSet, x int) bool {
	len := s.Len()
	alreadyHas := s.Has(x)
	s.Add(x)
	if !s.Has(x) {
		return false
	}
	if !alreadyHas {
		return (s.Len() == (len + 1))
	} else {
		return s.Len() == len
	}
}

func removeTest(s *IntSet, x int) bool {
	len := s.Len()
	has := s.Has(x)
	s.Remove(x)
	if s.Has(x) {
		return false
	}
	if has {
		return (s.Len() == (len - 1))
	} else {
		return s.Len() == len
	}
}

func Test6_1(t *testing.T) {
	const maxLen = 3 * 64

	s := IntSet{}
	for i := 0; i < maxLen; i++ {
		if !addTest(&s, i) {
			t.Fatal("failed test")
		}
		if !addTest(&s, i) {
			t.Fatal("failed test")
		}
	}

	if s.Len() != maxLen {
		t.Fatal("failed test")
	}
	if s.String() != s.Copy().String() {
		t.Fatal("failed test")
	}

	for i := 0; i < maxLen; i++ {
		if !removeTest(&s, i) {
			t.Fatal("failed test")
		}
		if !removeTest(&s, i) {
			t.Fatal("failed test")
		}
	}

	if s.Len() != 0 {
		t.Fatal("failed test")
	}
	if s.String() != s.Copy().String() {
		t.Fatal("failed test")
	}

	for i := 0; i < maxLen; i++ {
		if !addTest(&s, i) {
			t.Fatal("failed test")
		}
	}
	if s.Len() != maxLen {
		t.Fatal("failed test")
	}

	s.Clear()
	if s.Len() != 0 {
		t.Fatal("failed test")
	}
}
