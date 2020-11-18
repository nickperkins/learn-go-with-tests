package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Sleeper interface to allow dependency injection
type Sleeper interface {
	Sleep()
}

// DefaultSleeper for when we want to run this thing
type DefaultSleeper struct{}

// Sleep calls time.Sleep
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// Countdown from 3 to 1, print Go! and sleep 1 second between each print
func Countdown(writer io.Writer, sleeper Sleeper) {
	const finalWord = "Go!"
	const countdownStart = 3
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(writer, i)
	}

	sleeper.Sleep()
	fmt.Fprint(writer, finalWord)

}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
