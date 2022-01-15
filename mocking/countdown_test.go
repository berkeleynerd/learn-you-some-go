package main

import (
	"bytes"
    "reflect"
	"testing"
    "time"
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

type ObservableTime struct {
    durationSlept time.Duration
}

func (s *ObservableTime) Sleep(duration time.Duration) {
    s.durationSlept = duration
}

func TestConfigurableSleeper(t *testing.T) {
    sleepTime := 5 * time.Second

    observableTime := &ObservableTime{}
    sleeper := ConfigurableSleeper{sleepTime, observableTime.Sleep}
    sleeper.Sleep()

    if observableTime.durationSlept != sleepTime {
        t.Errorf("should have slept for %v but slept for %v", sleepTime, observableTime.durationSlept)
    }
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
