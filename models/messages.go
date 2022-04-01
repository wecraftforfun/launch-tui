package models

type UpdateListMessage struct {
	List []Process
}

type UpdateProcessStatusMessage struct {
	Process Process
}

type ErrorMessage struct {
	Err error
}

type CommandSuccessFullMessage struct {
	Cmd   string
	Label string
}
