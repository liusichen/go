package countdiff

import "crypto/sha256"

//CountDiff is count the diff between two hash code
func CountDiff(firstchar string, secondchar string) int {
	c1 := sha256.Sum256([]byte(firstchar))
	c2 := sha256.Sum256([]byte(secondchar))
	for i := range c1 {
		c1[i] = c1[i] ^ c2[i]
	}
	return countbitByte(c1)
}
func countBitint64(value int) int {
	var count int
	for value != 0 {
		value = value & (value - 1)
		count++
	}
	return count
}

func countbitByte(input [32]byte) int {
	var upCount int
	for i := range input {
		upCount += countBitint64(int(input[i]))
	}
	return upCount
}
