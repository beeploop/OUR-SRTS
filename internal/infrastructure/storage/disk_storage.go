package storage

import (
	"context"
	"io"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
)

type DiskStorage struct {
	tempDir   string
	uploadDir string
	dirPerms  fs.FileMode
}

func NewDiskStorage(uploadDir string) *DiskStorage {
	return &DiskStorage{
		tempDir:   "temp",
		uploadDir: uploadDir,
		dirPerms:  0777,
	}
}

func (s *DiskStorage) ConstructPath(ctx context.Context, folderName, filename string) string {
	return filepath.Join(s.uploadDir, folderName, filename)
}

func (s *DiskStorage) Save(ctx context.Context, path string, content io.Reader) error {
	if err := os.MkdirAll(filepath.Dir(path), s.dirPerms); err != nil {
		return err
	}

	// Save to temp location
	tempPath, err := s.saveTemp(filepath.Base(path), content)
	if err != nil {
		return err
	}

	// Copy to upload location
	tempFile, err := os.Open(tempPath)
	if err != nil {
		return err
	}
	defer tempFile.Close()

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := io.Copy(file, tempFile); err != nil {
		return err
	}

	return s.Delete(ctx, tempPath)
}

func (s *DiskStorage) Delete(ctx context.Context, path string) error {
	if err := os.Remove(path); err != nil {
		return err
	}

	slog.Info("File Deleted", "path", path)
	return nil
}

func (s *DiskStorage) saveTemp(filename string, content io.Reader) (string, error) {
	if err := os.MkdirAll(s.tempDir, s.dirPerms); err != nil {
		return "", err
	}

	path := filepath.Join(s.tempDir, filename)
	b, err := io.ReadAll(content)
	if err != nil {
		return "", err
	}

	if err := os.WriteFile(path, b, s.dirPerms); err != nil {
		return "", err
	}

	slog.Info("File Saved To Temp", "filename", filename, "path", path)
	return path, nil
}
