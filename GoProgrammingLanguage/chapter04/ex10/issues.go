/**
 * 一か月未満、一年未満、一年以上の期間で分類された結果を報告するようにissuesを修正しなさい。
 */
package main

import (
	"./github"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	daySec   = 60 * 60 * 24
	monthSec = 30 * daySec
	yearSec  = daySec * 365
)

func print(description string, issues []*github.Issue) {
	fmt.Println(description)
	for _, issue := range issues {
		fmt.Printf("#%-5d %9.9s %.55s %v\n",
			issue.Number, issue.User.Login, issue.Title, issue.CreatedAt)
	}
}

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("$d issues:\n", result.TotalCount)

	var ItemsWithin1Month []*github.Issue
	var ItemsWithin1Year []*github.Issue
	var ItemsOverYear []*github.Issue

	now := time.Now().Unix()
	for _, item := range result.Items {
		diff := now - item.CreatedAt.Unix()
		if diff < monthSec {
			ItemsWithin1Month = append(ItemsWithin1Month, item)
		} else if diff < yearSec {
			ItemsWithin1Year = append(ItemsWithin1Year, item)
		} else {
			ItemsOverYear = append(ItemsOverYear, item)
		}
	}

	print("---- Within 1 month ----", ItemsWithin1Month)
	print("---- Within 1 year ----", ItemsWithin1Year)
	print("---- Over year ----", ItemsOverYear)
}
