package infile

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func CreateFileIfNotExists(filePath, formData string) bool {
	if isExisting(filePath) {
		return true
	}
	return createFile(filePath, formData)
}

func ShowCSV(csvFilePath string) bool {
	editorPath, ok := getExcelPath()
	if !ok {
		fmt.Fprintf(os.Stderr, "Cannot find excel.exe\n")
		return false
	}
	cmd := exec.Command(editorPath, csvFilePath)
	err := cmd.Start()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return false
	}
	cmd.Wait()
	return true
}

func ReadCSV(csvFilePath string) ([][]string, bool) {
	f, err := os.Open(csvFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "os.Open() error [%v]\n", err)
		return nil, false
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	rows := make([][]string, 0)
	for s.Scan() {
		row := s.Text()
		if row[0] == '#' {
			continue
		}
		rows = append(rows, strings.Split(row, ","))
	}
	return rows, true
}

func isExisting(csvFilePath string) bool {
	_, err := os.Stat(csvFilePath)
	return err == nil
}

func createFile(filePath, formData string) bool {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_CREATE, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "os.OpenFile err[%v]\n", err)
		return false
	}
	defer f.Close()

	f.Write([]byte(formData))
	return true
}
