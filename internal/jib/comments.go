package jib

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

type CommentSection struct {
	Total int
	Comments []Comment
}

type Comment struct {
	ID string
	Author Author
	Body string
}

type Author struct {
	Name string
	Key string
	EmailAddress string
	DisplayName string
	Active bool
}

func (commentSection CommentSection) OutputToTerminal (writer io.Writer) {
	funcMap := map[string]interface{}{
		"Repeat": func(s string, count int) string { return strings.Repeat(s, count) },
	}
	b, err := ioutil.ReadFile(basepath+"/internal/jib/comments.txt.tmpl")
	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}
	s := string(b)

	err = template.Must(template.New("comments").Funcs(funcMap).Parse(s)).Execute(writer, commentSection)
	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}
}