package jib

import (
	"fmt"
	"github.com/zalando/go-keyring"
	"os"
)

func AddInstance (config Config) () {

	origin := getOrigin()

	if config.checkOriginExists(origin) {
		fmt.Print("Instance for this repo is already linked! Overwrite configuration? ")
		if !askForConfirmation() {
			TaskSummary(config)
			os.Exit(0)
		}
	}

	host 	   := requestHost()
	port 	   := requestPort()
	username   := requestUsername()
	password   := requestPassword()
	mainBranch := requestMainBranch()

	config.SetInstance(Instance{
		Origin: origin,
		Host: host,
		Port: port,
		Username: username,
		MainBranch: mainBranch,
	})

	config.Persist(basepath+"/config/config.json")

	err := keyring.Set(host, username, string(password))

	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}
}
