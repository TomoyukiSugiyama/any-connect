package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func connect() {
	home, _ := os.UserHomeDir()
	cisco := home + "/.cisco_vpn/"
	config := "credentials"
	cmd := "/opt/cisco/anyconnect/bin/vpn -s < " + cisco + config
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(string(out))
}
