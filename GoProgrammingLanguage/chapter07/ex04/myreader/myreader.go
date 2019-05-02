/**
 * strings.NewReader 関数は、その引数である文字列から読み込むことで io.Reader インタフェース (とほかのインタフェース) を満足する値を返します。
 * 皆さん自身で簡単な NewReader を実装し、HTML パーサ (5.2節) が文字列からの入力を受け取るようにしなさい。
 */
package myreader

import "io"

type myReader struct {
	s string
	n int
}

func (r *myReader) Read(p []byte) (int, error) {
	copiedN := copy(p, r.s[r.n:])
	r.n += copiedN
	if r.n < len(r.s) {
		return copiedN, nil
	}
	// All read
	return copiedN, io.EOF
}

func NewReader(s string) io.Reader {
	return &myReader{s: s, n: 0}
}