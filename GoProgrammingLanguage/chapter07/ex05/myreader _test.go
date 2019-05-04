package myreader

import (
	"io"
	"testing"
)

type testReader struct {
	s string
	cur int
}

func (r *testReader) Read(p []byte) (int, error) {
	n := copy(p, r.s[r.cur:])
	r.cur += n
	if r.cur < len(r.s) {
		return n, nil
	}
	return n, io.EOF
}

func NewTestReader(n int64) io.Reader {
	s := ""
	for i := int64(0); i < n; i++ {
		s += "a"
	}
	return &testReader{s: s, cur: 0}
}

func TestReader(t *testing.T) {
	const maxSize = 10

	tr := NewTestReader(maxSize)
	r := LimitReader(tr, maxSize)
	
	p := make([]byte, maxSize-1)
	n, err := r.Read(p)
	if err != nil || n != len(p) {
		t.Errorf("maxSize=%d, len(p)=%d, error=%v, n=%d\n", maxSize, len(p), err, n)
	}

	tr = NewTestReader(maxSize)
	r = LimitReader(tr, maxSize)

	p = make([]byte, maxSize)
	n, err = r.Read(p)
	if err != io.EOF || n != maxSize {
		t.Errorf("maxSize=%d, len(p)=%d, error=%v, n=%d\n", maxSize, len(p), err, n)
	}

	tr = NewTestReader(maxSize)
	r = LimitReader(tr, maxSize)

	p = make([]byte, maxSize+1)
	if err != io.EOF || n != maxSize {
		t.Errorf("maxSize=%d, len(p)=%d, error=%v, n=%d\n", maxSize, len(p), err, n)
	}

	tr = NewTestReader(maxSize)
	r = LimitReader(tr, 0)
	if err != io.EOF || n != 0 {
		t.Errorf("maxSize=%d, len(p)=%d, error=%v, n=%d\n", maxSize, len(p), err, n)
	}

	tr = NewTestReader(maxSize)
	r = LimitReader(tr, -1)
	if err != io.EOF || n != 0 {
		t.Errorf("maxSize=%d, len(p)=%d, error=%v, n=%d\n", maxSize, len(p), err, n)
	}
}
