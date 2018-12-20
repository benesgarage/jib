package jib

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/zalando/go-keyring"
	"net/http"
	"os"
)

func GetSummary (core Core) (summary Summary) {

	issueURI := core.Instance.Host+"/rest/api/latest/issue/"+core.TaskNumber

	req, err   := http.NewRequest( "GET", issueURI, nil)

	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	response := doJiraRequest(core.Instance, req)
	jsonParser := json.NewDecoder(response.Body)
	decoderErr := jsonParser.Decode(&summary)

	if nil != decoderErr {
		fmt.Println(decoderErr)
		os.Exit(1)
	}

	return summary
}

func GetCommentSection (core Core) (commentSection CommentSection) {
	issueURI := core.Instance.Host+"/rest/api/latest/issue/"+core.TaskNumber+"/comment"

	req, err := http.NewRequest( "GET", issueURI, nil)

	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	response := doJiraRequest(core.Instance, req)
	jsonParser := json.NewDecoder(response.Body)
	decoderErr := jsonParser.Decode(&commentSection)

	if nil != decoderErr {
		fmt.Println(decoderErr)
		os.Exit(1)
	}

	return commentSection
}

func doJiraRequest (instance Instance, request *http.Request) (response *http.Response) {
	password, err := keyring.Get(instance.Host, instance.Username)

	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	request.Header.Add("Authorization", "Basic " + basicAuth(instance.Username, password))
	request.Header.Add("Content-Type", "application/json")

	client := http.Client{CheckRedirect:redirectPolicyFunc}

	response, httpErr := client.Do(request)

	if nil != httpErr {
		fmt.Println(httpErr)
		os.Exit(1)
	}

	return response
}

func redirectPolicyFunc(req *http.Request, via []*http.Request) error{
	lastRequest := via[len(via)-1]
	// Copy the headers from last request
	req.Header = lastRequest.Header
	// If domain has changed, remove the Authorization-header if it exists
	if req.URL.Host != lastRequest.URL.Host {
		req.Header.Del("Authorization")
	}

	return nil
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
