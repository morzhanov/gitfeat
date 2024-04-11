package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: gitfeat <task_name>")
		os.Exit(1)
	}

	taskName := os.Args[1]

	// Git stash
	cmd := exec.Command("git", "stash")
	if err := cmd.Run(); err != nil {
		fmt.Println("Git stash failed. Do you want to continue without stashing? (y/n)")
		reader := bufio.NewReader(os.Stdin)
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		if choice != "y" && choice != "Y" {
			fmt.Println("Aborted.")
			os.Exit(0)
		}
	}

	// Checkout to master
	cmd = exec.Command("git", "checkout", "master")
	runCmd(cmd)

	// Reset hard master
	cmd = exec.Command("git", "reset", "--hard", "origin/master")
	runCmd(cmd)

	// Pull master
	cmd = exec.Command("git", "pull", "origin", "master")
	runCmd(cmd)

	// Checkout to a new branch
	cmd = exec.Command("git", "checkout", "-b", taskName)
	runCmd(cmd)

	fmt.Printf("Preparation completed successfully. Switched to a new task branch %s", taskName)
}

func runCmd(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		os.Exit(1)
	}
}
