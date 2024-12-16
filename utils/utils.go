package utils

import (
    "fmt"
    "os"
    "os/user"
    "path/filepath"
    "strings"
)

const (
    GoInstallDir = "$HOME/.goman/versions"
    CurrentLink  = "$HOME/.goman/current"
    ShellrcFiles = ".bashrc .zshrc .bash_profile .zprofile"
)

func PrintUsage() {
    fmt.Println("Goman - Go Manager  By Codemelon")
    fmt.Println("Go Version Manager:")
    fmt.Println("    goman install <version>   - Install a specific Go version")
    fmt.Println("    goman use <version>       - Switch to a specific Go version")
    fmt.Println("    goman list                - List installed Go versions")
    fmt.Println("    goman remove <version>    - Remove a specific Go version")
    fmt.Println("    goman purge               - Remove all installed Go versions")
    fmt.Println("    goman setup               - Setup Goman (update PATH)")
    fmt.Println("Options:")
    fmt.Println("    -h                        - Display help")
    fmt.Println("    -v                        - Display version")
	fmt.Println("    -c 					   - Contact the developer")
}

func UpdateShellConfig() error {
    // Get current user
    currentUser, err := user.Current()
    if err != nil {
        return fmt.Errorf("failed to get current user: %v", err)
    }

    // List of shell configuration files to check
    shellFiles := strings.Split(ShellrcFiles, " ")

    // Path to be added
    goPath := filepath.Join(os.ExpandEnv("$HOME"), ".goman", "current", "bin")
    pathLine := fmt.Sprintf("\n# Added by Goman\nexport PATH=$PATH:%s\n", goPath)

    // Track if configuration was updated
    updated := false

    // Check and update shell configuration files
    for _, filename := range shellFiles {
        filepath := filepath.Join(currentUser.HomeDir, filename)
        
        // Check if file exists
        if _, err := os.Stat(filepath); os.IsNotExist(err) {
            continue
        }

        // Read current contents
        content, err := os.ReadFile(filepath)
        if err != nil {
            fmt.Printf("Warning: Could not read %s: %v\n", filename, err)
            continue
        }

        // Check if path is already configured
        if !strings.Contains(string(content), goPath) {
            // Append path configuration
            f, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, 0644)
            if err != nil {
                fmt.Printf("Warning: Could not open %s: %v\n", filename, err)
                continue
            }
            defer f.Close()

            if _, err := f.WriteString(pathLine); err != nil {
                fmt.Printf("Warning: Could not update %s: %v\n", filename, err)
                continue
            }
            
            updated = true
            fmt.Printf("Updated %s\n", filename)
        }
    }

    if updated {
        fmt.Println("Goman PATH configuration updated. Please restart your terminal or run 'source ~/.{shell profile}' on your shell configuration file.")
    } else {
        fmt.Println("No shell configuration files needed updating.")
    }

    return nil
}