/*
This file contains all types in commands package
*/

package commands

import (
	"inMem/memory"
)

type CommandStruct struct {
	Prefix      string
	Description string
	Function    func(c []string, dir *string, fs *memory.FileSystem)
}

type CommandInputs struct {
	TokenizedCommands []string
	Directory         *string
	MemFS             *memory.FileSystem
}

type Commands map[string]CommandStruct

// process struct (used as a registry entry)
type CommandProcess struct {
	// name of the process
	ProcessName string

	// command info
	Command CommandStruct

	// function used to kill the process
	KillFunc func()

	// killed state
	Killed bool

	// Created time (unix)
	Created int64

	// End time
	End int64

	Deleted bool
}

// process store (used to store Processes)
type CommandProcessStore []CommandProcess
