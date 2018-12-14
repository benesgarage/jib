package jib

import (
	"fmt"
	"github.com/zalando/go-keyring"
	"os"
)

func AddInstance () () {

	origin := getOrigin()

	if checkOriginExists(origin) {
		fmt.Print("Instance for this repo is already linked! Overwrite configuration? ")
		if !askForConfirmation() {
			TaskSummary()
			os.Exit(0)
		}
	}

	host 	 := requestHost()
	port 	 := requestPort()
	username := requestUsername()
	password := requestPassword()
	mainBranch := requestMainBranch()

	config, err := LoadConfigs(basepath+"/config/config.json")

	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	config.SetInstance(Instance{
		Origin: origin,
		Host: host,
		Port: port,
		Username: username,
		MainBranch: mainBranch,
	})

	config.Persist(basepath+"/config/config.json")

	err = keyring.Set(host, username, string(password))
}
