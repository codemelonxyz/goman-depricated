
package commands

import (
    "fmt"
    "os"
    "path/filepath"

    "goman/utils"
)

func Remove(args []string) error {
    if len(args) != 1 {
        fmt.Println("Usage: goman remove <version>")
        os.Exit(1)
    }
    version := args[0]
    versionPath := filepath.Join(os.ExpandEnv(utils.GoInstallDir), version)
    
    // Check if version exists
    if _, err := os.Stat(versionPath); os.IsNotExist(err) {
        return fmt.Errorf("Go version %s is not installed", version)
    }

    // Remove the version directory
    if err := os.RemoveAll(versionPath); err != nil {
        return fmt.Errorf("failed to remove Go version %s: %v", err)
    }

    fmt.Printf("Removed Go version %s\n", version)
    return nil
}