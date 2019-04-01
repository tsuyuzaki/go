package counter

import (
	"fmt"
	"testing"
)

func TestCounter(t *testing.T) {
	tests := []struct {
		str string
		wc  WordCounter
		lc  LineCounter
	}{
		{"", 0, 0},
		{"あめんぼ 赤いな  あいうえお", 3, 1},
		{`あめんぼ　赤いな あいうえ
お`, 4, 2},
		{`あめんぼ
赤いな
あいうえお`, 3, 3},
	}

	for _, test := range tests {
		var wc WordCounter
		fmt.Fprintf(&wc, test.str)
		if wc != test.wc {
			t.Errorf("WordCounter error: res[%d] testdata[%s] expected[%v]", wc, test.str, test.wc)
		}

		var lc LineCounter
		fmt.Fprintf(&lc, test.str)
		if lc != test.lc {
			t.Errorf("LineCounter error: res[%d] testdata[%s] expected[%v]", lc, test.str, test.lc)
		}
	}
}
