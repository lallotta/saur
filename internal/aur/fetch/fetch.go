package fetch

import (
	"bytes"
	"fmt"
	"os/exec"
)

const baseURL = "https://aur.archlinux.org/"

// runGit runs a git command
func runGit(args ...string) error {
	var stderr bytes.Buffer

	cmd := exec.Command("git", args...)
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s", &stderr)
	}

	return nil
}

// gitClone runs a git clone command
func gitClone(url string) error {
	return runGit("clone", url)
}

// GetPkgbuild retrieves the PKGBUILD for the given package
func GetPkgbuild(pkgName string) error {
	return gitClone(baseURL + pkgName + ".git")
}
