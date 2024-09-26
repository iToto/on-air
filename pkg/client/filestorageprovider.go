package client

import (
	"context"
	"io"
)

//go:generate mockgen --build_flags=-mod=mod -package=client -destination=mock_filestorageprovider.go . FileStorageProvider

// FileStorageProvider represents a provider for storing and retrieving documents
type FileStorageProvider interface {
	Save(ctx context.Context, path string, content io.Reader) error
	Load(ctx context.Context, path string, content io.Writer) error
	Close() error
}
