package main

import (
	"fmt"
	"os"
	"time"

	mocking "github.com/mathiasb/learngowithtests/09.mocking"
)

func main() {
	sleeper := mocking.NewConfigurableSleeper(1 * time.Second, time.Sleep) // Use the constructor
	mocking.Countdown(os.Stdout, sleeper)
	fmt.Println("")
}
