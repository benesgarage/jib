package jib

import (
	"context"
	"flag"
	"fmt"
	"github.com/benesgarage/jib/jira_client"
	"github.com/google/subcommands"
	strip "github.com/grokify/html-strip-tags-go"
	"github.com/zalando/go-keyring"
	"os"
	"regexp"
)

type activity struct {
	Username string
}

func NewActivity() *activity {
	return &activity{}
}

func (activity *activity) Name() string {
	return "activity"
}

func (activity *activity) Synopsis() string {
	return "Shows user activity for the given date."
}

func (activity *activity) Usage() string {
	return `jib activity <date> [-username <username>]
`
}

func (activity *activity) SetFlags (f *flag.FlagSet) {
	f.StringVar(&activity.Username, "username", "", "Provide a specific username to query.")

}

func (activity *activity) Execute (_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
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

	password, err := keyring.Get(instance.Host, instance.Username)

	client := jira_client.BasicAuthClient{
		Host:instance.Host,
		Username:instance.Username,
		Password:password,
	}

	username := instance.Username

	if "" != activity.Username {
		username = activity.Username
	}


	feed := client.ParseActivity(client.GetActivity(username, "1586995200000", "1587081599999"))

	for _, entry := range feed.Entry {
		space := regexp.MustCompile(`\s+`)
		fmt.Println("[" + entry.Updated + "] " + space.ReplaceAllString(strip.StripTags(entry.Title), " "))
	}

	return subcommands.ExitSuccess
}

