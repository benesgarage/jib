package jira_client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func (client BasicAuthClient) ParseIssue (response *http.Response) (issue Issue) {
	decoder := json.NewDecoder(response.Body)
	err := decoder.Decode(&issue)

	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	return issue
}

func (client BasicAuthClient) ParseCommentSection (response *http.Response) (commentSection CommentSection) {
	decoder := json.NewDecoder(response.Body)
	err := decoder.Decode(&commentSection)

	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	return commentSection
}

func (client BasicAuthClient) ParseIssueTransitions (response *http.Response) (issueTransitions IssueTransitions) {
	decoder := json.NewDecoder(response.Body)
	err := decoder.Decode(&issueTransitions)

	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	return issueTransitions
}

func (client BasicAuthClient) FindTransitionByName (response *http.Response, transitionName string) (transition Transition, found bool) {
	decoder := json.NewDecoder(response.Body)
	issueTransitions := new(IssueTransitions)
	err := decoder.Decode(&issueTransitions)

	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, transition := range issueTransitions.Transitions {
		if strings.ToLower(transitionName) == strings.ToLower(transition.Name) {
			return transition, true
		}
	}

	return transition, false
}