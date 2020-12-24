/*
This file contains all types in commands package
*/

package commands

import "inMem/memory"

type CommandStruct struct {
	Prefix      string
	Description string
}

// process struct (used as a registry entry)
type Process struct {
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
}

type CommandInputs struct {
	TokenizedCommands []string
	Directory         *string
	MemFS             *memory.FileSystem
}

type Commands map[string]CommandStruct
