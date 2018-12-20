package main

import "C"
import (
	"bufio"
	"flag"
	"fmt"
	"github.com/benesgarage/jib/internal/jib"
	"os"
)

func main() {
	i := flag.Bool("i", false, "Add a new JIRA instance")
	c := flag.Bool("c", false, "Show task comments")
	flag.Parse()

	if *i { jib.AddInstance() }

	core := jib.NewCore()
	writer := bufio.NewWriter(os.Stdout)

	jib.GetSummary(*core).OutputToTerminal(writer)

	switch true {
	case *c:
		jib.GetCommentSection(*core).OutputToTerminal(writer)
	}

	err := writer.Flush()

	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}
}
