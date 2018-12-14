package jib

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/zalando/go-keyring"
	"net/http"
	"os"
)

type JiraClient struct {
	Instance Instance
	httpClient http.Client
}

func (client JiraClient) GetTaskSummary (taskNumber string) (responseContent SummaryResponse) {
	issueURI := client.Instance.Host+"/rest/api/latest/issue/"+taskNumber

	req, _ := http.NewRequest( "GET", issueURI, nil)

	password, err := keyring.Get(client.Instance.Host, client.Instance.Username)

	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	req.Header.Add("Authorization", "Basic " + basicAuth(client.Instance.Username, password))
	req.Header.Add("Content-Type", "application/json")

	resp, _ := client.httpClient.Do(req)

	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	jsonParser := json.NewDecoder(resp.Body)

	err = jsonParser.Decode(&responseContent)

	return responseContent
}

func NewJiraClient(instance Instance) *JiraClient {
	return &JiraClient{instance, http.Client{CheckRedirect: redirectPolicyFunc}}
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
