package main

import (
	logger "github.com/sirupsen/logrus"
	"os"
	"os/exec"
)

func isGitInstalled() bool {
	cmd := exec.Command("git", "--version")
	err := cmd.Run()

	return err == nil
}

func installGit() {
	logger.Println("Trying installing Git...")

	cmd := exec.Command("go", "get", "-u", "github.com/git/git")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		logger.Errorf("Error installing Git: %s\n", err)
		os.Exit(1)
	}

	logger.Println("Git has been installed successfully!")
}
