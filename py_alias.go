package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"

	"github.com/google/uuid"
)

// func getTimestamp() string {
// 	return time.Now().Format("20060102-150405")
// }

func snapshotDirectory(prog_error string, uuid string) error {
	var appDataDir string
	usr, err := user.Current()
    if err != nil {
		fmt.Print("Error getting user directory:", err)
    }

	fmt.Println("User directory:", usr.HomeDir)
	fmt.Println("OS:", runtime.GOOS)

    if runtime.GOOS == "windows" {
        appDataDir = usr.HomeDir + "\\AppData\\Local\\PythonCapture"
    } else if runtime.GOOS == "darwin" { // macOS
        appDataDir = usr.HomeDir + "/Library/Application Support/PythonCapture"
    } else {
        // Handle other OS if needed
        appDataDir = usr.HomeDir + "/.PythonCapture" // Fallback to home directory
    }

	// Ensure the directory exists
	err = os.MkdirAll(appDataDir, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return err
	}

	// Create or open the snapshot text file
	snapshotFilePath := filepath.Join(appDataDir, uuid+"_context.txt")
	errorFilePath := filepath.Join(appDataDir, uuid+"_error.txt")
	snapshotFile, err := os.Create(snapshotFilePath)
	if err != nil {
		fmt.Println("Error creating snapshot file:", err)
		return err
	}
	errorFile, err := os.Create(errorFilePath)
	if err != nil {
		fmt.Println("Error creating error file:", err)
		return err
	}
	defer snapshotFile.Close()
	defer errorFile.Close()


	// Walk through all files in the directory
	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip the snapshot file itself
		if path == "snapshot.txt" || info.IsDir() || filepath.Ext(path) != ".py" {
			return nil
		}

		_, err = snapshotFile.WriteString("FILE NAME: " + path + "\n")
		if err != nil {
			return err
		}

		// Read and write the file contents to the snapshot file
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		_, err = snapshotFile.Write(content)
		if err != nil {
			return err
		}

		// Write a separator for the next file
		_, err = snapshotFile.WriteString("\n\n")
		if err != nil {
			return err
		}

		return nil
	})

	errorFile.WriteString(prog_error + "\n")
	return nil

}

func runApp(args []string) {
	// If no arguments are provided, start the Python REPL
	if len(args) < 1 {
		cmd := exec.Command("python3")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Run()
		return
	}

	// Get the Python script and its arguments from the command line
	pythonScript := args[0]
	pythonArgs := args[1:]

	// Prepare the command and its arguments
	uuid := uuid.New().String()
	cmdArgs := append([]string{pythonScript}, pythonArgs...)
	cmd := exec.Command("python3", cmdArgs...)
	var stderr bytes.Buffer
	var stdout bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
	
	// Run the Python script and capture stderr
	err := cmd.Run()
	fullOutput := stdout.String()
	fmt.Print(fullOutput)
	fmt.Fprint(os.Stderr, stderr.String())
	fmt.Printf("👍\n")

	// Check if the Python script ran successfully
	if err != nil {
		// Take a snapshot of the current directory
		err := snapshotDirectory(stderr.String(), uuid)
		if err != nil {
			fmt.Println("Python Wrapper Failed. No Errors were captured:", err)
		}

		// At this point, you would store pythonError and the snapshot in a database
	} else {
		fmt.Printf("👍\n")
	}
}