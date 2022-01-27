package _select

import (
	"math/rand"
	"testing"
	"time"
)

func TestDemoSelect_waitForHeartbeat(t *testing.T) {
	t.Run("#select for timeout", func(t *testing.T) {
		demo := DemoSelect{
			signal: make(chan struct{}, 1),
		}

		go demo.waitForHeartbeat()

		go func() {
			for i := 1; i <= 10; i++ {
				time.Sleep(time.Duration(rand.Intn(i)) * time.Second * 2)
				demo.signal <- struct{}{}
			}
		}()

		time.Sleep(time.Minute)
	})
}
