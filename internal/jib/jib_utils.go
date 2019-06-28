package jib

import (
	"github.com/benesgarage/jib/jira_client"
	"log"
	"regexp"
	"strings"
	"unicode"
)

func GenerateBranchName(issue jira_client.Issue) string {
	issueType := strings.ToLower(issue.Fields.IssueType.Name)
	summary := strings.ToLower(removeNonAlphanumeric(strings.Replace(splitSummary(issue.Fields.Summary), " ", "-", -1)))

	return issueType + "/" + issue.Key + "-" + summary
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
