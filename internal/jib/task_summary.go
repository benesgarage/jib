package jib

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"io"
	"os"
	"strings"
	"text/template"
)

type Summary struct {
	Expand string
	ID string
	Self string
	Key string
	Fields Fields
}

type Fields struct {
	Summary string
	Reporter Reporter
	Assignee Assignee
	Status Status
}

type Reporter struct {
	DisplayName string
	EmailAddress string
}

type Assignee struct {
	DisplayName string
	EmailAddress string
}

type Status struct {
	Name string
	StatusCategory StatusCategory
}

type StatusCategory struct {
	ColorName string
}

func (summary Summary) OutputToTerminal (writer io.Writer) {
	funcMap := map[string]interface{}{
		"Repeat": func(s string, count int) string { return strings.Repeat(s, count) },
	}
	err := template.Must(template.New("summary").Funcs(funcMap).Parse(
`
--{{ Repeat "-" (len .Fields.Summary) }}--
| {{ .Fields.Summary }} |
--{{ Repeat "-" (len .Fields.Summary) }}--
------------
Status: {{.Fields.Status.GetColorFormattedName}}
Reporter: {{.Fields.Reporter.DisplayName}} <{{.Fields.Reporter.EmailAddress}}>
Assignee: {{.Fields.Assignee.DisplayName}} <{{.Fields.Assignee.EmailAddress}}>
`)).Execute(writer, summary)

	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (status Status) GetColorFormattedName() (name string) {
	if "blue-gray" == status.StatusCategory.ColorName {
		name = aurora.Blue(status.Name).String()
	}
	if "yellow" == status.StatusCategory.ColorName {
		name = aurora.Brown(status.Name).String()
	}
	if "green" == status.StatusCategory.ColorName {
		name = aurora.Green(status.Name).String()
	}
	return name
}