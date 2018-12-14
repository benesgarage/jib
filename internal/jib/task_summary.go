package jib

func TaskSummary(config Config) {
	instance, _		:= config.GetInstance(getOrigin())
	jiraClient      := NewJiraClient(instance)
	summaryResponse := jiraClient.GetTaskSummary(extractTaskNumber())

	OutputTaskSummary(summaryResponse)
}


func extractTaskNumber() (taskNumber string) {
	branch := GetBranch()
	taskNumber = GetBranchTaskNumber(branch)

	return taskNumber
}
