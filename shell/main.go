package shell

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"inMem/commands"
	"inMem/internal_processes"
	"inMem/memory"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

// starts shell interface and starts accepting input from stdin
func StartShell() {
	// This variable represents the current directory
	var currentDirectory = "/"

	// Gets a list of commands
	commandList := commands.GetCommands()

	// The in memory file system
	memFS := memory.CreateMemoryFileSystem()

	err := internal_processes.InterSTD.InitInternalStd(&memFS)
	if err != nil {
		log.Fatalf("cant initailize internal process stdout")
	}
	interStd = internal_processes.InterSTD

	// checks every 15 seconds if there are any cleaning jobs to do (goroutine)
	CleanUpProcesses()

	// This will capture user input, execute the necessary command and then loop back
	for {
		fmt.Printf("inMem %s >", currentDirectory)

		var command string

		// Captures the user input
		reader := bufio.NewReader(os.Stdin)
		command, _ = reader.ReadString('\n')

		// Tokenize the input
		tokenizedCommands := commandTokenizer(command)

		// Picks the correct command
		c := commandPicker(command, commandList)

		var wg sync.WaitGroup
		wg.Add(1)
		// Runs it

		c(tokenizedCommands, &currentDirectory, &memFS, &wg)

		wg.Wait()

		fmt.Print("\n")
	}

}

var interStd internal_processes.InternalStd

// checks for proceses marked as deleted in the internal process queue and deletes them
func CleanUpProcesses() {
	go func() {
		for {
			p := internal_processes.InternalProcesses.GetProcessByType(internal_processes.CleanUpProcess)
			if len(p) > 0 {
				for i := 0; i < len(p); i++ {
					tempP := p[i]
					id := tempP.Run()
					internal_processes.InternalProcesses.RemoveProcess(id)
					err := interStd.Println("Cleaned")
					if err != nil {
						log.Fatalln(err)
					}
				}
			}

			// wait 15 seconds
			time.Sleep(15 * time.Second)
		}
	}()
}

// returns the matching command
func commandPicker(command string, commandList commands.Commands) func(c []string, dir *string, fs *memory.FileSystem, wg *sync.WaitGroup) {

	for key := range commandList {
		if commandMatch(command, commandList[key]) {
			return commandList[key].Function
		}
	}

	return func(_ []string, _ *string, _ *memory.FileSystem, wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Printf("Invalid command: %s", command)
	}
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

// prints shell info
func PrintShellInfo() {
	d := color.New(color.FgHiGreen, color.Bold)

	_, _ = d.Println("FlewSoftware InMem Shell")

	color.Cyan("Created with ❤ by Tarith Jayasooriya & contributors.")
	color.Cyan("Open sourced under MIT Licence")

	println("Type \"help\" to get a list of commands and their usages")
	println("Go to https://github.com/flew-software/inMem to learn more about inMem")

	print("\n\n")
}
