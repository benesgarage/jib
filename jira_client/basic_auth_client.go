package jira_client

import (
	"net/http"
)

type BasicAuthClient struct {
	http.Client
	Host string
	Username string
	Password string
}

func (client BasicAuthClient) DoAuthorized (req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(client.Username, client.Password)
	client.CheckRedirect = client.redirectFuncPolicy

	return client.Do(req)
}

func (client BasicAuthClient) redirectFuncPolicy (req *http.Request, via []*http.Request) error {
	lastRequest := via[len(via)-1]
	// Copy the headers from last request
	req.Header = lastRequest.Header
	// If domain has changed, remove the Authorization-header if it exists
	if req.URL.Host != lastRequest.URL.Host {
		req.Header.Del("Authorization")
	}

	return nil
}