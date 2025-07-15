package rendering

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

type Website struct {
	Url       string
	Name      string
	ImageUrl  string
	IsFetched bool
}

const CacheUrl = "cache"

func GetCacheDir() (string, error) {
	baseDir, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}
	cacheDir := filepath.Join(baseDir, "startpage")
	err = os.MkdirAll(cacheDir, 0o755)
	if err != nil {
		return "", err
	}
	return cacheDir, nil
}

func hashUrl(url string) string {
	h := sha1.New()
	io.WriteString(h, url)
	return hex.EncodeToString(h.Sum(nil))
}

func (w *Website) Cache() error {
	if w.IsFetched {
		return nil
	}

	cacheDir, err := GetCacheDir()
	if err != nil {
		return err
	}

	filename := hashUrl(w.ImageUrl) + filepath.Ext(w.ImageUrl)
	targetPath := filepath.Join(cacheDir, filename)

	// if the file was already downloaded it doesn't need to be downloaded again
	if _, err := os.Stat(targetPath); errors.Is(err, os.ErrNotExist) {
		resp, err := http.Get(w.ImageUrl)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		file, err := os.Create(targetPath)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(file, resp.Body)

		if err != nil {
			return err
		}
	}

	// set the value in the struct to the current file
	w.ImageUrl, _ = url.JoinPath(CacheUrl, filename)
	w.IsFetched = true
	return nil
}
