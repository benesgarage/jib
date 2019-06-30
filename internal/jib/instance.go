package jib

import (
	"fmt"
	"github.com/zalando/go-keyring"
	"os"
)

func addInstance (instance Instance) () {
	var err error

	instance.Location, err = os.Getwd()

	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	if jibConfig.checkWorkingDirExists(instance.Location) {
		fmt.Print("Instance for this repository is already configured! Overwrite configuration? \n")
		if !askForConfirmation() {
			os.Exit(0)
		}
	}

	jibConfig.setInstance(instance)

	password := requestPassword()

	jibConfig.persist()

	err = keyring.Set(instance.Host, instance.Username, string(password))

	if nil != err {
		fmt.Println("Something happened trying to save the password: " + err.Error())
		os.Exit(1)
	}
}

func removeInstance () {
	wd, err := os.Getwd()

	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	instance, err := jibConfig.GetInstance(wd)

	if nil != err {
		fmt.Printf("Error trying to find instance: %s \n", err)
		os.Exit(1)
	}

	err = keyring.Delete(instance.Host, instance.Username)

	if nil != err {
		fmt.Printf("Error trying to remove instance password from keychain: %s \n", err)
		fmt.Print("Disregard error and proceed with instance removal? \n")
		if !askForConfirmation() {
			os.Exit(1)
		}
	}

	jibConfig.removeInstance(instance.Location)
	jibConfig.persist()
}
