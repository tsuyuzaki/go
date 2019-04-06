package csv

import (
	"fmt"
	"os"
	"strings"
	"os/exec"
)

const bat = `@echo off
if "%~1"=="" (
	assoc .xls
) else if not "%~1"=="" (
	ftype "%~1"
)`
const batName = "find_csv_cmd.bat"

func getExcelPath() (string, bool) {
	if !createFile(batName, bat) {
		return "", false
	}
	defer func() {
		if err := os.Remove(batName); err != nil {
			fmt.Fprintf(os.Stderr, "os.Remove error [%v]\n", err)
		}
	}()
	
	out, err := exec.Command(batName).Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can not find Excel. err[%v]\n", err)
		return "", false
	}
	splitted := strings.Split(string(out), "=")
	if len(splitted) < 2 {
		fmt.Fprintf(os.Stderr, "Can not find Excel\n\t $ assoc .xls\n%s\n", out)
		return "", false
	}
	
	out, err = exec.Command(batName, splitted[1]).Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can not find Excel. err[%v]\n", err)
		return "", false
	}
	return toExcelPath(string(out))
}

func toExcelPath(s string) (string, bool) {
	index := strings.Index(s, `"`)
	if index == -1 {
		fmt.Fprintf(os.Stderr, "Can not find Excel. cmd out[%s]\n", s)
		return "", false
	}
	s = s[index+1:]
	index = strings.Index(s[1:], `"`)
	if index == -1 {
		fmt.Fprintf(os.Stderr, "Can not find Excel. cmd out[%s]\n", s)
		return "", false
	}
	return s[:index+1], true
}
