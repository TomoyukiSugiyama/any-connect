package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func connect() {
	configIfNotExist()

	cmd := "/opt/cisco/anyconnect/bin/vpn -s < " + c.configPath
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(string(out))
}
