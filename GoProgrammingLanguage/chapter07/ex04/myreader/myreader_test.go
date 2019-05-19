package myreader

import (
	"io"
	"bytes"
	"testing"
)

func TestReader(t *testing.T) {
	s := "0123456789"
	{
		r := NewReader(s)
		p := make([]byte, len(s)-1)
		n, err := r.Read(p)
		if !bytes.Equal([]byte(s[:n]), p) || n != len(s)-1 || err != nil {
			t.Errorf("p=%s, n=%d, err=%v", string(p), n, err)
		}
		p = make([]byte, 1)
		n, err = r.Read(p)
		if !bytes.Equal([]byte(s[len(s)-n:]), p) || n != 1 || err != io.EOF {
			t.Errorf("p=%s, n=%d, err=%v", string(p), n, err)
		}
	}
	
	{
		r := NewReader(s)
		p := make([]byte, len(s))
		n, err := r.Read(p)
		if !bytes.Equal([]byte(s), p) || n != len(s) || err != io.EOF {
			t.Errorf("p=%s, n=%d, err=%v", string(p), n, err)
		}
	}
	
	{
		r := NewReader(s)
		p := make([]byte, len(s)+1)
		n, err := r.Read(p)
		if !bytes.Equal([]byte(s), p[:len(s)]) || n != len(s) || err != io.EOF {
			t.Errorf("p=%s, n=%d, err=%v", string(p), n, err)
		}
	}
	
}