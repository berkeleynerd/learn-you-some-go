package main

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
    Sleep()
}

type DefaultSleeper struct {}

type ConfigurableSleeper struct {
    duration time.Duration
    sleep    func(time.Duration)
}

func (d *DefaultSleeper) Sleep() {
    time.Sleep(1 * time.Second)
}

func (c *ConfigurableSleeper) Sleep() {
    c.sleep(c.duration)
}

const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer, sleeper Sleeper) {
    for i := countdownStart; i > 0; i-- {
        sleeper.Sleep()
        fmt.Fprintln(out, i)
    }

    sleeper.Sleep()
    fmt.Fprint(out, finalWord)
}
