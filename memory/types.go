package memory

import "github.com/blang/vfs/memfs"

// file system struct
type FileSystem struct {
	MFileSystem *memfs.MemFS
}

// file system store
type FileSystemStore map[string]FileSystemStoreEntry

// file system store entry
type FileSystemStoreEntry struct {

	// file system
	FS FileSystem

	// the time when the file system was stashed
	Stored int64
}
