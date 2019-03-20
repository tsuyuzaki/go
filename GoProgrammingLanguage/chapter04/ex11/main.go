/**
 * コマンドラインからユーザがGitHubのイシューを作成、読み出し、更新、クローズできるツールを構築しなさい。
 * 大量のテキストを入力する必要がある場合には、ユーザの好みのテキストエディタを起動するようにしなさい。
 */
package main

import (
	"./github"
	"fmt"
	"os"
)

const msg = `Please specify operation type.
    c: Create new issue
    s: Show and Update current issue

operation type: `

func main() {
	txt, ok := github.ScanText(msg)
	if !ok {
		os.Exit(1)
	}
	if txt == "c" {
		github.PostNewIssue()
	} else if txt == "s" {
		github.ShowIssue()
	} else {
		fmt.Println("Invalid operation type [", txt, "]")
	}
}
