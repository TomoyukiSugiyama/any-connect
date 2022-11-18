package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/manifoldco/promptui"
)

func createConfig() {
	home, _ := os.UserHomeDir()
	if err := os.Mkdir(home+"/.cisco__vpn", 0755); err != nil {
		log.Fatal(err)
	}
	f, err := os.Create(home + "/.cisco__vpn/credentials")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
}

func selectHost() {
	cmd := "/opt/cisco/anyconnect/bin/vpn hosts"
	out, _ := exec.Command("sh", "-c", cmd).Output()
	s := string(out)
	reg := "\r\n|\n"
	arr1 := regexp.MustCompile(reg).Split(s, -1)
	var items []string

	for _, s := range arr1 {
		if !strings.Contains(s, ">") {
			continue
		}
		re := regexp.MustCompile("\\s+").Split(s, -1)
		items = append(items, re[2])
	}
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    "What's your text editor",
			Items:    items,
			AddLabel: "Other",
		}

		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
		}
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
}

func setGroup() {

	validate := func(input string) error {
		if len(input) < 1 {
			return errors.New("Group must have more than 1 characters")
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | faint }} ",
	}

	prompt := promptui.Prompt{
		Label:     "Group",
		Validate:  validate,
		Templates: templates,
	}

	_, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
}

func setUser() {

	validate := func(input string) error {
		if len(input) < 1 {
			return errors.New("User must have more than 1 characters")
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | faint }} ",
	}

	prompt := promptui.Prompt{
		Label:     "User",
		Validate:  validate,
		Templates: templates,
	}

	_, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
}

func setPassword() {
	validate := func(input string) error {
		if len(input) < 1 {
			return errors.New("Password must have more than 1 characters")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Password",
		Validate: validate,
		Mask:     '*',
	}

	_, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
}

func connect() {
	// home, _ := os.UserHomeDir()
	// config_path := home + "/.cisco__vpn/"
	// config := config_path + "credentials"

	// f, err := os.Stat(config)
	// if os.IsNotExist(err) || f.IsDir() {
	// 	fmt.Println("not exist")
	// 	createConfig()
	// }

	selectHost()
	setGroup()
	setUser()
	setPassword()
	// cmd := "/opt/cisco/anyconnect/bin/vpn -s < " + cisco + config
	// out, err := exec.Command("sh", "-c", cmd).Output()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	os.Exit(1)
	// }
	// fmt.Println(string(out))
}
