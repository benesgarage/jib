package jira_client

import (
	"bytes"
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

func (client BasicAuthClient) RequestIssueTransitions (issueID string) *http.Response {
	req, err := http.NewRequest("GET", client.Host+"/rest/api/latest/issue/"+issueID+"/transitions", nil)

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

func (client BasicAuthClient) PostTransition (issueID string, transitionID string) *http.Response {
	requestByte := []byte(fmt.Sprintf(`{"transition":{"id":%s}}`, transitionID))

	req, err := http.NewRequest(
		"POST",
		client.Host+"/rest/api/latest/issue/"+issueID+"/transitions",
		bytes.NewReader(requestByte),
	)

	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	if req == nil {
		os.Exit(1)
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.DoAuthorized(req)

	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	return resp
}