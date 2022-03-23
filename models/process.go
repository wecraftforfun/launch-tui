package models

import "fmt"

type Process struct {
	Pid    string
	Status int
	Label  string
}

func (p Process) FilterValue() string {
	return p.Label
}
func (p Process) Title() string {
	return p.Label
}

func (p Process) Description() string {
	return fmt.Sprintf("Running with PID %v, and Status %v.", p.Pid, p.Status)
}
