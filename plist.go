package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

// createAndLoadPlist creates a .plist file for the application, moves it to the appropriate directory,
// and then loads the service using launchctl.
func createAndLoadPlist() error {
	// Step 1: Generate the .plist content
	plistContent := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>Label</key>
	<string>com.yourcompany.yourapp</string>
	<key>ProgramArguments</key>
	<array>
		<string>/path/to/your/app</string>
	</array>
	<key>RunAtLoad</key>
	<true/>
	<key>KeepAlive</key>
	<true/>
	<key>StandardOutPath</key>
	<string>/dev/null</string>
	<key>StandardErrorPath</key>
	<string>/dev/null</string>
	<key>NSHighResolutionCapable</key>
	<string>True</string>

	<!-- avoid showing the app on the Dock -->
	<key>LSUIElement</key>
	<string>1</string>
</dict>
</plist>`

	// Step 2: Write the .plist file to a temporary location
	tempFilePath := "/tmp/com.yourcompany.yourapp.plist"
    err := os.WriteFile(tempFilePath, []byte(plistContent), 0644)
    if err != nil {
        return fmt.Errorf("error writing .plist file: %v", err)
    }

    // Step 3: Move the .plist file to ~/Library/LaunchAgents/
    homeDir, err := os.UserHomeDir()
    if err != nil {
        return fmt.Errorf("error getting home directory: %v", err)
    }
    destFilePath := homeDir + "/Library/LaunchAgents/com.yourcompany.yourapp.plist"

    cmd := exec.Command("mv", tempFilePath, destFilePath)
    var stderr bytes.Buffer
    cmd.Stderr = &stderr
    err = cmd.Run()
    if err != nil {
        return fmt.Errorf("error moving .plist file: %v, stderr: %s", err, stderr.String())
    }

	// Step 4: Load the service using launchctl
	// This may also require administrative permissions
	cmd = exec.Command("launchctl", "load", destFilePath)
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("error loading service: %v", err)
	}

	return nil
}