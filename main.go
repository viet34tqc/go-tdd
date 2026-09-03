package main

import (
	"os"
	"test/mocking"
	"time"
)

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	sleeper := &DefaultSleeper{}
	mocking.Countdown(os.Stdout, sleeper)
}
