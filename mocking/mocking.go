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

// A ConfigurableSleeper is a Sleeper which can be configured as we go
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
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
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
