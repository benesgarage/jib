package internal

import (
	"fmt"
	"github.com/zalando/go-keyring"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os"
	"os/exec"
	"strings"
)

func AddInstance () () {

	origin := getOrigin()

	if checkOriginExists(origin) {
		//TODO: Ask if wants to edit
		fmt.Println("Instance for this repo is already linked! Sorry, edits are not implemented yet :(")
		os.Exit(1)
	}

	host 	 := requestHost()
	port 	 := requestPort()
	username := requestUsername()
	password := requestPassword()
	project  := requestProject()

	config, err := LoadConfigs(basepath+"/config/config.json")

	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	config.Instances = append(config.Instances, Instance{
		Origin: origin,
		Host: host,
		Port: port,
		Username: username,
		Projects: []Project{
			{
				Prefix:project,
			},
		},
	})

	SaveConfig(basepath+"/config/config.json", config)

	err = keyring.Set(host, username, string(password))
}

func requestHost() (host string) {
	fmt.Print("Host: ")
	if    _, err := fmt.Scan(&host);    err != nil {
		log.Print("  Scan for host failed, due to ", err)
		return
	}

	return host
}

func requestPort() (port int) {
	fmt.Print("Port [80]: ")
	if    _, err := fmt.Scan(&port);    err != nil {
		log.Print("  Scan for port failed, due to ", err)
		return
	}

	return port
}

func requestUsername() (username string) {
	fmt.Print("Username: ")
	if    _, err := fmt.Scan(&username);    err != nil {
		log.Print("  Scan for host failed, due to ", err)
		return
	}

	return username
}

func requestPassword() string {
	fmt.Print("Enter password: ")
	password, err := terminal.ReadPassword(0)

	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println()

	return string(password)
}

func requestProject() (project string) {
	fmt.Print("Enter project issue prefix (Example: INV, OIO, TACO): ")
	if    _, err := fmt.Scan(&project);    err != nil {
		log.Print("  Scan for host failed, due to ", err)
		return
	}

	return project
}

func getOrigin() string {
	wd, err := os.Getwd()
	if nil != err {
		fmt.Println(err)
	}
	out, err := exec.Command("git", "-C", wd, "config", "--get", "remote.origin.url").Output()
	origin := strings.TrimSuffix(string(out), "\n")

	if nil != err {
		fmt.Println(err)
	}

	return string(origin)
}

func checkOriginExists(repository string) bool {
	config, err := LoadConfigs(basepath+"/config/config.json")

	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, instance := range config.Instances {
		if repository == instance.Origin {

			return true
		}
	}

	return false
}