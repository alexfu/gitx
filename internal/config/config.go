package config

import (
	"errors"
	"os"
	"path"
	"path/filepath"
)

func EnsureGitxHome() (string, error) {
	home := GitxHome()
	if len(home) == 0 {
		return "", errors.New("unable to resolve gitx home")
	}
	os.MkdirAll(home, 0o755)
	return home, nil
}

func GitxHome() string {
	userhome, err := os.UserHomeDir()
	if err != nil || len(userhome) == 0 {
		return ""
	}
	return path.Join(userhome, ".gitx")
}

func GitxDataDir() string {
	if xdgDataHome := os.Getenv("XDG_DATA_HOME"); xdgDataHome != "" {
		return filepath.Join(xdgDataHome, "gitx")
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}

	return filepath.Join(home, ".local", "share", "gitx")
}
