/**
 * (*IntSet).UnionWith はワード単位のビット和算子である | を使って二つのビットの和集合を計算しています。
 * セット操作に対応するメソッド IntersectWith, DifferenceWith, SymmetricDifference を実装しなさい。
 * (二つの集合の対称差 (symmetric difference) は、どちらかの集合にはあるが、両方にはない要素を含む集合です。)
 */
package intset

import (
	"bytes"
	"fmt"
	"../popcount"
)


type IntSet struct {
	words []uint64
}

func (s *IntSet) Len() int {
    len := 0
	for _, word := range s.words {
		len += popcount.PopCount(word)
	}
	return len
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	s.add(x)
}

func (s *IntSet) AddAll(xs ...int) {
	for _, x := range xs {
		s.add(x)
	}
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word >= len(s.words) {
		return // Nothing to do
	}
	s.words[word] &= ^(1 << bit)
}

func (s *IntSet) Clear() {
	s.words = []uint64{}
}

func (s *IntSet) Copy() *IntSet {
	var copied IntSet
	copied.words = append(copied.words, s.words...)
	return &copied
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			return
		}
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		s.words[i] &^= tword
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		s.words[i] ^= tword
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}