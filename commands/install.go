
package commands

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"goman/utils"
)

func Install(args []string) error {
	if len(args) != 1 {
		fmt.Println("Usage: goman install <version>")
		os.Exit(1)
	}
	version := args[0]
	fmt.Println("Installing Go version", version)
	installPath := filepath.Join(os.ExpandEnv(utils.GoInstallDir), version)

	// Create installation directory
	if err := os.MkdirAll(installPath, 0755); err != nil {
		return err
	}

	// Construct download URL
	osName := runtime.GOOS
	archName := runtime.GOARCH
	fmt.Println("Downloading the specified version ... ")
	downloadURL := fmt.Sprintf("https://golang.org/dl/go%s.%s-%s.tar.gz", version, osName, archName)

	// Download Go archive
	resp, err := http.Get(downloadURL)
	if err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download: HTTP %d", resp.StatusCode)
	}

	// Create tar.gz file
	tarPath := filepath.Join(os.TempDir(), fmt.Sprintf("go%s.tar.gz", version))
	out, err := os.Create(tarPath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Save downloaded archive
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	// Open the tar.gz file
	f, err := os.Open(tarPath)
	if err != nil {
		return err
	}
	defer f.Close()

	// Create gzip reader
	gzr, err := gzip.NewReader(f)
	if err != nil {
		return err
	}
	defer gzr.Close()

	// Create tar reader
	tr := tar.NewReader(gzr)

	// Extract files
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		target := filepath.Join(installPath, strings.TrimPrefix(header.Name, "go/"))

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(target, 0755); err != nil {
				return err
			}
		case tar.TypeReg:
			if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
				return err
			}

			func() {
				outFile, err := os.Create(target)
				if err != nil {
					return
				}
				defer outFile.Close()

				if _, err := io.Copy(outFile, tr); err != nil {
					return
				}

				os.Chmod(target, os.FileMode(header.Mode))
			}()
		}
	}

	// Clean up tar file
	os.Remove(tarPath)

	fmt.Printf("Go %s installed successfully\n", version)
	return nil
}