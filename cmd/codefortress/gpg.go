package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func installSSH() {
	fmt.Println("Installing SSH...")

	cmd := exec.Command("go", "get", "-u", "golang.org/x/crypto/ssh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error installing SSH: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("SSH has been installed successfully.")
}

func isSSHInstalled() bool {
	cmd := exec.Command("ssh", "-V")
	err := cmd.Run()

	return err == nil
}

func runSSH() {
	fmt.Println("Running SSH...")

	// Get user input for SSH parameters
	fileName := getUserInput("Enter the file name for the SSH key (e.g., id_rsa): ")
	keyType := getUserInput("Enter the type of the SSH key (e.g., rsa): ")
	keySizeStr := getUserInput("Enter the key size for the SSH key (e.g., 2048): ")
	password := getUserInput("Enter the password for the SSH key (press Enter for no password): ")

	// Convert keySize to integer
	keySize, err := strconv.Atoi(keySizeStr)
	if err != nil {
		fmt.Println("Invalid key size. Please enter a valid number.")
		os.Exit(1)
	}

	// Create the .ssh folder if it doesn't exist
	sshFolder := filepath.Join(os.Getenv("HOME"), ".ssh")
	if _, err := os.Stat(sshFolder); os.IsNotExist(err) {
		err := os.Mkdir(sshFolder, 0700)
		if err != nil {
			fmt.Printf("Error creating .ssh folder: %v\n", err)
			os.Exit(1)
		}
	}

	// Run ssh-keygen with the provided parameters
	keyPath := filepath.Join(sshFolder, fileName)
	cmd := exec.Command("ssh-keygen", "-t", keyType, "-b", strconv.Itoa(keySize), "-f", keyPath, "-N", password)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error running SSH: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("SSH key has been generated successfully and saved to: %s\n", keyPath)
}

func getUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')

	return strings.TrimSpace(input)
}
