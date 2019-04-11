package jib

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
	"unicode"
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

	r, err := git.PlainOpen(wd)
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

func CreateBranchFromIssue(instance Instance, issue Issue) string {
	wd, err := os.Getwd()
	if nil != err {
		fmt.Println("Something happened trying to detect working directory. Exiting...")
		os.Exit(1)
	}

	r, err := git.PlainOpen(wd)
	if nil != err {
		fmt.Println("Something happened trying to open git repo. Exiting...")
		os.Exit(1)
	}

	masterRef, err := r.Reference(plumbing.ReferenceName("refs/heads/" + instance.MainBranch), false)

	if nil != err {
		fmt.Println("Something happened trying to fetch main branch reference:" + err.Error())
		os.Exit(1)
	}

	issueType := strings.ToLower(issue.Fields.IssueType.Name)
	summary := strings.ToLower(removeNonAlphanumeric(strings.Replace(splitSummary(issue.Fields.Summary), " ", "-", -1)))

	branchName := "refs/heads/" + issueType + "/" + issue.Key + "-" + summary

	ref := plumbing.NewHashReference(plumbing.ReferenceName(branchName), masterRef.Hash())

	err = r.Storer.SetReference(ref)

	if nil != err {
		fmt.Println("Something happened trying to create branch. Exiting...")
		os.Exit(1)
	}

	return branchName
}

func CheckoutBranch(branchName string) {
	wd, err := os.Getwd()
	if nil != err {
		fmt.Println("Something happened trying to detect working directory. Exiting...")
		os.Exit(1)
	}

	r, err := git.PlainOpen(wd)
	if nil != err {
		fmt.Println("Something happened trying to open git repo. Exiting...")
		os.Exit(1)
	}

	wt, err := r.Worktree()

	if nil != err {
		fmt.Println("Something happened trying to fetch the git worktree: " + err.Error())
		os.Exit(1)
	}

	err = wt.Checkout(&git.CheckoutOptions{
		Branch: plumbing.ReferenceName(branchName),
	})

	if nil != err {
		fmt.Println("Something happened trying to checkout: " + err.Error())
		os.Exit(1)
	}
}

func removeNonAlphanumeric (string string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9\\s-_]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(string, "")
}

func splitSummary(summary string) string {
	splittingPos := 0
	for pos, char := range summary[0:25] {
		if unicode.IsSpace(char) {
			splittingPos = pos
		}
	}

	if 0 == splittingPos {
		return summary[0:25]
	}

	return summary[0:splittingPos]
}