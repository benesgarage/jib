package jira_client

import (
	"github.com/logrusorgru/aurora"
)

type UpdateIssue struct {
	Transition Transition
}

type IssueTransitions struct {
	Transitions []Transition
}

type Transition struct {
	ID string
	Name string
	To Status
}

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
	ID string
	StatusCategory StatusCategory
}

type IssueType struct {
	Name string
}

type StatusCategory struct {
	ColorName string
}

type Feed struct {
	Entry []Entry `xml:"entry"`
}

type Entry struct {
	ID string `xml:"id"`
	Title string `xml:"title"`
	Author Authore `xml:"author"`
	Updated string `xml:"updated"`
}

type Authore struct {
	Name string `xml:"name"`
	Email string `xml:"email"`
	URI string `xml:"uri"`
	Username string `xml:"username"`
}

// TODO: Move this? yes, move this
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