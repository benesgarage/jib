package jib

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/subcommands"
	"os"
)

type setup struct {
	Instance Instance
	List bool
	Remove bool
	RemoteName string
}

func NewSetup() *setup {
	return &setup{}
}

func (*setup) Name() string {
	return "setup"
}

func (*setup) Synopsis() string {
	return "Configures a Jira/git integration."
}

func (*setup) Usage() string {
	return `jib setup [-port <port>] [-username <username>] [-main-branch <main-branch>] <host>:
	Configures a Jira/git integration. If a username is provided, the user will be prompted to
	enter a password.
`
}

func (setup *setup) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&setup.List, "list", false, "List all configured instances.")
	f.BoolVar(&setup.Remove, "remove", false, "Remove a configured instance based on the host provided.")

	f.StringVar(&setup.Instance.Host, "host", "", "JIRA host.")

	port := *f.Uint("port", 80, "JIRA host port.")
	if 65535 < port || port < 0 {
		fmt.Println("Port out of range [0-65535].")
		os.Exit(1)
	}
	setup.Instance.Port = uint16(port)

	f.StringVar(&setup.Instance.Username, "username", "", "JIRA username.")
	f.StringVar(&setup.RemoteName, "remote-name", "origin", "The name given to the remote repository within the local git configuration.")
	f.StringVar(&setup.Instance.MainBranch, "main-branch", "master", "Main branch for git repository.")
}

func (setup *setup) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if true == setup.List {
		setup.listInstances()
		return subcommands.ExitSuccess
	} else if true == setup.Remove {
		setup.removeInstance()
		return subcommands.ExitSuccess
	}
	setup.addInstance()
	return subcommands.ExitSuccess
}

func (setup *setup) listInstances ()  {
	fmt.Print(jibConfig.toTOML())
}

func (setup *setup) removeInstance () {
	if 0 == len(setup.RemoteName) {
		fmt.Println("No remote was provided. Please specify remote with -remote-name.")
		os.Exit(1)
	}

	removeInstance()
}

func (setup *setup) addInstance () {
	if 0 == len(setup.Instance.Host) {
		fmt.Println("No JIRA host was provided. Please specify host with -host.")
		os.Exit(1)
	}
	if 0 == len(setup.Instance.Username) {
		fmt.Println("No username was provided for JIRA. Please specify username with -username.")
		os.Exit(1)
	}
	if 0 == len(setup.RemoteName) {
		fmt.Println("No remote was provided. Please specify remote with -remote-name.")
		os.Exit(1)
	}

	addInstance(setup.Instance)
}