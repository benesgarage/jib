package jira_client

import (
	"github.com/logrusorgru/aurora"
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