/*
This file contains global vars used by commands package
*/

package commands

import "inMem/memory"

var processes []Process

var sessionStore = make(memory.FileSystemStore)
