package jib

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/subcommands"
	"os"
)

type branch struct {
	Instance Instance
	List bool
	Remove bool
	RemoteName string
}

func NewBranch() *branch {
	return &branch{}
}

func (*branch) Name() string {
	return "branch"
}

func (*branch) Synopsis() string {
	return "Create a git branch from a JIRA issue."
}

func (*branch) Usage() string {
	return `jib branch <issue-number>
`
}

func (setup *branch) SetFlags(f *flag.FlagSet) {
}

func (setup *branch) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	wd, err := os.Getwd()

	if nil != err {
		fmt.Println("Something happened trying to detect working directory. Exiting...")
	}

	instance, err := jibConfig.GetInstance(wd)

	if nil != err {
		if _, ok := err.(*UnconfiguredInstanceError); false == ok {
			fmt.Println("There is no instance configured for this repository. Please configure an instance with `jib setup`.")
			os.Exit(1)
		} else {
			fmt.Println("Something happened trying to fetch the configured instance for this repository. Exiting...")
			os.Exit(1)
		}
	}

	issueIdentifier := f.Arg(0)
	issue := GetIssue(instance, issueIdentifier)
	branchName := CreateBranchFromIssue(instance, issue)
	CheckoutBranch(branchName)

	return subcommands.ExitSuccess
}
