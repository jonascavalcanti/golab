package sorter

// recusion(x int): Inform a int number
func Recusion(x int) int {
	if x <= 1 {
		return x
	} else {
		return Recusion(x-1) * 2
	}
}
