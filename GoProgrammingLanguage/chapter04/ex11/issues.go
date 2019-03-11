/**
 * コマンドラインからユーザがGitHubのイシューを作成、読み出し、更新、クローズできるツールを構築しなさい。
 * 大量のテキストを入力する必要がある場合には、ユーザの好みのテキストエディタを起動するようにしなさい。
 */
package main

import "./github"

func main() {
    github.PostNewIssue()
}
