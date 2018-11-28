package internal

type SummaryResponse struct {
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