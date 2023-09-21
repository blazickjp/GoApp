package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/google/uuid"
)

// func getTimestamp() string {
// 	return time.Now().Format("20060102-150405")
// }


func snapshotDirectory(prog_error string, uuid string) error {
	// Create or open the snapshot text file
	snapshotFile, err := os.Create(uuid + "_context.txt")
	if err != nil {
		fmt.Println("Error creating snapshot file:", err)
		return err
	}
	errorFile, err := os.Create(uuid + "_error.txt")
	if err != nil {
		fmt.Println("Error creating error file:", err)
		return err
	}
	defer snapshotFile.Close()

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