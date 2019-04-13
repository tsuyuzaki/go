/**
 * 下記のシグニチャを持つ関数 CountingWriter を書きなさい。
 * io.Writer が与えられたなら、それを含む新たな Writer と int64 変数へのポインタを返します。
 * その変数は新たな Writer に書き込まれたバイト数を常に保持しています。
 *
 * func CountingWriter(w io.Writer) (io.Writer, *int64)
 */
package counter

import (
	"io"
)

type countingWriter struct {
	w io.Writer
	sum int64
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := &countingWriter{w, 0}
	return cw, &(cw.sum)
}

func (cw *countingWriter) Write(p []byte) (int, error) {
	n, err := cw.w.Write(p)
	cw.sum += int64(n)
	return n, err
}
