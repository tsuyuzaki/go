package main

import (
	"./gitlab"
	"./infile"
	"bufio"
	"fmt"
	"os"
)

const csvFilePath = "input.csv"
const formData = "#ProjectURL,GroupAccess(10/20/30/40 or 50),GroupURL1,GroupURL2,...\n#https://%[1]s/project/path,30,https://%[1]s/group/path1,https://%[1]s/group/path2,..."

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please input your token.")
		return
	}
	if !infile.CreateFileIfNotExists(csvFilePath, fmt.Sprintf(formData, "gitlab.com")) ||
		!infile.ShowCSV(csvFilePath) {
		return
	}

	if !confirm() {
		return
	}

	rows, ok := infile.ReadCSV(csvFilePath)
	if !ok {
		return
	}
	for _, splittedRow := range rows {
		if len(splittedRow) <= 2 {
			fmt.Fprintf(os.Stderr, "Invalid row[%v].\n", splittedRow)
			continue
		}
		gitlab.AddGroups(os.Args[1], splittedRow[0], splittedRow[1], splittedRow[2:])
	}
}

func confirm() bool {
	for {
		fmt.Printf("Do you want to continue to add access groups? [Y/N]: ")
		s := bufio.NewScanner(os.Stdin)
		if ok := s.Scan(); !ok {
			fmt.Fprintf(os.Stderr, "Scan error\n")
			return false
		}
		answer := s.Text()
		if answer == "Y" {
			return true
		} else if answer == "N" {
			return false
		}
	}
}
