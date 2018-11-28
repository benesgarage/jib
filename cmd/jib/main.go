package main

import "C"
import (
	"flag"
	"github.com/benesgarage/jib/internal"
)

func main() {
	i := flag.Bool("i", false, "Add a new JIRA instance")
	flag.Parse()

	switch true {
	case *i:
		internal.AddInstance()
		break
	default:
		internal.TaskSummary()
	}
}
