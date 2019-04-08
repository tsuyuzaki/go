package main

import (
	"./gitlab"
	"./infile"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

const membersCSV = "users.csv"
const addMembersFormData = "#GroupURL,GroupAccess(10/20/30/40 or 50),UserName1,UserName2,...\n#https://%[1]s/group/path,30,norihiko-tsuyuzaki,kimio-hashimoto,..."

const groupsCSV = "groups.csv"
const addGroupsFormData = "#ProjectURL,GroupAccess(10/20/30/40 or 50),GroupURL1,GroupURL2,...\n#https://%[1]s/project/path,30,https://%[1]s/group/path1,https://%[1]s/group/path2,..."

func main() {
	if len(os.Args) != 3 {
		fmt.Printf(`Usage:
 $ %s [g or p] <Personal Access Token>`, filepath.Base(os.Args[0]))
		return
	}
	var csvFilePath, formData string
	if os.Args[1] == "g" {
		csvFilePath = membersCSV
		formData = addMembersFormData
	} else if os.Args[1] == "p" {
		csvFilePath = groupsCSV
		formData = addGroupsFormData
	} else {
		fmt.Println("First argument must be g or p")
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
	if os.Args[1] == "g" {
		addMembers(rows)
	} else if os.Args[1] == "p" {
		addGroups(rows)
	}
}

func addMembers(rows [][]string) {
	for _, splittedRow := range rows {
		if len(splittedRow) <= 2 {
			fmt.Fprintf(os.Stderr, "Invalid row[%v].\n", splittedRow)
			continue
		}
		gitlab.AddMembers(os.Args[2], splittedRow[0], splittedRow[1], splittedRow[2:])
	}
}

func addGroups(rows [][]string) {
	for _, splittedRow := range rows {
		if len(splittedRow) <= 2 {
			fmt.Fprintf(os.Stderr, "Invalid row[%v].\n", splittedRow)
			continue
		}
		gitlab.AddGroups(os.Args[2], splittedRow[0], splittedRow[1], splittedRow[2:])
	}
}

func confirm() bool {
	for {
		fmt.Printf("Do you want to continue a process? [Y/N]: ")
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
