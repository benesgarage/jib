package internal

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
)

func OutputTaskSummary(summaryResponse SummaryResponse)  {
	fmt.Println(summaryResponse.Fields.Summary)
	fmt.Println(strings.Repeat("=", len(summaryResponse.Fields.Summary)))
	var col = color.New()

	if "blue-gray" == summaryResponse.Fields.Status.StatusCategory.ColorName {
		col = color.New(color.FgBlue)
	}
	if "yellow" == summaryResponse.Fields.Status.StatusCategory.ColorName {
		col = color.New(color.FgYellow)
	}
	if "green" == summaryResponse.Fields.Status.StatusCategory.ColorName {
		col = color.New(color.FgGreen)
	}
	fmt.Print("Status: ")
	_, err := col.Println(summaryResponse.Fields.Status.Name)

	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Reporter: "+summaryResponse.Fields.Reporter.DisplayName+" <"+summaryResponse.Fields.Reporter.EmailAddress+">")
	fmt.Println("Assignee: "+summaryResponse.Fields.Assignee.DisplayName+" <"+summaryResponse.Fields.Assignee.EmailAddress+">")
}