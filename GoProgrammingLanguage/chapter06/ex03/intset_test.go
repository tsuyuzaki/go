/**
 * >go test -run Test6_2
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
			t.Errorf("failed test")
		}
		if !addTest(&s, i) {
			t.Errorf("failed test")
		}
	}

	if s.Len() != maxLen {
		t.Errorf("failed test")
	}
	if s.String() != s.Copy().String() {
		t.Errorf("failed test")
	}

	for i := 0; i < maxLen; i++ {
		if !removeTest(&s, i) {
			t.Errorf("failed test")
		}
		if !removeTest(&s, i) {
			t.Errorf("failed test")
		}
	}
	
	if s.Len() != 0 {
		t.Errorf("failed test")
	}
	if s.String() != s.Copy().String() {
		t.Errorf("failed test")
	}
	
	for i := 0; i < maxLen; i++ {
		if !addTest(&s, i) {
			t.Errorf("failed test")
		}
	}
	if s.Len() != maxLen {
		t.Errorf("failed test")
	}

	s.Clear()
	if s.Len() != 0 {
		t.Errorf("failed test")
	}
}

func Test6_2(t *testing.T) {
	s := IntSet{}
	added := []int{1, 100, 1000}
	s.AddAll(added...)
	if s.Len() != len(added) {
		t.Errorf("failed test")
	}

	for _, x := range added {
		if !s.Has(x) {
			t.Errorf("failed test")
		}
	}
}


func testIntersectWith(t *testing.T) {
	testdata := []struct{
		in1      []int
		in2      []int
		expected []int
	} {
		{
			[]int{0, 1, 10},
			[]int{0, 2, 10},
			[]int{0, 10},
		},
		{
			[]int{1, 2, 3, 65},
			[]int{1, 2, 3, 65},
			[]int{1, 2, 3, 65},
		},
		{
			[]int{1, 2, 3, 65},
			[]int{4, 5, 6, 66},
			[]int{},
		},
	}

	for _, td := range testdata {
		s := IntSet{}
		s.AddAll(td.in1...)

		passed := IntSet{}
		passed.AddAll(td.in2...)
		
		s.IntersectWith(&passed)
		for _, e := range td.expected {
			if !s.Has(e) {
				t.Errorf("failed test %v [%s]", td, s.String())
			}
		}
	}
}

func testDifferenceWith(t *testing.T) {
	testdata := []struct{
		in1      []int
		in2      []int
		expected []int
	} {
		{
			[]int{0, 1, 10},
			[]int{0, 2, 10},
			[]int{1},
		},
		{
			[]int{1, 2, 3, 65},
			[]int{1, 2, 3, 65},
			[]int{},
		},
		{
			[]int{1, 2, 3, 65},
			[]int{4, 5, 6, 66},
			[]int{1, 2, 3, 65},
		},
	}

	for _, td := range testdata {
		s := IntSet{}
		s.AddAll(td.in1...)

		passed := IntSet{}
		passed.AddAll(td.in2...)
		
		s.DifferenceWith(&passed)
		for _, e := range td.expected {
			if !s.Has(e) {
				t.Errorf("failed test %v [%s]", td, s.String())
			}
		}
	}
}

func testSymmetricDifference(t *testing.T) {
	testdata := []struct{
		in1      []int
		in2      []int
		expected []int
	} {
		{
			[]int{0, 1, 10},
			[]int{0, 2, 10},
			[]int{1, 2},
		},
		{
			[]int{1, 2, 3, 65},
			[]int{1, 2, 3, 65},
			[]int{},
		},
		{
			[]int{1, 2, 3, 65},
			[]int{4, 5, 6, 66},
			[]int{1, 2, 3, 4, 5, 6, 65, 66},
		},
	}

	for _, td := range testdata {
		s := IntSet{}
		s.AddAll(td.in1...)

		passed := IntSet{}
		passed.AddAll(td.in2...)
		
		s.SymmetricDifference(&passed)
		for _, e := range td.expected {
			if !s.Has(e) {
				t.Errorf("failed test %v [%s]", td, s.String())
			}
		}
	}
}

func Test6_3(t *testing.T) {
	testIntersectWith(t)
	testDifferenceWith(t)
	testSymmetricDifference(t)
}