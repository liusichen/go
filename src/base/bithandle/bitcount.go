package bithandle

func countBitint64(value int) int {
	var count int
	for value != 0 {
		value = value & (value - 1)
		count++
	}
	return count
}
