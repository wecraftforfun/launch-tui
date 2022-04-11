package cmds

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/wecraftforfun/launch-tui/models"
)

// List returns the list of user created Agents from launchctl.
func List() tea.Msg {
	cmd := exec.Command("launchctl", "list")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	home, _ := os.UserHomeDir()
	files, _ := ioutil.ReadDir(path.Join(home, "Library", "LaunchAgents"))
	f := map[string]int{}
	for _, v := range files {
		f[strings.ReplaceAll(v.Name(), ".plist", "")]++
	}

	s := out.String()
	list := strings.Split(s, "\n")
	processes := []models.Process{}
	for i, v := range list {
		if i == 0 || i == len(list)-1 {
			continue
		}
		s := strings.Split(v, "\t")
		label := strings.Trim(s[2], " ")
		pid := strings.Trim(s[0], " ")
		status, _ := strconv.Atoi(strings.Trim(s[1], " "))
		if _, found := f[label]; found {
			processes = append(processes, models.Process{
				Pid:      pid,
				Status:   status,
				Label:    label,
				IsLoaded: true,
			})
			delete(f, label)
		}
	}
	for k := range f {
		processes = append(processes, models.Process{
			Pid:      "-",
			Status:   0,
			Label:    k,
			IsLoaded: false,
		})
	}
	return models.UpdateListMessage{
		List: processes,
	}
}

// Start use launchctl to start the specified Agent.
func Start(label string) tea.Cmd {
	return func() tea.Msg {
		cmd := exec.Command("launchctl", "start", label)
		err := cmd.Run()
		if err != nil {
			return models.ErrorMessage{
				Err: err,
			}
		}
		return models.CommandSuccessFullMessage{
			Cmd:   "start",
			Label: label,
		}
	}

}

// GetStatus update the status for the specified Agent
func GetStatus(label string) tea.Cmd {
	return func() tea.Msg {
		grep := exec.Command("grep", label)
		cmd := exec.Command("launchctl", "list")
		pipe, _ := cmd.StdoutPipe()
		defer pipe.Close()
		grep.Stdin = pipe
		cmd.Start()

		res, err := grep.Output()
		if err != nil {
			return models.ErrorMessage{
				Err: err,
			}
		}
		if result := string(res); result != "" {
			s := strings.Split(result, "\t")
			label := strings.Trim(s[2], " ")
			pid := strings.Trim(s[0], " ")
			status, _ := strconv.Atoi(strings.Trim(s[1], " "))
			return models.UpdateProcessStatusMessage{
				Process: models.Process{
					Label:    label,
					Pid:      pid,
					Status:   status,
					IsLoaded: true,
				},
			}
		} else {
			return models.UpdateProcessStatusMessage{
				Process: models.Process{
					Label:    label,
					Pid:      "-",
					Status:   0,
					IsLoaded: false,
				},
			}
		}

	}
}

func Load(label string) tea.Msg {
	home, _ := os.UserHomeDir()
	userId := strconv.Itoa(os.Getuid())
	return func() tea.Msg {
		cmd := exec.Command("launchctl", "bootstrap", "gui/"+userId, path.Join(home, "Library", "LaunchAgents", label+".plist"))
		err := cmd.Run()
		if err != nil {
			return models.ErrorMessage{
				Err: err,
			}
		}
		return models.CommandSuccessFullMessage{
			Cmd:   "load",
			Label: label,
		}
	}
}

func Unload(label string) tea.Msg {
	userId := strconv.Itoa(os.Getuid())
	return func() tea.Msg {
		cmd := exec.Command("launchctl", "bootout", "gui/"+userId+"/"+label)
		err := cmd.Run()
		if err != nil {
			return models.ErrorMessage{
				Err: err,
			}
		}
		return models.CommandSuccessFullMessage{
			Cmd:   "unload",
			Label: label,
		}
	}
}

// Stop use launchctl to stop the specified Agent.
func Stop(label string) tea.Cmd {
	return func() tea.Msg {
		cmd := exec.Command("launchctl", "stop", label)
		err := cmd.Run()
		if err != nil {
			return models.ErrorMessage{
				Err: err,
			}
		}
		return models.CommandSuccessFullMessage{
			Cmd:   "stop",
			Label: label,
		}
	}

}
