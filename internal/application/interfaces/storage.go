package interfaces

import (
	"context"
	"io"
)

type Storage interface {
	ConstructPath(ctx context.Context, folderName, filename string) string
	Save(ctx context.Context, path string, content io.Reader) error
	Delete(ctx context.Context, path string) error
}
