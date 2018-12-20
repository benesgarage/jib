package jib

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

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

func GetOrigin() string {
	wd, err := os.Getwd()
	if nil != err {
		fmt.Println(err)
	}
	out, err := exec.Command("git", "-C", wd, "config", "--get", "remote.origin.url").Output()
	origin := strings.TrimSuffix(string(out), "\n")

	if nil != err {
		fmt.Println(err)
	}

	return string(origin)
}

func ExtractTaskNumber() (taskNumber string) {
	branch := GetBranch()
	taskNumber = GetBranchTaskNumber(branch)

	return taskNumber
}