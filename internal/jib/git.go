package jib

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"time"
)

type RemoteNotFoundError struct {
	When time.Time
	What string
}

func (e RemoteNotFoundError) Error() string  {
	return e.What
}

func GetBranchTaskNumber(branch string) (taskNumber string){
	reg := regexp.MustCompile("[a-zA-Z]+[-]?[0-9]+")
	taskNumber = reg.FindString(branch)

	if taskNumber == "" {
		fmt.Println("Could not find a task number within the current branch name.")
		os.Exit(1)
	}

	return taskNumber
}

func GetBranch() (branch string){
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

func ExtractTaskNumber() (taskNumber string) {
	branch := GetBranch()
	taskNumber = GetBranchTaskNumber(branch)

	return taskNumber
}