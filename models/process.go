package models

import "fmt"

type Process struct {
	Pid      string
	Status   int
	Label    string
	IsLoaded bool
}

func (p Process) FilterValue() string {
	return p.Label
}
func (p Process) Title() string {
	return p.Label
}

func (p Process) Description() string {
	if !p.IsLoaded {
		return fmt.Sprintf("Agent is not currently loaded")
	}
	if p.Pid == "-" {
		return fmt.Sprintf("Agent is not currently running with Status %v.", p.Status)
	}
	return fmt.Sprintf("Agent is running with PID %v, and Status %v.", p.Pid, p.Status)
}
