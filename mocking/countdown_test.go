package main

import (
	"bytes"
    "reflect"
	"testing"
)

const write = "write"
const sleep = "sleep"

type ObserverCountdownOperations struct {
    Calls []string
}

func (s *ObserverCountdownOperations) Sleep() {
    s.Calls = append(s.Calls, sleep)
}

func (s *ObserverCountdownOperations) Write(p []byte) (n int, err error) {
    s.Calls = append(s.Calls, write)
    return
}

func TestCountdown(t *testing.T) {

    t.Run("prints 3 to Go!", func(t *testing.T) {
        buffer := &bytes.Buffer{}
        Countdown(buffer, &ObserverCountdownOperations{})

        got := buffer.String()
        want := `3
2
1
Go!`

        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    })

    t.Run("sleep before every print", func(t *testing.T) {
        observerSleepPrinter := &ObserverCountdownOperations{}
        Countdown(observerSleepPrinter, observerSleepPrinter)

        want := []string{
            sleep,
            write,
            sleep,
            write,
            sleep,
            write,
            sleep,
            write,
        }

        if !reflect.DeepEqual(want, observerSleepPrinter.Calls) {
            t.Errorf("wanted calls %v got %v", want, observerSleepPrinter.Calls)
        }
    })

}
