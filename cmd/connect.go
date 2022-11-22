package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/Songmu/timeout"
)

func showPS() {
	b, err := exec.Command("ps", "j").Output()
	fmt.Println(string(b), err)
}

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

	// var stdout bytes.Buffer
	// var stderr bytes.Buffer
	// cmd.Stdout = &stdout
	// cmd.Stderr = &stderr

	stdout, _ := cmd.StdoutPipe()
	// stdout, err := cmd.StdoutPipe()
	// if err != nil {
	// 	fmt.Println("cmd.StdoutPipe() error: " + err.Error())
	// }
	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
		}
	}()

	tio := &timeout.Timeout{
		Cmd:       cmd,
		Duration:  10 * time.Second,
		KillAfter: 5 * time.Second,
	}
	exitStatus, out, _, err := tio.Run()
	if exitStatus.Signaled {
		fmt.Println(string(out))
		fmt.Println("ERROR : Timeout is occurred. Please reconfiguration your setting by using `any-connect config`.")
		os.Exit(1)
	}
	if err != nil {
		fmt.Println(string(out))
		fmt.Println(err.Error())
		os.Exit(1)
	}
	// err := cmd.Run()
	// get pipe to standard output
	//stdout, err := cmd.StdoutPipe()
	// if err != nil {
	// 	fmt.Println("cmd.StdoutPipe() error: " + err.Error())
	// }

	// cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	// cmd.Start()
	// fmt.Println("process start", cmd.Process.Pid)
	// time.Sleep(time.Second)
	// showPS()
	// if err != nil {
	// 	fmt.Printf("Stdout: %s\n", stdout.String())
	// 	fmt.Printf("Stderr: %s\n", stderr.String())
	// } else {
	// 	fmt.Printf("Stdout: %s\n", stdout.String())
	// }
	// t := time.NewTimer(5 * time.Second)
	// <-t.C
	// syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
	// fmt.Println(cmd.Process.Pid)
	// fmt.Println("ERROR : Timeout is occurred. Please reconfiguration your setting by using `any-connect config`.")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	os.Exit(1)
	// }

	// go func() {
	// 	err := cmd.Wait()
	// 	fmt.Println("process done:", err)
	// }()

	// t := time.NewTimer(5 * time.Second)
	// <-t.C
	// fmt.Println("process kill")

	// syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
	// time.Sleep(time.Second)
	//fmt.Println(string(stdout.Bytes()))
	showPS()
	fmt.Println(out)
	//fmt.Println(string(out))
}
