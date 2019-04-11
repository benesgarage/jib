package jib

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

type Issue struct {
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
	IssueType IssueType
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

type IssueType struct {
	Name string
}

type StatusCategory struct {
	ColorName string
}

func (summary Issue) OutputToTerminal (writer io.Writer) {
	funcMap := map[string]interface{}{
		"Repeat": func(s string, count int) string { return strings.Repeat(s, count) },
	}
	b, err := ioutil.ReadFile(basepath+"/internal/jib/issue.txt.tmpl")
	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}
	s := string(b)

	err = template.Must(template.New("summary").Funcs(funcMap).Parse(s)).Execute(writer, summary)
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