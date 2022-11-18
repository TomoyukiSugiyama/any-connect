package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func disconnect() {
	cmd := "/opt/cisco/anyconnect/bin/vpn disconnect"
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(string(out))
}
