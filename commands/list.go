
package commands

import (
    "fmt"
    "os"

    "goman/utils"
)

func List() error {
    installPath := os.ExpandEnv(utils.GoInstallDir)
    
    // Check if installation directory exists
    if _, err := os.Stat(installPath); os.IsNotExist(err) {
        fmt.Println("No Go versions installed")
        return nil
    }

    // Read directory contents
    entries, err := os.ReadDir(installPath)
    if err != nil {
        return err
    }

    // Check if any versions are installed
    if len(entries) == 0 {
        fmt.Println("No Go versions installed")
        return nil
    }

    fmt.Println("Installed Go Versions:")
    for _, entry := range entries {
        if entry.IsDir() {
            fmt.Println(entry.Name())
        }
    }

    return nil
}