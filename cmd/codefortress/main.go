package main

import (
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
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

	var createCmd = &cobra.Command{
		Use:   "create ssh",
		Short: "Create a new SSH key",
		Long:  "Create a new SSH key and save it.",
		Run: func(cmd *cobra.Command, args []string) {
			runCreateSSH()
		},
	}

	var addCmd = &cobra.Command{
		Use:   "add ssh",
		Short: "Add an existing SSH key",
		Long:  "Add an existing SSH key to the .ssh folder.",
		Run: func(cmd *cobra.Command, args []string) {
			runAddSSH()
		},
	}

	rootCmd.AddCommand(installCmd, createCmd, addCmd)

	if err := rootCmd.Execute(); err != nil {
		logger.Errorf("unable to execute Codefortress. Here the reason: %s\n", err)
		os.Exit(1)
	}
}

func installTool(tool string) {
	switch tool {
	case "git":
		if isGitInstalled() {
			logger.Println("Git is already installed.")
		} else {
			installGit()
		}
	case "ssh":
		if isSSHInstalled() {
			logger.Println("SSH is already installed.")
		} else {
			installSSH()
		}

		runSSH()
	default:
		logger.Printf("Unknown tool: %s", tool)
	}
}
