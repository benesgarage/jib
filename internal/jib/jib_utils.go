package jib

import (
	"fmt"
	"github.com/benesgarage/jib/jira_client"
	"log"
	"regexp"
	"strings"
	"time"
	"unicode"
)

func GenerateBranchName(issue jira_client.Issue) string {
	issueType := strings.ToLower(issue.Fields.IssueType.Name)
	summary := strings.ToLower(removeNonAlphanumeric(strings.Replace(splitSummary(issue.Fields.Summary), " ", "-", -1)))

	return issueType + "/" + issue.Key + "-" + summary
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

func removeNonAlphanumeric (string string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9\\s-_]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(string, "")
}

func splitSummary(summary string) string {
	splittingPos := 0
	var end = 25
	if len(summary) < end {
		end = len(summary)
	}
	for pos, char := range summary[0:end] {
		if unicode.IsSpace(char) {
			splittingPos = pos
		}
	}

	if 0 == splittingPos {
		return summary[0:end]
	}

	return summary[0:splittingPos]
}
