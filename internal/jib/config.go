package jib

import (
	"bufio"
	"fmt"
	"github.com/pelletier/go-toml"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var (
	_, b, _, _  = runtime.Caller(0)
	basepath    = filepath.Dir(b)+"/../.."
	configRoute = basepath+"/config/"
	filename 	= "config.toml"
	jibConfig Config
)

type Config struct {
	Instances []Instance
}

type Instance struct {
	Location string
	Host string
	Port uint16
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

func initConfig() {
	file, err := os.Open(configRoute+filename)
	if nil != err {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	decoder := toml.NewDecoder(bufio.NewReader(file))

	if err := decoder.Decode(&jibConfig); nil != err {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func (config Config) toTOML() string {
	configToml, err := toml.Marshal(config)
	if nil != err {
		panic(fmt.Errorf("Fatal error marshalling configuration struct: %s \n", err))
	}
	return string(configToml)
}

func (config Config) persist() {
	configToml, err := toml.Marshal(config)
	if nil != err {
		panic(fmt.Errorf("Fatal error marshalling configuration struct: %s \n", err))
	}
	if err := ioutil.WriteFile(configRoute+filename, configToml, 0644); nil != err {
		panic(fmt.Errorf("Fatal error writing configuration file: %s \n", err))
	}
}

func (config Config) GetInstance(wd string) (instance Instance, err error) {
	for _, instance := range config.Instances {
		if wd == instance.Location {
			return instance, err
		}
	}

	err = UnconfiguredInstanceError{
		time.Now(),
		fmt.Sprintf("No instance configured for working directory %s", wd),
	}

	return instance, err
}

func (config *Config) setInstance (newInstance Instance) {
	for index, instance := range config.Instances {
		if newInstance.Location == instance.Location {
			config.Instances[index] = newInstance
			return
		}
	}
	config.Instances = append(config.Instances, newInstance)
}

func (config *Config) removeInstance (wd string) {
	for index, instance := range config.Instances {
		if wd == instance.Location {
			config.Instances = append(config.Instances[:index], config.Instances[index+1:]...)
			return
		}
	}
}

func (config Config) checkWorkingDirExists (workingDir string) bool {
	for _, instance := range config.Instances {
		if workingDir == instance.Location {

			return true
		}
	}

	return false
}