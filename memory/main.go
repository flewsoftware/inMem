package memory

import (
	"crypto/sha256"
	"github.com/awnumar/memguard"
	"github.com/blang/vfs/memfs"
	"github.com/diskfs/go-diskfs/filesystem"
	"io"
	"os"
)

// reads a file in the memory file system
func (fs *FileSystem) ReadFS(file string) (*memguard.LockedBuffer, error) {
	f, err := fs.MFileSystem.OpenFile(file, os.O_RDONLY, os.ModeAppend)
	if err != nil {
		return nil, err
	}
	lf, err := memguard.NewBufferFromEntireReader(f)
	if err != nil {
		return nil, err
	}

	return lf, nil
}

// creates a file in the memory file system
func (fs *FileSystem) CreateFile(file string) (filesystem.File, error) {
	return fs.MFileSystem.OpenFile(file, os.O_CREATE, os.ModeAppend)
}

// writes a file in the memory file system
func (fs *FileSystem) WriteFS(file string, data []byte) (int, error) {
	f, err := fs.MFileSystem.OpenFile(file, os.O_RDWR, os.ModeAppend)
	if err != nil {
		return 0, err
	}
	n, err := f.Write(data)
	if err != nil {
		return 0, err
	}
	return n, nil
}

// returns a sha-256 hash of the file
func (fs *FileSystem) GetHash(file string) ([]byte, error) {
	h := sha256.New()
	f, err := fs.ReadFS(file)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(h, f.Reader())
	if err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

// replaces memfs
func (fs *FileSystem) ReplaceFS(f *memfs.MemFS) {
	fs.MFileSystem = f
}

// replaces memfs with a new one
func (fs *FileSystem) ClearFS() {
	fs.MFileSystem = CreateMemoryFileSystem().MFileSystem
}

// creates a in memory file system
func CreateMemoryFileSystem() FileSystem {
	return FileSystem{MFileSystem: memfs.Create()}
}
