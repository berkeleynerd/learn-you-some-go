package main

import "os"

func main() {
    sleeper := &DefaultSleeper{}
    Countdown(os.Stdout, sleeper)
}
