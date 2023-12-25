package main

import (
	"bufio"
	"fmt"
	logger "github.com/sirupsen/logrus"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func installSSH() {
	logger.Println("Installing SSH...")

	cmd := exec.Command("go", "get", "-u", "golang.org/x/crypto/ssh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		logger.Errorf("Error installing SSH: %v", err)
		os.Exit(1)
	}

	logger.Println("SSH has been installed successfully.")
}

func isSSHInstalled() bool {
	cmd := exec.Command("ssh", "-V")
	err := cmd.Run()

	return err == nil
}

func runSSH() {
	logger.Println("Configuring SSH...")

	// Get user input for SSH operation
	operation := getUserInput("Do you want to generate a new SSH key (type 'new') or add an existing one (type 'add')? ")

	switch strings.ToLower(operation) {
	case "new":
		runCreateSSH()
	case "add":
		runAddSSH()
	default:
		logger.Println("Invalid operation. Please type 'new' to generate a new SSH key or 'add' to add an existing one.")
		os.Exit(1)
	}
}

func runCreateSSH() {
	logger.Println("Creating a new SSH key...")

	// Get user input for new SSH key parameters
	fileName := getUserInput("Enter the file name for the new SSH key (e.g., id_rsa): ")
	keyType := getUserInput("Enter the type of the new SSH key (e.g., rsa): ")
	keySizeStr := getUserInput("Enter the key size for the new SSH key (e.g., 2048): ")
	password := getUserInput("Enter the password for the new SSH key (press Enter for no password): ")

	// Convert keySize to integer
	keySize, err := strconv.Atoi(keySizeStr)
	if err != nil {
		logger.Errorf("Invalid key size. Please enter a valid number.")
		os.Exit(1)
	}

	// Create the .ssh folder if it doesn't exist
	sshFolder := filepath.Join(os.Getenv("HOME"), ".ssh")
	if _, err := os.Stat(sshFolder); os.IsNotExist(err) {
		err := os.Mkdir(sshFolder, 0700)
		if err != nil {
			logger.Errorf("Error creating .ssh folder: %v\n", err)
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
		logger.Errorf("Error running SSH: %v\n", err)
		os.Exit(1)
	}

	logger.Printf("New SSH key has been generated successfully and saved to: %s\n", keyPath)
}

func runAddSSH() {
	logger.Println("Adding an existing SSH key...")

	// Get user input for existing SSH key parameters
	existingKeyPath := getUserInput("Enter the path to the existing SSH key file: ")

	// Check if the file exists
	if _, err := os.Stat(existingKeyPath); os.IsNotExist(err) {
		logger.Errorf("Error: File %s does not exist.\n", existingKeyPath)
		os.Exit(1)
	}

	// Create the .ssh folder if it doesn't exist
	sshFolder := filepath.Join(os.Getenv("HOME"), ".ssh")
	if _, err := os.Stat(sshFolder); os.IsNotExist(err) {
		err := os.Mkdir(sshFolder, 0700)
		if err != nil {
			logger.Errorf("Error creating .ssh folder: %v\n", err)
			os.Exit(1)
		}
	}

	// Copy the existing key to the .ssh folder
	newKeyPath := filepath.Join(sshFolder, filepath.Base(existingKeyPath))

	err := copyFile(existingKeyPath, newKeyPath)
	if err != nil {
		logger.Errorf("Error copying existing SSH key: %v\n", err)
		os.Exit(1)
	}

	logger.Printf("Existing SSH key has been imported successfully and saved to: %s\n", newKeyPath)
}

func copyFile(src, dest string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("unable to open file: %w", err)
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("unable to create file: %w", err)
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return fmt.Errorf("unable to copy file: %w", err)
	}

	return nil
}

func getUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(prompt)

	input, _ := reader.ReadString('\n')

	return strings.TrimSpace(input)
}
