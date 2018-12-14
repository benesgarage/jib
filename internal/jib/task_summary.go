package jib

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/zalando/go-keyring"
	"net/http"
	"os"
)

func TaskSummary() {
	instance   		:= determineInstance(getOrigin())
	summaryResponse := getTaskSummary(instance, extractTaskNumber())

	OutputTaskSummary(summaryResponse)
}


func extractTaskNumber() (taskNumber string) {
	branch := GetBranch()
	taskNumber = GetBranchTaskNumber(branch)

	return taskNumber
}

func determineInstance(origin string) Instance {
	config, err := LoadConfigs(basepath+"/config/config.json")
	instance, err := config.GetInstance(origin)

	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	return instance
}

func getTaskSummary(instance Instance, taskNumber string) (responseContent SummaryResponse) {

	issueUri := instance.Host+"/rest/api/latest/issue/"+taskNumber

	client := http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	req, _ := http.NewRequest(
		"GET",
		issueUri,
		nil,
	)

	password, err := keyring.Get(instance.Host, instance.Username)

	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	var basic = "Basic " + basicAuth(instance.Username, password)

	req.Header.Add("Authorization", basic)
	req.Header.Add("Content-Type", "application/json")

	resp, _ := client.Do(req)

	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	jsonParser := json.NewDecoder(resp.Body)

	err = jsonParser.Decode(&responseContent)

	return responseContent
}


func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
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
