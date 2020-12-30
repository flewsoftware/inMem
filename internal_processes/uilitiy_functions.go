package internal_processes

func removeProcessItem(t []Process, item int) []Process {
	lo := item
	t = append(t[:lo], t[lo+1:]...)

	return t
}
