// Package _select use in timeout ==> waitGroup在timeout中使用
package _select

import (
	"fmt"
	"time"
)

type DemoSelect struct {
	signal chan struct{}
}

func (d DemoSelect) waitForHeartbeat() {
	for {
		select {
		case <-d.signal:
			fmt.Println("信号到达")
			// do some business
		case <-time.After(time.Second * 5):
			fmt.Println("5s超时")
			// do some timeout business
		}
	}
}
