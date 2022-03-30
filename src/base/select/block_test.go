package _select

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestDemoSelect(t *testing.T) {
	t.Run("直接阻塞", func(t *testing.T) {
		go func() {
			for {
				select {}
				fmt.Println("hello")
			}
		}()
		stopCh := make(chan os.Signal, 1)
		signal.Notify(stopCh, os.Interrupt, syscall.SIGTERM)
		<-stopCh
		fmt.Println("ready to exit")
		time.Sleep(time.Second * 5)
		fmt.Println("bye")
	})

	t.Run("单一管道", func(t *testing.T) {
		ch := make(chan struct{})

		go func() {
			for {
				select {
				case <-ch:
					fmt.Println("get signal")
				}
			}
		}()
		ch <- struct{}{}

		stopCh := make(chan os.Signal, 1)
		signal.Notify(stopCh, os.Interrupt, syscall.SIGTERM)
		<-stopCh
	})

	t.Run("非阻塞操作", func(t *testing.T) {
		ch := make(chan struct{})
		exitSignal := make(chan struct{})

		go func() {
			for {
				select {
				case <-ch:
					fmt.Println("get signal")
				case <-exitSignal:
					fmt.Println("get exit signal")
					return
				default:
					fmt.Println("no signal")
					time.Sleep(time.Second)
				}
			}
		}()
		time.Sleep(time.Second * time.Duration(rand.Intn(5)))
		ch <- struct{}{}
		time.Sleep(time.Second * time.Duration(rand.Intn(5)))
		exitSignal <- struct{}{}
	})

	t.Run("多个通道", func(t *testing.T) {
		// select中,多个case都满足的话,会随机选择一个执行;如果按顺序执行,可能会有case会被饿死
		ch := make(chan struct{})
		exitSignal := make(chan struct{})

		go func() {
			for {
				select {
				case <-ch:
					fmt.Println("get signal 1")
				case <-ch:
					fmt.Println("get signal 2")
				case <-exitSignal:
					fmt.Println("get exit signal")
					return
				default:
					fmt.Println("no signal")
					time.Sleep(time.Second)
				}
			}
		}()
		time.Sleep(time.Second * time.Duration(rand.Intn(5)))
		for i := 0; i < 5; i++ {
			ch <- struct{}{}
		}
		time.Sleep(time.Second * time.Duration(rand.Intn(5)))
		exitSignal <- struct{}{}
	})

}
