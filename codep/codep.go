package codep

import (
	"os/exec"

	"github.com/naoki85/entrykit"
)

func init() {
	entrykit.Cmds["codep"] = Run
}

func Run(config *entrykit.Config) {
	defer entrykit.Exec(config.Exec)
	err := Codep(config.Tasks)
	if err != nil {
		entrykit.Error(err)
	}
}

func Codep(tasks map[string]string) error {
	if len(tasks) == 0 {
		return nil
	}

	done := make(chan error)
	cmds := []*exec.Cmd{}

	for _, task := range tasks {
		cmd := entrykit.CommandTask(task)
		err := cmd.Start()
		if err != nil {
			return err
		}
		cmds = append(cmds, cmd)
		go func() {
			done <- cmd.Wait()
		}()
	}
	entrykit.ProxySignals(cmds)

	err := <-done
	if err != nil {
		return err
	}
	return nil
}
