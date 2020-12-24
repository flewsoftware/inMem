package memory

import "github.com/blang/vfs/memfs"

// file system struct
type FileSystem struct {
	MFileSystem *memfs.MemFS
}

// file system store
type FileSystemStore map[string]FileSystemStoreEntry

// file system entry
type FileSystemStoreEntry struct {
	FS     FileSystem
	Stored int64
}
