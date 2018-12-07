package jib

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func askForConfirmation() bool {
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}
	okayResponses := []string{"y", "Y", "yes", "Yes", "YES"}
	nokayResponses := []string{"n", "N", "no", "No", "NO"}
	if containsString(okayResponses, response) {
		return true
	} else if containsString(nokayResponses, response) {
		return false
	} else {
		fmt.Println("Please type yes or no and then press enter:")
		return askForConfirmation()
	}
}

func requestHost() (host string) {
	for {

		fmt.Print("Host: ")

		scanner := bufio.NewScanner(os.Stdin)

		scanner.Scan()
		host = scanner.Text()

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading from input: ", err)
		}

		if len(host) != 0 {
			return host
		}

		fmt.Println("Can not insert empty host parameter. Please insert a valid host.")
	}
}

func requestPort() (port int) {
	for {

		fmt.Print("Port[80]: ")

		scanner := bufio.NewScanner(os.Stdin)

		scanner.Scan()
		out := scanner.Text()

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading from input: ", err)
			os.Exit(1)
		}

		if len(out) != 0 {
			var err error

			if port, err = strconv.Atoi(out); err != nil {
				fmt.Println("Invalid port. Please input a valid integer.")
				continue
			}
		} else {
			port = 80
		}

		return port
	}
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

func requestMainBranch() (mainBranch string) {
	defaultBranch := "master"

	fmt.Print("Main branch["+defaultBranch+"]: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if nil != err {
		log.Print("  Scan for host failed, due to ", err)
		os.Exit(1)
	}
	mainBranch = strings.Trim(strings.Replace(input, "\n", "", -1), " ")

	if len(mainBranch) == 0 {
		mainBranch = defaultBranch
	}

	return mainBranch
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