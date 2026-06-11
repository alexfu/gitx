package repo

import (
	"errors"
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

	path := extensionPath(name)
	if path == "" {
		return "", errors.New("failed to resolve extension path")
	}

	_, err := git.PlainClone(path, &git.CloneOptions{
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

func extensionPath(name string) string {
	extensionDir := config.GitxDataDir()
	return filepath.Join(extensionDir, name)
}

func isFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	return !info.IsDir()
}
