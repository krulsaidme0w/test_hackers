package rps

import (
	"fmt"
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

const maxGoroutines = 10000

func TestRPS() {
	rate := vegeta.Rate{Freq: 10000, Per: time.Second}
	duration := 5 * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    "http://localhost:8010/json/hackers",
	})
	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Test rps") {
		metrics.Add(res)
	}
	metrics.Close()

	fmt.Println("RATE: ", metrics.Rate)
}
