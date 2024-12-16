package core

import (
    "log"
    "os/exec"
    "strings"
)

type GoInfo struct {
    Version   string
    Processes string
}

func getGoVersion() (string, error) {
    output, err := exec.Command("go", "version").Output()
    if err != nil {
        return "", err
    }
    return string(output), nil
}

func getGoProcesses() (string, error) {
    cmd := exec.Command("ps", "aux")
    output, err := cmd.Output()
    if err != nil {
        return "", err
    }

    grep := exec.Command("grep", "go")
    grep.Stdin = strings.NewReader(string(output))
    filtered, err := grep.Output()
    if err != nil {
        return "No Go processes running", nil
    }
    return string(filtered), nil
}

func GetGoInfo() GoInfo {
    // Get Go version
    version, err := getGoVersion()
    if err != nil {
        log.Fatal("Error getting Go version:", err)
    }

    // Get running processes
    processes, err := getGoProcesses()
    if err != nil {
        log.Fatal("Error getting Go processes:", err)
    }

    return GoInfo{
        Version:   version,
        Processes: processes,
    }
}