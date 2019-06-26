package jib

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"github.com/benesgarage/jib/jira_client"
	"github.com/google/subcommands"
	"github.com/zalando/go-keyring"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

type status struct {
	IssueIdentifier string
	Comments bool
}

func NewStatus() *status {
	return &status{}
}

func (*status) Name() string {
	return "status"
}

func (*status) Synopsis() string {
	return "Fetches and shows the status of a JIRA ticket linked to the current branch."
}

func (*status) Usage() string {
	return `jib status [-issue <issue-number>] [-comments]
Defaults to finding the issue number within the current branch. If -issue is specified, will query
the JIRA host for the specific issue identifier.
`
}

func (status *status) SetFlags (f *flag.FlagSet) {
	defaultIssueIdentifier, err := GetBranchTaskNumber(GetBranch())

	if _, ok := err.(*IssueIdentifierNotFoundError); true == ok {
		defaultIssueIdentifier = ""
	}

	f.StringVar(&status.IssueIdentifier, "issue", defaultIssueIdentifier, "Provide a specific issue identifier to query.")
	f.BoolVar(&status.Comments, "comments", false, "Show issue comments.")
}

func (status *status) Execute (_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	wd, err := os.Getwd()

	if nil != err {
		fmt.Println("Something happened trying to detect working directory. Exiting...")
		os.Exit(1)
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

	if "" == status.IssueIdentifier {
		fmt.Println("No issue identifier was provided. Please specify an issue identifier with the -issue flag.")
		os.Exit(1)
	}

	password, err := keyring.Get(instance.Host, instance.Username)

	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	client := jira_client.BasicAuthClient{
		Host:instance.Host,
		Username:instance.Username,
		Password:password,
	}

	funcMap := map[string]interface{}{
		"Repeat": func(s string, count int) string { return strings.Repeat(s, count) },
	}

	issue := client.ParseIssue(client.RequestIssue(status.IssueIdentifier))
	writer := bufio.NewWriter(os.Stdout)
	b, err := ioutil.ReadFile(basepath+"/internal/jib/issue.txt.tmpl")
	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}
	s := string(b)

	err = template.Must(template.New("summary").Funcs(funcMap).Parse(s)).Execute(writer, issue)
	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	if true == status.Comments {
		commentSection := client.ParseCommentSection(client.RequestCommentSection(status.IssueIdentifier))
		b, err := ioutil.ReadFile(basepath+"/internal/jib/comments.txt.tmpl")
		if nil != err {
			fmt.Println(err)
			os.Exit(1)
		}
		s := string(b)

		err = template.Must(template.New("comments").Funcs(funcMap).Parse(s)).Execute(writer, commentSection)
		if nil != err {
			fmt.Println(err)
			os.Exit(1)
		}	}

	err = writer.Flush()

	if nil != err {
		fmt.Printf("Something happened trying to show ticket status: %s", err)
		os.Exit(1)
	}

	return subcommands.ExitSuccess
}