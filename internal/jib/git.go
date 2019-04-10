package jib

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"os"
	"regexp"
	"time"
)

type IssueIdentifierNotFoundError struct {
	When time.Time
	What string
}

func (e IssueIdentifierNotFoundError) Error() string  {
	return e.What
}

type RemoteNotFoundError struct {
	When time.Time
	What string
}

func (e RemoteNotFoundError) Error() string  {
	return e.What
}

func GetBranchTaskNumber(branch string) (taskNumber string, err error){
	reg := regexp.MustCompile("[a-zA-Z]+[-]?[0-9]+")
	taskNumber = reg.FindString(branch)

	if taskNumber == "" {
		return taskNumber, IssueIdentifierNotFoundError{
			time.Now(),
			fmt.Sprintf("Could not find issue identifier in branch '%s'", branch),
		}
	}

	return taskNumber, nil
}

func GetBranch() string {
	wd, err := os.Getwd()
	if nil != err {
		fmt.Println("Something happened trying to detect working directory. Exiting...")
		os.Exit(1)
	}

	r, err := git.PlainOpen(wd+"/.git/")
	if nil != err {
		fmt.Println("Something happened trying to open git repo. Exiting...")
		os.Exit(1)
	}

	head, err :=r.Head()

	if nil != err {
		fmt.Println("Something happened trying to get head. Exiting...")
		os.Exit(1)
	}

	return string(head.Name())
}