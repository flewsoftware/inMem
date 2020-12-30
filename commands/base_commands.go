/*
This file contains all base functions used by commands
*/

package commands

import (
	"fmt"
	"github.com/inancgumus/screen"
	"inMem/memory"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func HttpGetToMem(memfs *memory.FileSystem, url string, fileName string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	_, err = memfs.CreateFile(fileName)
	if err != nil {
		return err
	}

	_, err = memfs.WriteFS(fileName, body)
	if err != nil {
		return err
	}

	return nil
}

func HostData(memfs *memory.FileSystem, location string, port int, pattern string) error {
	f, err := memfs.ReadFS(location)
	if err != nil {
		return err
	}

	b, err := ioutil.ReadAll(f.Reader())
	if err != nil {
		return err
	}

	KeepHosting := true
	p := len(CommandProcesses)
	t := time.Now()
	CommandProcesses = append(CommandProcesses, CommandProcess{
		ProcessName: "Hosting " + location + " in port " + strconv.Itoa(port),
		Command:     GetCommands()["host"],
		KillFunc: func() {
			KeepHosting = false
			CommandProcesses[p].Killed = true
		},
		Killed:  false,
		Created: t.Unix(),
		End:     0,
	})

	defer func() {
		CommandProcesses[p].setKilled()
	}()

	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if KeepHosting {
			_, err = w.Write(b)
			if err != nil {
				log.Fatalln(err)
			}
		}
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

	return nil
}

func StashSession(memfs *memory.FileSystem, id string) {
	sessionStore[id] = memory.FileSystemStoreEntry{
		FS:     *memfs,
		Stored: time.Now().Unix(),
	}
	memfs.ClearFS()
}

func CollectSession(memfs *memory.FileSystem, id string, stashCurrent bool, newId string) {
	s := sessionStore[id]
	fs := s.FS.MFileSystem
	if stashCurrent {
		sessionStore[newId] = memory.FileSystemStoreEntry{
			FS:     *memfs,
			Stored: time.Now().Unix(),
		}
	}
	memfs.ReplaceFS(fs)
}

func ClearScreen() {
	screen.Clear()
	screen.MoveTopLeft()
}
