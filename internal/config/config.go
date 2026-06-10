package config

import (
	"errors"
	"os"
	"path"
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
