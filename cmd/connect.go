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

var c *Config

func init() {
	c = New()
}

type Config struct {
	configDir  string
	configPath string
	configName string
	hostName   string
	groupName  string
	userName   string
	password   string
}

func New() *Config {
	c := new(Config)
	home, _ := os.UserHomeDir()
	c.configDir = home + "/.cisco_vpn"
	c.configName = "credentials"
	c.configPath = c.configDir + "/" + c.configName
	c.hostName = ""
	c.groupName = ""
	c.userName = ""
	c.password = ""
	return c
}

func createConfigDir(c Config) {
	if err := os.Mkdir(c.configDir, 0755); err != nil {
		log.Fatal(err)
	}
}

func setConfig(c Config) {
	f, err := os.OpenFile(c.configPath, os.O_RDWR|os.O_CREATE, 0600)
	f.WriteString("connect " + c.hostName + "\n")
	f.WriteString("y\n")
	f.WriteString(c.groupName + "\n")
	f.WriteString(c.userName + "\n")
	f.WriteString(c.password + "\n")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
}

func selectHost(c *Config) {
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
	c.hostName = result
}

func setGroup(c *Config) {

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

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	c.groupName = result
}

func setUser(c *Config) {

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

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	c.userName = result
}

func setPassword(c *Config) {
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

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	c.password = result
}

func connect() {
	if _, err := os.Stat(c.configDir); os.IsNotExist(err) {
		createConfigDir(*c)
	}

	if f, err := os.Stat(c.configPath); os.IsNotExist(err) || f.IsDir() {
		selectHost(c)
		setGroup(c)
		setUser(c)
		setPassword(c)
		setConfig(*c)
	}

	cmd := "/opt/cisco/anyconnect/bin/vpn -s < " + c.configPath
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(string(out))
}
