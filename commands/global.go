/*
This file contains global vars used by commands package
*/

package commands

import "inMem/memory"

// stores information about running CommandProcesses
var CommandProcesses CommandProcessStore

// stores stashed sessions(file systems)
var sessionStore = make(memory.FileSystemStore)
