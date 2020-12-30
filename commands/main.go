package commands

// returns a list of commandStruct
func GetCommands() Commands {
	var m = make(map[string]CommandStruct)

	// commands
	m["download"] = CommandStruct{
		Prefix:      "download",
		Description: "downloads the given url",
		Function:    DownloadCommand,
	}

	m["host"] = CommandStruct{
		Prefix:      "host",
		Description: "hosts the given file",
		Function:    HostCommand,
	}

	m["exit"] = CommandStruct{
		Prefix:      "exit",
		Description: "exits inMem",
		Function:    ExitCommand,
	}

	m["cd"] = CommandStruct{
		Prefix:      "cd",
		Description: "change directories",
		Function:    ChangeDirCommand,
	}

	m["mkdir"] = CommandStruct{
		Prefix:      "mkdir",
		Description: "make directories",
		Function:    MakeDirCommand,
	}

	m["ls"] = CommandStruct{
		Prefix:      "ls",
		Description: "list directories",
		Function:    ListCommand,
	}

	m["kill"] = CommandStruct{
		Prefix:      "kill",
		Description: "kills CommandProcesses",
		Function:    KillCommand,
	}

	m["lp"] = CommandStruct{
		Prefix:      "lp",
		Description: "Lists CommandProcesses",
		Function:    ListProcessesCommand,
	}

	m["nsession"] = CommandStruct{
		Prefix:      "nsession",
		Description: "New session",
		Function:    NewSessionCommand,
	}

	m["fh"] = CommandStruct{
		Prefix:      "fh",
		Description: "file hash",
		Function:    FileHashCommand,
	}

	m["help"] = CommandStruct{
		Prefix:      "help",
		Description: "displays commands",
		Function:    HelpCommand,
	}

	m["mkfile"] = CommandStruct{
		Prefix:      "mkfile",
		Description: "creates a file",
		Function:    MakeFileCommand,
	}

	m["ssession"] = CommandStruct{
		Prefix:      "ssession",
		Description: "stashes current session",
		Function:    StashSessionCommand,
	}

	m["csession"] = CommandStruct{
		Prefix:      "csession",
		Description: "collect session and drop current session",
		Function:    CollectSessionCommand,
	}

	m["lssessions"] = CommandStruct{
		Prefix:      "lssessions",
		Description: "lists sessions in store",
		Function:    ListSessionsCommand,
	}

	m["cpl"] = CommandStruct{
		Prefix:      "cpl",
		Description: "clean process list",
		Function:    CleanProcessList,
	}

	m["pout"] = CommandStruct{
		Prefix:      "pout",
		Description: "outputs process stdout",
		Function:    ProcessOutCommand,
	}

	m["clear"] = CommandStruct{
		Prefix:      "clear",
		Description: "clears the screen",
		Function:    ClearCommand,
	}
	return m
}
