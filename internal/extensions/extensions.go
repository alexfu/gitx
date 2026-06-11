package extensions

import (
	"errors"
	"io/fs"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"gitx/internal/config"
	"gitx/internal/fsutils"

	"github.com/go-git/go-git/v6"
)

func Download(name string) (string, error) {
	url := url.URL{
		Scheme: "https",
		Host:   "github.com",
		Path:   name,
	}

	path := extensionDataPath(name)
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

func Install(path string) error {
	gitxhome, err := config.EnsureGitxHome()
	if err != nil {
		return err
	}

	filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		basename := filepath.Base(path)
		if strings.HasPrefix(basename, "git-") && !d.IsDir() {
			os.Symlink(path, filepath.Join(gitxhome, basename))
		}
		return nil
	})
	return nil
}

type RemoveInfo struct {
	DeleteList []string
	Success    []string
	Failed     []string
}

func Remove(name string, confirmFunc func(deleteList []string) bool) RemoveInfo {
	extDataPath := extensionDataPath(name)
	info := RemoveInfo{}

	// Collect files to delete
	filepath.WalkDir(config.GitxHome(), func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			realPath, err := filepath.EvalSymlinks(path)
			if err == nil && strings.HasPrefix(realPath, extDataPath) {
				info.DeleteList = append(info.DeleteList, path)
			}
		}
		return nil
	})

	// If it's a dir, clean up source as well
	if fsutils.IsDir(extDataPath) {
		info.DeleteList = append(info.DeleteList, extDataPath)
	}

	// Nothing to delete
	if len(info.DeleteList) == 0 {
		return info
	}

	if confirmFunc(info.DeleteList) {
		for _, path := range info.DeleteList {
			err := os.RemoveAll(path)
			if err != nil {
				info.Failed = append(info.Failed, path)
			} else {
				info.Success = append(info.Success, path)
			}
		}
		return info
	}

	return info
}

func extensionDataPath(name string) string {
	datadir := config.GitxDataDir()
	return filepath.Join(datadir, name)
}
