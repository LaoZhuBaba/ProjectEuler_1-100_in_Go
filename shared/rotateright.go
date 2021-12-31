package shared

func RotateRight(l []int) {
	// fmt.Printf("rotateRight() called with %v\n", l)
	var temp int
	maxIndex := len(l) - 1
	if maxIndex < 1 {
		return
	}
	for c := maxIndex; c > 0; c-- {
		if c == maxIndex {
			temp = l[c]
		}
		l[c] = l[c-1]
	}
	l[0] = temp
}
