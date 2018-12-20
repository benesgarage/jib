package jib

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _  = runtime.Caller(0)
	basepath    = filepath.Dir(b)+"/../.."
	configRoute = basepath+"/config/config.json"
)

type Core struct {
	Config Config
	Instance Instance
	TaskNumber string
}

func NewCore() *Core {
	config, _ := LoadConfig(configRoute)
	instance, err := config.GetInstance(GetOrigin())

	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	taskNumber := ExtractTaskNumber()

	return &Core{
		Config:config,
		Instance:instance,
		TaskNumber:taskNumber,
	}
}