package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"strings"
	"os/exec"
	"os"
)

func main() {
	theme := os.Args[1]
	parse(theme)
	
}

func parse(theme string) {
	home, err := os.UserHomeDir()
	checkErr(err)
	
	file, err := ioutil.ReadFile(home + "/.zshrc")
	checkErr(err)

	lines := strings.Split(string(file), "\n")

	for i, line := range lines {
		if strings.Contains(line, "ZSH_THEME=\"") {
			lines[i] = "ZSH_THEME=\""+ theme + "\""
		}
	}

	edited := strings.Join(lines, "\n")

	err = ioutil.WriteFile(home + "/.zshrc", []byte(edited), 0644)
	checkErr(err)

	reloadZsh := [3]string{"zsh", "-c", "source ~/.zshrc"}
	newTerminalTab := [3]string{"zsh", "-c", "open -na Terminal"}
	executeShellCommand(reloadZsh)
	executeShellCommand(newTerminalTab)

}

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}


func executeShellCommand(args [3]string) {
	cmd := exec.Command(args[0], args[1], args[2])
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
		return
	}
	fmt.Println(string((output)))
}