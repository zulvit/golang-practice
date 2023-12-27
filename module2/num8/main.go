package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")

		if input == "\\quit" {
			break
		}

		args := strings.Fields(input)
		switch args[0] {
		case "cd":
			if len(args) < 2 {
				fmt.Println("Path required")
				continue
			}
			err := os.Chdir(args[1])
			if err != nil {
				fmt.Println("Error:", err)
			}
		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Println(dir)
		case "echo":
			fmt.Println(strings.Join(args[1:], " "))
		case "kill":
			if len(args) < 2 {
				fmt.Println("PID required")
				continue
			}
			pid, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("Invalid PID")
				continue
			}
			process, err := os.FindProcess(pid)
			if err != nil {
				fmt.Println("Process not found")
				continue
			}
			err = process.Kill()
			if err != nil {
				fmt.Println("Error killing process:", err)
			} else {
				fmt.Println("Process killed")
			}
		case "ps":
			cmd := exec.Command("ps")
			cmd.Stdout = os.Stdout
			err := cmd.Run()
			if err != nil {
				fmt.Println("Error executing ps:", err)
			}
		default:
			cmd := exec.Command(args[0], args[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println("Error executing command:", err)
			}
		}
	}
}
