package storage

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type DiskStorage struct {
	uploadDir string
}

func NewDiskStorage(uploadDir string) *DiskStorage {
	return &DiskStorage{
		uploadDir: uploadDir,
	}
}

func (s *DiskStorage) ConstructPath(ctx context.Context, folderName, filename string) string {
	return fmt.Sprintf("%s/%s/%s", s.uploadDir, folderName, filename)
}

func (s *DiskStorage) Save(ctx context.Context, path string, content io.Reader) error {
	if err := os.MkdirAll(filepath.Dir(path), 0777); err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := io.Copy(file, content); err != nil {
		return err
	}

	return nil
}

func (s *DiskStorage) Delete(ctx context.Context, path string) error {
	// TODO: Implement method
	return errors.New("unimplemented method")
}
