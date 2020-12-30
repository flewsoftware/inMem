package commands

func removeCommandProcessItem(c []CommandProcess, item int) []CommandProcess {
	lo := item
	c = append(c[:lo], c[lo+1:]...)

	return c
}
