package cache

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"gitea.elara.ws/Hazel/transfem-startpage/internal/utils"
	"github.com/labstack/echo/v4"
)

type Cache struct {
	CacheDir string
	Disabled bool
}

func getCacheDir() (string, error) {
	baseDir, err := os.UserCacheDir()
	if err != nil {
		baseDir = "/tmp"
	}
	cacheDir := filepath.Join(baseDir, utils.Name)
	err = os.MkdirAll(cacheDir, 0o755)
	if err != nil {
		return "", err
	}
	return cacheDir, nil
}

func getProfileCacheDir(profile string) (string, error) {
	var profileCacheDir string

	cacheDir, err := getCacheDir()
	if err != nil {
		return profileCacheDir, err
	}

	profileCacheDir = filepath.Join(cacheDir, profile)
	err = os.MkdirAll(cacheDir, 0o755)
	return profileCacheDir, err
}

func NewCache(profile string) Cache {
	cacheDir, err := getProfileCacheDir(profile)

	return Cache{
		CacheDir: cacheDir,
		Disabled: err != nil,
	}
}

const baseCacheUrl = "cache"

func (c Cache) StartStaticServer(e *echo.Echo) error {
	e.Static("/"+baseCacheUrl, c.CacheDir)
	return nil
}

func hashUrl(url string) string {
	h := sha1.New()
	io.WriteString(h, url)
	return hex.EncodeToString(h.Sum(nil))
}

func (c Cache) CacheUrl(urlString string) (string, error) {
	filename := hashUrl(urlString) + filepath.Ext(urlString)
	targetPath := filepath.Join(c.CacheDir, filename)

	// if the file was already downloaded it doesn't need to be downloaded again
	if _, err := os.Stat(targetPath); errors.Is(err, os.ErrNotExist) {
		resp, err := http.Get(urlString)
		if !errors.Is(err, os.ErrNotExist) {
			return urlString, err
		}
		defer resp.Body.Close()

		file, err := os.Create(targetPath)
		if err != nil {
			return urlString, err
		}
		defer file.Close()

		_, err = io.Copy(file, resp.Body)

		if err != nil {
			return urlString, err
		}
	} else {
		return url.JoinPath(baseCacheUrl, filename)
	}

	return url.JoinPath(baseCacheUrl, filename)
}
