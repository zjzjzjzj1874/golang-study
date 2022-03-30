package _select

import (
	"errors"
	"math/rand"
	"sync"
)

type task struct{}

func (t task) Run() error {
	if rand.Intn(5)%2 == 0 {
		return errors.New("new error")
	}
	return nil
}

func nonBlock() error {
	// 这个demo只关心有没有错误,不关心有多少个任务执行失败
	tasks := make([]task, 5)
	errCh := make(chan error, len(tasks))
	wg := sync.WaitGroup{}
	wg.Add(len(tasks))
	for i := range tasks {
		go func() {
			defer wg.Done()
			if err := tasks[i].Run(); err != nil {
				errCh <- err
			}
		}()
	}
	wg.Wait()

	select {
	case err := <-errCh:
		return err
	default:
		return nil
	}
}
