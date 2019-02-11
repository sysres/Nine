package fs

import (
	"os"
	"strings"
	"sync"
	"time"
)

type (
	Dir interface {
		Len() int
		Names() []string
		Files() []*FileData
		Add(*FileData)
		Remove(*FileData)
	}

	FileData struct {
		name    string
		data    []byte
		memdir  *Mem
		dir     bool
		mode    os.FileMode
		modtime time.Time
	}

	Mem struct {
		mu   sync.RWMutex
		data map[string]*FileData
	}

	// Walkfn is a callback signature to walk the files in the memory fs
	Walkfn func(f *FileData) error
)

func isValidPath(name string) bool {
	return name != "" && !strings.Contains(name, "/")
}

// NewInMem creates a new in-memory fs
func NewInMem() *Mem {
	return &Mem{
		data: make(map[string]*FileData),
	}
}

// Mkfile makes a new file in memory
func (m *Mem) Mkfile(name string, data []byte) error {
	if !isValidPath(name) {
		return ErrInvalidFilename
	}

	f := &FileData{
		name: name,
		data: data,
	}

	m.mu.Lock()
	m.data[name] = f
	m.mu.Unlock()
	return nil
}

// Mkdir makes a new directory
func (m *Mem) Mkdir(name string) error {
	if !isValidPath(name) {
		return ErrInvalidFilename
	}

	d := &FileData{
		name:   name,
		memdir: NewInMem(),
		dir:    true,
	}

	m.mu.Lock()
	m.data[name] = d
	m.mu.Unlock()
	return nil
}

// Walk the files in this fs level
func (m *Mem) Walk(fn Walkfn) error {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, fdata := range m.data {
		err := fn(fdata)
		if err != nil {
			return err
		}
	}
	return nil
}

// Data returns the file data specified by name
func (m *Mem) Data(name string) (*FileData, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	fdata, ok := m.data[name]
	if !ok {
		return nil, ErrFileNotFound
	}

	return fdata, nil
}

func (f *FileData) Name() string { return f.name }
func (f *FileData) Data() []byte { return f.data }
func (f *FileData) Dir() *Mem    { return f.memdir }
