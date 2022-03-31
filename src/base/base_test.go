package base

import (
	"fmt"
	"testing"
	"time"
)

func Test_timeTicker(t *testing.T) {
	t.Run("#time.Ticker test case", func(t *testing.T) {
		tc := time.NewTicker(time.Second) // second ticker => to execute every second task
		for v := range tc.C {
			fmt.Println(v)
		}
		time.Sleep(time.Minute)
	})

	t.Run("#time.Ticker test case", func(t *testing.T) {
		tc := time.NewTicker(time.Second) // second ticker => to execute every second task
		for range tc.C {
			fmt.Println(time.Now().Format(time.RFC3339))
		}
		time.Sleep(time.Minute)
	})
}
