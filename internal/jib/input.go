package jib

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os"
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
