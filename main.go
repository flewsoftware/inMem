package main

import (
	"bufio"
	"fmt"
	"inMem/commands"
	"inMem/memory"
	"os"
	"strings"
)

func main() {
	var directory = "/"
	commandList := commands.GetCommands()
	memFS := memory.CreateMemoryFileSystem()

	for {
		fmt.Printf("inMem %s >", directory)
		var command string

		reader := bufio.NewReader(os.Stdin)
		command, _ = reader.ReadString('\n')
		tokenizedCommands := commandTokenizer(command)

		commandPicker(command, commandList, commands.CommandInputs{TokenizedCommands: tokenizedCommands, Directory: &directory, MemFS: &memFS})

	}

}

// executes the correct command depending on the user input
func commandPicker(command string, commandList commands.Commands, inputs commands.CommandInputs) {
	memfs := inputs.MemFS
	directory := inputs.Directory
	tokenizedCommands := inputs.TokenizedCommands

	if commandMatch(command, commandList["download"]) {
		commands.DownloadCommand(tokenizedCommands, directory, memfs)
	} else if commandMatch(command, commandList["host"]) {
		go commands.HostCommand(tokenizedCommands, directory, memfs)
	} else if commandMatch(command, commandList["exit"]) {
		commands.ExitCommand(tokenizedCommands, directory, memfs)
	} else if commandMatch(command, commandList["cd"]) {
		commands.CdCommand(tokenizedCommands, directory, memfs)
	} else if commandMatch(command, commandList["mkdir"]) {
		commands.MkDirCommand(tokenizedCommands, directory, memfs)
	} else if commandMatch(command, commandList["ls"]) {
		commands.LsCommand(tokenizedCommands, directory, memfs)
	} else if commandMatch(command, commandList["kill"]) {
		commands.KillCommand(tokenizedCommands, directory, memfs)
	} else if commandMatch(command, commandList["lp"]) {
		commands.ListProcesses(tokenizedCommands, directory, memfs)
	} else if commandMatch(command, commandList["ns"]) {
		commands.NewSessionCommand(tokenizedCommands, directory, memfs)
	} else if commandMatch(command, commandList["fh"]) {
		commands.FHcommand(tokenizedCommands, directory, memfs)
	} else if commandMatch(command, commandList["help"]) {
		commands.HelpCommand(tokenizedCommands, directory, memfs)
	} else if commandMatch(command, commandList["mkfile"]) {
		commands.CreateFileCommand(tokenizedCommands, directory, memfs)
	} else if commandMatch(command, commandList["csession"]) {
		commands.CollectSessionCommand(tokenizedCommands, directory, memfs)
	} else if commandMatch(command, commandList["ssession"]) {
		commands.StashSessionCommand(tokenizedCommands, directory, memfs)
	} else if commandMatch(command, commandList["lssessions"]) {
		commands.LsSessionsCommand(tokenizedCommands, directory, memfs)
	} else {
		fmt.Printf("Invalid command: %s", command)
	}
	print("\n")

}

// returns if inputString matches the command prefix
func commandMatch(inputString string, command commands.CommandStruct) bool {
	tokens := commandTokenizer(inputString)
	if tokens[0] == command.Prefix {
		return true
	}
	return false
}

// tokenize inputString
func commandTokenizer(inputString string) []string {
	return sliceLowerCase(strings.Split(inputString, " "))
}

func sliceLowerCase(slice2 []string) []string {
	for i := 0; i < len(slice2); i++ {
		slice2[i] = strings.ReplaceAll(strings.ReplaceAll(strings.ToLower(slice2[i]), "\n", ""), "\r", "")
	}
	return slice2
}
