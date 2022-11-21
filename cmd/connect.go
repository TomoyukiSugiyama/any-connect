package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/Songmu/timeout"
)

func connect() {
	configIfNotExist()

	cmd := "/opt/cisco/anyconnect/bin/vpn -s < " + c.configPath

	tio := &timeout.Timeout{
		Cmd:       exec.Command("sh", "-c", cmd),
		Duration:  5 * time.Second,
		KillAfter: 5 * time.Second,
	}

	exitStatus, out, _, err := tio.Run()
	if exitStatus.Signaled {
		fmt.Println("ERROR : Timeout is occurred. Please reconfiguration your setting by using `any-connect config`.")
		os.Exit(1)
	}

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(string(out))
}
