package process

import (
	"fmt"
	"os/exec"
)

func RunDetached(cmds string) error {
	cmd := exec.Command(cmds)
	err := cmd.Start()
	if err != nil {
		return err
	}
	pid := cmd.Process.Pid
	// use goroutine waiting, manage process
	// this is important, otherwise the process becomes in S mode
	go func() {
		err = cmd.Wait()
		fmt.Printf("Command finished with error: %v", err)
	}()

	fmt.Printf("Command started with pid: %d\n", pid)

	return nil
}
