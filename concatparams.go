package piscine

func ConcatParams(args []string) string {
	rep := args[0]
	for _, ch := range args[1:] {
		rep += "\n" + ch
	}
	return rep
}
