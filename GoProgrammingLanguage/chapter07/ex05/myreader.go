/**
 * io パッケージの LimitReader 関数は io.Reader である r とバイト数 n を受け取り、
 * r から読み出す別の Reader を返しますが、n バイトを読みだした後にファイルの終わりの状態を報告します。
 * その関数を実装しなさい。
 *
 * func LimitReader(r io.Reader, n int64) io.Reader
 */
package myreader

import "io"

type myReader struct {
	r io.Reader // underlying reader
	n int64     // max bytes remaining
}

func (r *myReader) Read(p []byte) (n int, err error) {
	if r.n <= 0 {
		return 0, io.EOF
	}

	if int64(len(p)) > r.n {
		bs := make([]byte, r.n)
		n, err = r.r.Read(bs)
		copy(p, bs)
	} else {
		n, err = r.r.Read(p)
	}
	r.n -= int64(n)
	if err == nil && r.n <= 0{
		err = io.EOF
	}
	return n, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &myReader{r: r, n: n}
}