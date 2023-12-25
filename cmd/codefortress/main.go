package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use: "codefortress",
	}

	var installCmd = &cobra.Command{
		Use:   "install",
		Short: "Install a tool",
		Long:  "Install a tool provided as an argument.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			tool := args[0]
			installTool(tool)
		},
	}

	rootCmd.AddCommand(installCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func installTool(tool string) {
	switch tool {
	case "git":
		if isGitInstalled() {
			fmt.Println("Git is already installed.")
		} else {
			installGit()
		}
	default:
		log.Printf("Unknown tool: %s\n", tool)
	}
}

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
		fmt.Errorf("Error installing git: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Git has been installed successfully.")
}
