
package commands

import (
	"fmt"
	"os"

	"goman/utils"
)

func Purge() error {
	installPath := os.ExpandEnv(utils.GoInstallDir)

	// Check if installation directory exists
	if _, err := os.Stat(installPath); os.IsNotExist(err) {
		fmt.Println("No Go versions installed")
		return nil
	}

	// Remove entire installation directory
	if err := os.RemoveAll(installPath); err != nil {
		return fmt.Errorf("failed to purge Go versions: %v", err)
	}

	fmt.Println("Removed all installed Go versions")
	return nil
}