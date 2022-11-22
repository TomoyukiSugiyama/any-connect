package cmd

import (
	"bytes"
	"fmt"
	"os/exec"
)

func connect() {
	configIfNotExist()

	cmdstr := "/opt/cisco/anyconnect/bin/vpn -s < " + c.configPath

	// tio := &timeout.Timeout{
	// 	Cmd:       exec.Command("sh", "-c", cmd),
	// 	Duration:  10 * time.Second,
	// 	KillAfter: 5 * time.Second,
	// }

	// exitStatus, out, _, err := tio.Run()
	// if exitStatus.Signaled {
	// 	fmt.Println(string(out))
	// 	fmt.Println("ERROR : Timeout is occurred. Please reconfiguration your setting by using `any-connect config`.")
	// 	os.Exit(1)
	// }

	cmd := exec.Command("sh", "-c", cmdstr)

	// out, err := exec.Command("sh", "-c", cmdstr).Output()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	os.Exit(1)
	// }

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {
		fmt.Printf("Stdout: %s\n", stdout.String())
		fmt.Printf("Stderr: %s\n", stderr.String())
	} else {
		fmt.Printf("Stdout: %s\n", stdout.String())
	}
	// t := time.NewTimer(5 * time.Second)
	// <-t.C
	// syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
	// fmt.Println(cmd.Process.Pid)
	// fmt.Println("ERROR : Timeout is occurred. Please reconfiguration your setting by using `any-connect config`.")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	os.Exit(1)
	// }

	// // fmt.Println(out)
	//fmt.Println(string(out))

}
