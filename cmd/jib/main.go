package main

import "C"
import (
	"flag"
	"github.com/benesgarage/jib/internal/jib"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)+"/../.."
)

func main() {
	i := flag.Bool("i", false, "Add a new JIRA instance")
	flag.Parse()
	config, _ := jib.LoadConfig(basepath+"/config/config.json")
	switch true {
	case *i:
		jib.AddInstance(config)
		break
	default:
		jib.TaskSummary(config)
	}
}
