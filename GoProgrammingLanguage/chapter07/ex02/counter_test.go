package counter

import (
	"os"
	"testing"
)

func TestCounter(t *testing.T) {
	cw, sum := CountingWriter(os.Stdout)
	expected := 0
	if *sum != int64(expected) {
		t.Errorf("CountingWriter error: *sum[%d] expected[%d]", *sum, expected)
	}

	in := []string{"あめんぼ", " 赤いな", " あいうえお"}
	for _, s := range in {
		n, err := cw.Write([]byte(s))
		if n != len(s) || err != nil {
			t.Errorf("cw.Write() returns n[%d (expected[%d])], err[%v]", n, len(s), err)
		}
		expected += n
		if *sum != int64(expected) {
			t.Errorf("CountingWriter error: *sum[%d] expected[%d]", *sum, expected)
		}
	}

}
