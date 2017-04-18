package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	xCode := sha256.Sum256([]byte("x"))
	XCode := sha256.Sum256([]byte("X"))
	fmt.Printf("x:%x\nX:%x\n%t\n%T\n", xCode, XCode, xCode == XCode, xCode)
}
