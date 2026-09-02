package main

import (
	"os"
	"test/mocking"
)

func main() {
	mocking.Countdown(os.Stdout)
}
