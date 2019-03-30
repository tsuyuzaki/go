/**
 * IntSet で使われている個々のワードの型は uint64 ですが、64ビット演算は 32ビットプラットフォーム上では非効率かもしれません。
 * プラットフォームに対して最も効率的な符号なし変数である uint 型を使うようにプログラムを修正しなさい。
 * 64で割る代わりに、unit の実質的サイズのビット数である 32 あるいは 64 を保持する定数を定義しなさい。
 * そのためには、おそらくかなり賢い式 32 << (^uint(0) >> 63) を使えます。
 */
package intset

import (
	"bytes"
	"fmt"
	"../popcount"
)

const bitSize = 32 << (^uint(0) >> 63)

type IntSet struct {
	words []uint
}

func (s *IntSet) Len() int {
    len := 0
	for _, word := range s.words {
		len += popcount.PopCount(uint64(word))
	}
	return len
}

func (s *IntSet) Elems() []int {
	var ret []int
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for bit := 0; bit < bitSize; bit++ {
			if word&(1<<uint(bit)) != 0 {
				x := i*bitSize + bit
				ret = append(ret, x)
			}
		}
	}
	return ret
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/bitSize, uint(x%bitSize)
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
	word, bit := x/bitSize, uint(x%bitSize)
	if word >= len(s.words) {
		return // Nothing to do
	}
	s.words[word] &= ^(1 << bit)
}

func (s *IntSet) Clear() {
	s.words = []uint{}
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
		for j := 0; j < bitSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", bitSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) add(x int) {
	word, bit := x/bitSize, uint(x%bitSize)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}