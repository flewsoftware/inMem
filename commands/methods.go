/*
This file contains methods in commands package
*/

package commands

import (
	"inMem/internal_processes"
	"log"
	"time"
)

// clears killed CommandProcesses
func (p *CommandProcessStore) Clear() {
	pCopy := *p
	for key, val := range pCopy {
		if val.Killed {
			pCopy[key].Deleted = true

			clearProcess := internal_processes.Process{
				ProcessType: internal_processes.CleanUpProcess,
				ProcessFunc: func(id string) {

					removeCommandProcessItem(pCopy, key)

					err := internal_processes.InterSTD.Println("removed command process with id: " + id + "from command process store")
					if err != nil {
						log.Fatalf("cant save process output: %v", err)
					}
				},
				Created: time.Now(),
			}
			internal_processes.InternalProcesses.AddProcessToQueue(clearProcess)
		}
	}
}

// set process has been killed without calling KillFunc
func (p *CommandProcess) setKilled() {
	p.Killed = true
	p.End = time.Now().Unix()
}
