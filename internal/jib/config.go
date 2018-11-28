package jib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"time"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)+"/../.."
)

type Config struct {
	Instances []Instance
}

type Instance struct {
	Origin string
	Host string
	Port int
	Username string
	Projects []Project
}

type Project struct {
	Prefix string
}

type UnconfiguredInstanceError struct {
	When time.Time
	What string
}

func (e UnconfiguredInstanceError) Error() string  {
	return e.What
}

func LoadConfigs(filename string) (Config, error) {
	var config Config
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if nil != err {
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)

	return config, err
}

func SaveConfig(filename string, config Config) {
	configJson, _ := json.Marshal(config)
	_ = ioutil.WriteFile(filename, configJson, 0644)
}

func GetInstance(TaskNumber string, config Config) (instance Instance, err error) {
	reg := regexp.MustCompile("[a-zA-Z]+")
	prefix := reg.FindString(TaskNumber)
	for _, instance := range config.Instances {
		for _, project := range instance.Projects {
			if prefix == project.Prefix {
				return instance, err
			}
		}
	}

	err = UnconfiguredInstanceError{
		time.Now(),
		fmt.Sprintf("No JIRA instance configured for task %s", TaskNumber),
	}

	return instance, err
}