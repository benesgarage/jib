package jib

import (
	"context"
	"flag"
	"fmt"
	"github.com/benesgarage/jib/jira_client"
	"github.com/google/subcommands"
	"github.com/zalando/go-keyring"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing"
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

	switch CreateBranchFromIssueID(f.Arg(0), instance) {
	case 0:
		return subcommands.ExitSuccess
	default:
		return subcommands.ExitFailure
	}
}

func CreateBranchFromIssueID(issueID string, instance Instance) int {
	password, err := keyring.Get(instance.Host, instance.Username)

	if nil != err {
		fmt.Println(err)
		return 1
	}
	client := jira_client.BasicAuthClient{
		Host:instance.Host,
		Username:instance.Username,
		Password:password,
	}
	issue := client.ParseIssue(client.RequestIssue(issueID))

	repository, err := GetRepositoryFromWD()

	if nil != err {
		fmt.Println("Something happened trying to get the git repository: "+err.Error())
		return 1
	}


	branchName := GenerateBranchName(issue)

	worktree, err := repository.Worktree()

	if nil != err {
		fmt.Println("Something happened trying to get the git worktree: "+err.Error())
		return 1
	}

	mainRef, err := repository.Reference(plumbing.ReferenceName("refs/heads/" + instance.MainBranch), false)
	if nil != err {
		fmt.Println("Somethin happened trying to get the main branch ref: " + err.Error())
		return 1
	}

	err = worktree.Checkout(&git.CheckoutOptions{
		Hash: mainRef.Hash(),
		Create: false,
		Force: false,
	})

	if nil != err {
		fmt.Println("Something happened trying to checkout main branch: " + err.Error())
		return 1
	}

	// Create new branch
	err = repository.CreateBranch(&config.Branch{
		Name: branchName,
		Merge: plumbing.ReferenceName("refs/heads/" + branchName),
	})

	if nil != err {
		fmt.Println("Something happened trying to create the new branch: "+err.Error())
		return 1
	}

	err = worktree.Checkout(&git.CheckoutOptions{
		Hash: mainRef.Hash(),
		Branch: plumbing.ReferenceName("refs/heads/" + branchName),
		Create: true,
		Force: false,
	})

	if nil != err {
		fmt.Println("Something happened trying to get the main branch ref: "+err.Error())
		return 1
	}

	return 0
}
