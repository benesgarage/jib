package jira_client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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