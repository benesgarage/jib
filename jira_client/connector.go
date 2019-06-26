package jira_client

import (
	"fmt"
	"net/http"
	"os"
)

func (client BasicAuthClient) RequestIssue (issueID string) *http.Response {
	req, err := http.NewRequest("GET", client.Host+"/rest/api/latest/issue/"+issueID, nil)

	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	if req == nil {
		os.Exit(1)
	}

	resp, err := client.DoAuthorized(req)

	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	return resp
}

func (client BasicAuthClient) RequestCommentSection (issueID string) *http.Response {
	req, err := http.NewRequest("GET", client.Host+"/rest/api/latest/issue/"+issueID, nil)

	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	if req == nil {
		os.Exit(1)
	}

	resp, err := client.DoAuthorized(req)

	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	return resp
}