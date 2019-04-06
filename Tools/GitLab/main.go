package main

import (
	"./csv"
	"./gitlab"
	"fmt"
	"os"
)

const kCSVFilePath = "input.csv"

const kFormData = "#ProjectURL,GroupURL1,GroupURL2,...\n#https://%[1]s/project/path,https://%[1]s/group/path1,https://%[1]s/group/path2,..."

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please input your token.")
		return
	}
	if !csv.CreateFileIfNotExists(kCSVFilePath, fmt.Sprintf(kFormData, "gitlab.com")) ||
		!csv.ShowCSV(kCSVFilePath) {
		return
	}

	rows, ok := csv.ReadCSV(kCSVFilePath)
	if !ok {
		return
	}
	for _, splittedRow := range rows {
		if len(splittedRow) <= 1 {
			fmt.Fprintf(os.Stderr, "Invalid row[%v].\n", splittedRow)
			continue
		}
		gitlab.AddGroups(os.Args[1], splittedRow[0], splittedRow[1:])
	}
}
