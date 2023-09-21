package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/getlantern/systray"
)

func getShellType() string {
	shell := os.Getenv("SHELL")
	if shell == "" {
		return "unknown"
	}
	return shell
}

func onReady() {
	// Get the path of the currently running program
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable path:", err)
		// Handle error appropriately
	}
	fmt.Print("Starting the systray app\n")
	systray.SetTitle("My Go App")
	systray.SetTooltip("Right-click to see options")

    mSetAlias := systray.AddMenuItem("Set Alias", "Set the alias name")
	mRemoveAlias := systray.AddMenuItem("Remove Alias", "Remove the alias name")  // New line
    mShell := systray.AddMenuItem("Shell", "Choose shell type")
    mBash := mShell.AddSubMenuItemCheckbox("Bash", "Use Bash", false)
    mZsh := mShell.AddSubMenuItemCheckbox("Zsh", "Use Zsh", false)
    mQuit := systray.AddMenuItem("Quit", "Quit the application")

    shell := getShellType()
    shellType := "bashrc"  // default
    if shell == "/bin/zsh" {
        shellType = "zshrc"
    } else if shell == "/bin/bash" {
        shellType = "bashrc"
    }

    for {
        select {
		case <-mBash.ClickedCh:
            mBash.Check()
            mZsh.Uncheck()
            // Logic to switch to bash
        case <-mZsh.ClickedCh:
            mZsh.Check()
            mBash.Uncheck()
        case <-mSetAlias.ClickedCh:
            // Logic to set the alias
            // For demonstration, using a shell command
            cmd := exec.Command("bash", "-c", fmt.Sprintf(`echo 'alias %s="%s"' >> ~/.%s`, "py", exePath, shellType))
            cmd.Run()
		case <-mRemoveAlias.ClickedCh:  // New block
            cmd := exec.Command("bash", "-c", fmt.Sprintf(`sed -i '' '/alias %s/d' ~/.%s`, "py", shellType))
            cmd.Run()
		case <-mQuit.ClickedCh:
			err := os.Remove("/tmp/my_app.lock")
			if err != nil {
				fmt.Println("Error removing lock file:", err)
			}
			systray.Quit()
			return
		}
		fmt.Print("Waiting for user input\n")
	}
}

func main() {
	onExit := func() {
		os.Remove("/tmp/my_app.lock")
	}
	// Create and load plist only if it doesn't exist
	plistPath := "~/Library/LaunchAgents/com.yourcompany.yourapp.plist"
	if _, err := os.Stat(plistPath); os.IsNotExist(err) {
		err := createAndLoadPlist()
		if err != nil {
			fmt.Println("Error creating and loading plist:", err)
			return
		}
	}	
	// lockFilePath := "/tmp/my_app.lock"
	// lockFile, err := os.OpenFile(lockFilePath, os.O_CREATE|os.O_EXCL, 0644)
	// if err != nil {
	// 	fmt.Println("Another instance is already running.")
	// 	return
	// }
	// defer os.Remove(lockFilePath)
	// defer lockFile.Close()
	if len(os.Args) > 1 && os.Args[1] == "runApp" {
		// Remove the "runApp" argument and pass the rest to runApp
		runApp(os.Args[2:])
		return
	} else {
		systray.Run(onReady, onExit)
	}

}
