/**
 * @Author: harshkhajuria
 * @Date:   01-Jul-2019 06:58:28 am
 * @Email:  khajuriaharsh729@gmail.com
 * @Filename: potato.go
 * @Last modified by:   harshkhajuria
 * @Last modified time: 10-Jul-2019 07:01:44 am
 */

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Setting struct {
	Name, Description, Command, Filename string
}

var settings map[string]Setting

func reader(toRead string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(toRead)
	inputText, _ := reader.ReadString('\n')
	inputText = strings.TrimSuffix(inputText, "\n")
	return inputText
}

func readJson() {

  dir, direrr := os.Getwd()
  if direrr != nil {
    fmt.Println(direrr)
		return
  }
  dir = dir + "/settings.json"
	fi, fierr := os.Stat(dir)
	if fierr != nil {
		fmt.Println(fierr)
		return
	}

	if fi.Size() == 0 {
		return
	}

	data, err := ioutil.ReadFile("settings.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(data, &settings)
	if err != nil {
		fmt.Println(err)
	}

}

func writeJson() {
	file, err := json.Marshal(settings)
	if err != nil {
		fmt.Println(err)
	}
  dir, direrr := os.Getwd()
  if direrr != nil {
    fmt.Println(direrr)
    return
  }
  dir = dir + "/settings.json"
	_ = ioutil.WriteFile(dir, file, 0644)
}

func addSettingsHelper() {
	tempName := reader("Enter a name for the setting(1 to 10 characters) ")
	_, ok := settings[tempName]
	if ok == true {
		fmt.Println("The given key already exists")
		return
	}
	tempDescription := reader("Enter a short description for the setting (can be empty) ")
	tempCommand := reader("Enter the command that needs to be executed for this setting ")
	tempFilename := reader("Enter the filename in whih these settings will be saved. If empty name of the setting will be used. ")
	tempFileNameLen := len([]rune(tempFilename))
	if tempFileNameLen == 0 {
		tempFilename = tempName
	}
	settings[tempName] = Setting{
		tempName, tempDescription, tempCommand, tempFilename,
	}
}

func addSetting() {
	readJson()
	addSettingsHelper()
	writeJson()
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func deleteSetting(key string) {
	_, ok := settings[key]
	if ok == false {
		fmt.Println("No such setting exists")
	} else {
		fmt.Println("Deleting setting - " + key)
	}
	delete(settings, key)
	writeJson()
}

func displaySettings() {
	for key, _ := range settings {
		tempSetting := settings[key]
		if key == "Git" {
			fmt.Printf("Git Directory: %s \n", tempSetting.Description)
			fmt.Println()
			continue
		}
		fmt.Printf("Name: %s \n", tempSetting.Name)
		fmt.Printf("Description: %s \n", tempSetting.Description)
		fmt.Printf("Command: %s \n", tempSetting.Command)
		fmt.Printf("Filename: %s \n", tempSetting.Filename)
		fmt.Println()
	}
}

func setUpGit() {
	dir := reader("Enter the path to the Git Directory ")
	settings["Git"] = Setting{
		"Git", dir, "", "",
	}
	writeJson()
}

func writeToFile(filename string, data []byte) {
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}

func execCommand(command string, args []string) []byte {
	cmd := exec.Command(command, args...)
	stdoutStderr, err := cmd.CombinedOutput()
	data := []byte(stdoutStderr)
	if err != nil {
		fmt.Printf("%s\n", stdoutStderr)
	}
	return data
}

func isGitDirHelper(dir string) bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	cmd.Dir = dir
	_, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}
	return true
}

func gitPush(dir string) {
	fmt.Println("Creating GitHub backup")
	cmd := exec.Command("git", "add", ".")
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error while executing git add .")
		fmt.Println(string(out))
		fmt.Println(err)
		return
	}
	now := time.Now()
	cmd = exec.Command("git", "commit", "-m", "Updated settings - "+now.Format("2006-01-02 15:04:05"))
	cmd.Dir = dir
	out, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error while making a new commit")
		fmt.Println(string(out))
		return
	}
	cmd = exec.Command("git", "push", "origin", "master")
	cmd.Dir = dir
	out, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error while pushing the commit to remote")
		fmt.Println(string(out))
		return
	}
}

func parseCommand(command string) []string {
	var commands []string
	length := len([]rune(command))
	if length == 0 {
		return commands
	}
	var temp string
	temp = string(command[0])
	for i := 1; i < length; i++ {
		if string(command[i]) == "\\" {
			continue
		}
		if string(command[i]) != " " {
			temp = temp + string(command[i])
		} else {
			if string(command[i-1]) != "\\" {
				commands = append(commands, temp)
				temp = ""
			} else {
				temp += " "
			}
		}
	}
	commands = append(commands, temp)
	return commands
}

func runSettings() {
	_, ok := settings["Git"]
	if ok == false {
		fmt.Println("Set up a git directory before running")
		return
	}
	filename := settings["Git"].Description
	fileNameLen := len([]rune(filename))
	if fileNameLen == 0 {
		fmt.Println("Set up a git directory before running")
		return
	}
	for key, _ := range settings {
		if key == "Git" {
			continue
		}
		tempSetting := settings[key]
		tempCommandArguments := parseCommand(tempSetting.Command)
		fmt.Printf("Executing for setting : %s \n", tempSetting.Name)
		data := execCommand(tempCommandArguments[0], tempCommandArguments[1:])
		filename = settings["Git"].Description + "/" + tempSetting.Filename
		writeToFile(filename, data)
		fmt.Println("Done for " + tempSetting.Name)
		fmt.Println()
	}
	// fmt.Printf("Git dir %s\n", settings["Git"].Description)
	if isGitDirHelper(settings["Git"].Description) == false {
		fmt.Println("The directory " + settings["Git"].Description + "is not a git repo")
		fmt.Println("Run git init and set a remote tracking in it to create GitHub backups")
		return
	}
	gitPush(settings["Git"].Description)
}

func main() {
	app := cli.NewApp()

	app.EnableBashCompletion = true
	app.Name = "Potato"
	app.Usage = "Keep track of your system as a developer"

	settings = make(map[string]Setting)

	app.Flags = []cli.Flag{}

	app.Commands = []cli.Command{
		{
			Name:  "add",
			Usage: "Add a setting",
			Action: func(c *cli.Context) error {
				fmt.Println("Add a setting : ")
				addSetting()
				return nil
			},
		},
		{
			Name:  "delete",
			Usage: "Delete a setting",
			Action: func(c *cli.Context) error {
				fmt.Println("Delete a setting : ")
				tempName := reader("Enter the name of the string to be deleted: ")
				readJson()
				deleteSetting(tempName)
				return nil
			},
		},
		{
			Name:  "display",
			Usage: "Display all settings",
			Action: func(c *cli.Context) error {
				fmt.Println("Display all settings : ")
				readJson()
				displaySettings()
				return nil
			},
		},
		{
			Name:  "run",
			Usage: "Run the package",
			Action: func(c *cli.Context) error {
				fmt.Println("Creating backup : ")
				readJson()
				runSettings()
				return nil
			},
		},
		{
			Name:  "git",
			Usage: "Set/Update Git Directory where the packages will be backed up",
			Action: func(c *cli.Context) error {
				readJson()
				setUpGit()
				return nil
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Println("Running potato")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
