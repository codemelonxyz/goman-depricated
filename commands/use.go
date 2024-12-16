
package commands

import (
    "fmt"
    "os"
    "path/filepath"

    "goman/utils"
)

func Use(args []string) error {
    if len(args) != 1 {
        fmt.Println("Usage: goman use <version>")
        os.Exit(1)
    }
    version := args[0]
    return switchGoVersion(version)
}

func switchGoVersion(version string) error {
    // Check if version is installed
    versionPath := filepath.Join(os.ExpandEnv(utils.GoInstallDir), version)
    if _, err := os.Stat(versionPath); os.IsNotExist(err) {
        return fmt.Errorf("Go version %s is not installed. Install it first using 'goman install %s'", version, version)
    }

    // Expand the current link path
    currentPath := os.ExpandEnv(utils.CurrentLink)

    // Remove existing symlink if it exists
    if _, err := os.Lstat(currentPath); err == nil {
        if err := os.Remove(currentPath); err != nil {
            return fmt.Errorf("failed to remove existing symlink: %v", err)
        }
    }

    // Ensure parent directory exists
    if err := os.MkdirAll(filepath.Dir(currentPath), 0755); err != nil {
        return fmt.Errorf("failed to create parent directory: %v", err)
    }

    // Create new symlink
    if err := os.Symlink(versionPath, currentPath); err != nil {
        return fmt.Errorf("failed to create symlink: %v", err)
    }

    fmt.Printf("Switched to Go version %s\n", version)
    return nil
}