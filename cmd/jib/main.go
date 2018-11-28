package main

import "C"
import (
	"flag"
	"github.com/benesgarage/jib/internal/jib"
)

func main() {
	i := flag.Bool("i", false, "Add a new JIRA instance")
	flag.Parse()

	switch true {
	case *i:
		jib.AddInstance()
		break
	default:
		jib.TaskSummary()
	}
}
