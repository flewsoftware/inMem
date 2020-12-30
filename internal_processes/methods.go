package internal_processes

import (
	"inMem/memory"
	"io/ioutil"
	"math/rand"
	"time"
)

func (IPQ *InternalProcessQueue) GetProcessByType(pType ProcessType) []Process {
	cp := *IPQ
	var processes []Process

	for i := 0; i < len(cp); i++ {
		if cp[i].ProcessType == pType {
			processes = append(processes, cp[i])
		}
	}
	return processes
}

// adds a process to the queue and returns its queue id
func (IPQ *InternalProcessQueue) AddProcessToQueue(IProcess Process) string {
	rand.Seed(time.Now().UnixNano())
	IProcess.id = string(rand.Int())

	*IPQ = append(*IPQ, IProcess)
	return IProcess.id
}

// removes process by id
func (IPQ *InternalProcessQueue) RemoveProcess(id string) int {

	var removedProcessCount int

	cIPQ := *IPQ
	for i := 0; i < len(cIPQ); i++ {
		if cIPQ[i].id == id {
			cIPQ = removeProcessItem(cIPQ, i)
			removedProcessCount++
		}
	}
	return removedProcessCount
}

func (IPQ *InternalProcessQueue) RemoveItem(t []InternalProcessQueue, item int) []InternalProcessQueue {
	lo := item
	t = append(t[:lo], t[lo+1:]...)

	return t
}

func (p *Process) Run() string {
	p.ProcessFunc(p.id)
	return p.id
}
func (std *InternalStd) Println(data string) error {
	cstd := *std
	data = "\n" + data
	_, err := cstd.fs.WriteFS(cstd.location, []byte(data))
	if err != nil {
		return err
	}
	return nil
}

func (std *InternalStd) InitInternalStd(fs *memory.FileSystem) error {
	stdoutLocation := "stdout"
	_, err := fs.CreateFile(stdoutLocation)
	if err != nil {
		return err
	}
	_, err = fs.WriteFS(stdoutLocation, nil)
	if err != nil {
		return err
	}
	*std = InternalStd{
		location: stdoutLocation,
		fs:       fs,
	}
	return nil
}

func (std *InternalStd) PrintToStdOut() error {
	fs := std.fs
	lb, err := fs.ReadFS(std.location)
	if err != nil {
		return err
	}
	defer lb.Destroy()

	out, err := ioutil.ReadAll(lb.Reader())
	if err != nil {
		return err
	}
	println(string(out))
	return nil
}
