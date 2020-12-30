This package contains commands and their info,   

## Creating custom commands
1\. Create a base function in the base_commands.go file   
example:
```go
func SayLol() {
    println("lol")
}
```

2\. Use the base function and create a command in the commands.go file (The command function should be CommandStruct.Function) 
```go
func SayLolCommand(c []string, dir *string, fs *memory.FileSystem){
    SayLol()
}
```

3\.  Add the function command to main.go 
```go
func GetCommands() Commands {
    var m = make(map[string]CommandStruct)
    // .......
    // .......

    // add your custom commands here
    m["saylol"] = CommandStruct{
        prefix: "saylol",
		Description: "outputs lol to stdout,
		Function:    SayLolCommand,
    }
}
```

