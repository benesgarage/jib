package main

import "C"
import (
	"context"
	"flag"
	"github.com/benesgarage/jib/internal/jib"
	"github.com/google/subcommands"
	"os"
)

func main() {
	subcommands.Register(jib.NewSetup(), "")
	//i := flag.Bool("i", false, "Add a new JIRA instance")
	//c := flag.Bool("c", false, "Show task comments")
	//b := flag.String("b", "", "Create branch from task number")
	flag.Parse()

	ctx := context.Background()

	os.Exit(int(subcommands.Execute(ctx)))
	//core := jib.NewCore()
	//writer := bufio.NewWriter(os.Stdout)
	//
	//
	//summary := jib.GetSummary(*core)
	//if *b != "" {
	//	jib.CreateBranchFromSummary(*core, summary)
	//}
	//summary.OutputToTerminal(writer)
	//
	//switch true {
	//case *c:
	//	jib.GetCommentSection(*core).OutputToTerminal(writer)
	//}
	//
	//err := writer.Flush()
	//
	//if nil != err {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
}
