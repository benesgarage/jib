package jib

import (
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _  = runtime.Caller(0)
	basepath    = filepath.Dir(b)+"/../.."
	configRoute = basepath+"/config/"
	filename 	= "config.toml"
	jibConfig Config
)

func init() {
	initConfig()
}