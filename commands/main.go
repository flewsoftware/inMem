package commands

// returns a list of commandStruct
func GetCommands() Commands {
	var m = make(map[string]CommandStruct)

	// commands
	m["download"] = CommandStruct{
		Prefix:      "download",
		Description: "downloads the given url",
	}

	m["host"] = CommandStruct{
		Prefix:      "host",
		Description: "hosts the given file",
	}

	m["exit"] = CommandStruct{
		Prefix:      "exit",
		Description: "exits inMem",
	}

	m["cd"] = CommandStruct{
		Prefix:      "cd",
		Description: "change directories",
	}

	m["mkdir"] = CommandStruct{
		Prefix:      "mkdir",
		Description: "make directories",
	}

	m["ls"] = CommandStruct{
		Prefix:      "ls",
		Description: "list directories",
	}

	m["kill"] = CommandStruct{
		Prefix:      "kill",
		Description: "kills processes",
	}

	m["lp"] = CommandStruct{
		Prefix:      "lp",
		Description: "Lists processes",
	}

	m["nsession"] = CommandStruct{
		Prefix:      "nsession",
		Description: "New session",
	}

	m["fh"] = CommandStruct{
		Prefix:      "fh",
		Description: "session hash",
	}

	m["help"] = CommandStruct{
		Prefix:      "help",
		Description: "displays commands",
	}

	m["mkfile"] = CommandStruct{
		Prefix:      "mkfile",
		Description: "creates a file",
	}

	m["ssession"] = CommandStruct{
		Prefix:      "ssession",
		Description: "stashes current session",
	}

	m["csession"] = CommandStruct{
		Prefix:      "csession",
		Description: "collect session and drop current session",
	}

	m["lssessions"] = CommandStruct{
		Prefix:      "lssessions",
		Description: "lists sessions in store",
	}
	return m
}
