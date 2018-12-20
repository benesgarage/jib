package jib

import (
	"fmt"
	"io"
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
	t := template.Must(template.New("comments").Funcs(funcMap).Parse(
		`
------------
| Comments |
------------
{{ range $comment := .Comments }}
{{ $comment.Author.DisplayName }}
{{ Repeat "-" (len $comment.Author.DisplayName) }}
{{ $comment.Body }}
{{ end }}
`))
	err := t.Execute(writer, commentSection)

	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}
}