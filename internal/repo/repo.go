package repo

import (
	"io/fs"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"gitx/internal/config"

	"github.com/go-git/go-git/v6"
)

func DownloadExtention(name string) (string, error) {
	url := url.URL{
		Scheme: "https",
		Host:   "github.com",
		Path:   name,
	}

	path, err := extensionPath(name)
	if err != nil {
		return "", err
	}

	_, err = git.PlainClone(path, &git.CloneOptions{
		URL: url.String(),
	})
	if err != nil {
		return "", err
	}

	return path, nil
}

func InstallExtension(path string) error {
	gitxhome, err := config.EnsureGitxHome()
	if err != nil {
		return err
	}

	filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		basename := filepath.Base(path)
		if strings.HasPrefix(basename, "git-") && isFile(path) {
			os.Symlink(path, filepath.Join(gitxhome, basename))
		}
		return nil
	})
	return nil
}

func extensionPath(name string) (string, error) {
	extensionDir, err := dataDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(extensionDir, name), nil
}

func dataDir() (string, error) {
	if xdgDataHome := os.Getenv("XDG_DATA_HOME"); xdgDataHome != "" {
		return filepath.Join(xdgDataHome, "gitx"), nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, ".local", "share", "gitx"), nil
}

func isFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	return !info.IsDir()
}
