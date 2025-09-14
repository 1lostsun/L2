package storage

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

type Storage struct {
	baseDir string
}

func NewStorageOnDesktop(folderName string) *Storage {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "."
	}

	desktopDir := filepath.Join(homeDir, "Desktop", folderName)
	return &Storage{baseDir: desktopDir}
}

func (s *Storage) Save(pageURL string, content []byte) error {
	fileName := urlToFilename(pageURL)
	fullPath := filepath.Join(s.baseDir, fileName)

	if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
		return err
	}

	file, err := os.Create(fullPath)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer func(file *os.File) {
		closeErr := file.Close()
		if closeErr != nil {
			log.Fatal(closeErr)
		}
	}(file)

	if _, writeErr := file.Write(content); writeErr != nil {
		return fmt.Errorf("error writing to file: %w", writeErr)
	}

	return nil
}

func urlToFilename(pageURL string) string {
	u, err := url.Parse(pageURL)
	if err != nil {
		return "page.html"
	}

	var pathParts []string

	if u.Host != "" {
		pathParts = append(pathParts, u.Host)
	}

	if u.Path != "" && u.Path != "/" {
		cleanPath := strings.TrimPrefix(u.Path, "/")
		if cleanPath != "" {
			pathParts = append(pathParts, strings.Split(cleanPath, "/")...)
		}
	}

	if len(pathParts) == 0 || u.Path == "/" || strings.HasSuffix(u.Path, "/") {
		pathParts = append(pathParts, "index.html")
	} else {
		lastPart := pathParts[len(pathParts)-1]
		if !strings.HasSuffix(filepath.Ext(lastPart), ".") {
			pathParts[len(pathParts)-1] = lastPart + ".html"
		}
	}

	return filepath.Join(pathParts...)
}

func (s *Storage) GetBaseDir() string {
	return s.baseDir
}
