package internal_processes

import (
	"inMem/memory"
	"time"
)

type InternalProcessQueue []Process

type Process struct {
	ProcessType ProcessType
	ProcessFunc func(id string)
	Created     time.Time
	id          string
}

type ProcessType int

type InternalStd struct {
	location string
	fs       *memory.FileSystem
}
