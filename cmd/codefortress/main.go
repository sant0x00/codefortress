package main

import (
	"fmt"
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
	case "ssh":
		if isSSHInstalled() {
			fmt.Println("SSH is already installed.")
		} else {
			installSSH()
		}

		runSSH()
	default:
		fmt.Printf("Unknown tool: %s\n", tool)
	}
}
