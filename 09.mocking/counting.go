package mocking

import (
	"fmt"
	"io"
	"time"
)

const (
	countdownStart = 3
	finalCallout   = "Go!"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (ds *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// In your mocking package (09.mocking/counting.go)
func NewConfigurableSleeper(duration time.Duration, sleep func(time.Duration)) *ConfigurableSleeper {
	return &ConfigurableSleeper{
		duration: duration,
		sleep:    sleep,
	}
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}
	fmt.Fprint(writer, finalCallout)
}
