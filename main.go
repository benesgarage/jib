package main

import "C"
import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/zalando/go-keyring"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func main() {
	i := flag.Bool("i", false, "Add a new JIRA instance")
	flag.Parse()

	switch true {
	case *i:
		addInstance()
		break
	default:
		taskSummary()
	}
}

func taskSummary() {
	taskNumber 		:= extractTaskNumber()
	instance   		:= determineInstance(taskNumber)
	summaryResponse := getTaskSummary(instance, taskNumber)

	outputTaskSummary(summaryResponse)
}


func extractTaskNumber() (taskNumber string) {
	branch := getBranch()
	taskNumber = getBranchTaskNumber(branch)

	return taskNumber
}

func determineInstance(taskNumber string) Instance {
	config, err := loadConfigs(basepath+"/config/config.json")
	instance, err := getInstance(taskNumber, config)

	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	return instance
}

func getTaskSummary(instance Instance, taskNumber string) (responseContent summaryResponse) {

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
