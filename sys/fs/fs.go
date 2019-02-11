package fs

import (
	"errors"
	"io"
	"os"
)

type (
	// File represents a file in the filesystem.
	File interface {
		io.Closer
		io.Reader
		io.ReaderAt
		io.Seeker
		io.Writer
		io.WriterAt

		Name() string
		Readdir(count int) ([]os.FileInfo, error)
		Readdirnames(n int) ([]string, error)
		Stat() (os.FileInfo, error)
		Sync() error
		Truncate(size int64) error
		WriteString(s string) (ret int, err error)
	}

	// FS common interface
	FS interface {
		Create(name string) (File, error)
		Open(name string) (File, error)
	}
)

var (
	ErrInvalidFilename   = errors.New("Invalid file name")
	ErrFileClosed        = errors.New("File is closed")
	ErrOutOfRange        = errors.New("Out of range")
	ErrTooLarge          = errors.New("Too large")
	ErrFileNotFound      = os.ErrNotExist
	ErrFileExists        = os.ErrExist
	ErrDestinationExists = os.ErrExist
)
