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
	parts := strings.Split(branch, "/")
	reg := regexp.MustCompile("[a-zA-Z]+[-]?[0-9]+")
	taskNumber = reg.FindString(parts[1])

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

func getOrigin() string {
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