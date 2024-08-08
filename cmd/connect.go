package cmd

import (
	"bytes"
	"fmt"
	"os/exec"
	"syscall"
	"time"
)

func connect() {
	configIfNotExist()

	cmdstr := "/opt/cisco/anyconnect/bin/vpn -s < " + c.configPath
	cmd := exec.Command("sh", "-c", cmdstr)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}

	cmd.Start()

	go func() {
		t := time.NewTimer(120 * time.Second)
		<-t.C
		fmt.Println("process kill")
		syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
		fmt.Println("pid : ", cmd.Process.Pid)
		fmt.Println("ERROR : Timeout is occurred. Please reconfiguration your setting by using `any-connect config`.")
	}()

	cmd.Wait()

	fmt.Println(stdout.String())
}
