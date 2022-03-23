package cmds

import (
	"bytes"
	"log"
	"os/exec"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/wecraftforfun/launch-tui/models"
)

func List() tea.Msg {
	cmd := exec.Command("launchctl", "list")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
	s := out.String()
	list := strings.Split(s, "\n")
	processes := []models.Process{}
	for i, v := range list {
		if i == 0 || i == len(list)-1 {
			continue
		}
		s := strings.Split(v, "\t")
		pid := strings.Trim(s[0], " ")
		status, _ := strconv.Atoi(strings.Trim(s[1], " "))
		processes = append(processes, models.Process{
			Pid:    pid,
			Status: status,
			Label:  strings.Trim(s[2], " "),
		})
	}
	return models.UpdateListMessage{
		List: processes,
	}
}

func Start() {}

func Load() {}

func Unload() {}

func Stop() {}
