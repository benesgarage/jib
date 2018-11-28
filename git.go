package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"regexp"
)

func getBranchTaskNumber(branch string) (taskNumber string){
	parts := strings.Split(branch, "/")
	reg := regexp.MustCompile("[a-zA-Z]+[-]?[0-9]+")
	taskNumber = reg.FindString(parts[1])

	return taskNumber
}

func getBranch() (branch string){
	dir, err := os.Getwd()
	if nil != err {
		return branch
	}
	if _, err := os.Stat(dir+"/.git/HEAD"); os.IsNotExist(err) {
		fmt.Println(dir+"/.git/HEAD does not exist!")
	}
	file, err := os.Open(dir+"/.git/HEAD")
	defer file.Close()
	if nil != err {
		return branch
	}
	var scanner = bufio.NewScanner(file)

	scanner.Scan()
	branch = scanner.Text()
	return branch[16:]
}