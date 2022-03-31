package errgroup

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func errGroup() {
	var g errgroup.Group
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
	}
	for i := range urls {
		url := urls[i]
		g.Go(func() error {
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	} else {
		fmt.Println(err)
	}
	// 如果返回错误 — 这一组 Goroutine 最少返回一个错误；
	// 如果返回空值 — 所有 Goroutine 都成功执行；
}

// 这个demo:当有一个goroutine出错,马上通知其他goroutine取消
func cancelGoroutineInErrWithCtx() {
	eg, _ := errgroup.WithContext(context.Background())
	var urls = []string{
		"https://www.golang.org/",
		"https://www.google.com/",
		"https://www.baidu.com/",
	}
	for i := range urls {
		url := urls[i]
		eg.Go(func() error {
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
				fmt.Println(resp)
			}
			fmt.Println(url)
			return err
		})
	}
	if err := eg.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	} else {
		fmt.Println(err)
	}
}
