package main

import (
	"fmt"
	"os"
	"os/exec"
)

func isGitInstalled() bool {
	cmd := exec.Command("git", "--version")
	err := cmd.Run()

	return err == nil
}

func installGit() {
	fmt.Println("Installing Git...")

	cmd := exec.Command("go", "get", "-u", "github.com/git/git")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error installing git: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Git has been installed successfully.")
}
