package memory

import (
	"github.com/blang/vfs/memfs"
)

// creates a in memory file system
func CreateMemoryFileSystem() FileSystem {
	return FileSystem{MFileSystem: memfs.Create()}
}
