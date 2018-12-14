package jib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
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
	MainBranch string
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

func (config Config) Persist(filename string) {
	configJson, _ := json.Marshal(config)
	_ = ioutil.WriteFile(filename, configJson, 0644)
}

func (config Config) GetInstance(origin string) (instance Instance, err error) {
	for _, instance := range config.Instances {
		if origin == instance.Origin {
			return instance, err
		}
	}

	err = UnconfiguredInstanceError{
		time.Now(),
		fmt.Sprintf("No JIRA instance configured for origin %s", origin),
	}

	return instance, err
}

func (config *Config) SetInstance (newInstance Instance) {
	for index, instance := range config.Instances {
		if newInstance.Origin == instance.Origin {
			config.Instances[index] = newInstance
			return
		}
	}
	config.Instances = append(config.Instances, newInstance)
}