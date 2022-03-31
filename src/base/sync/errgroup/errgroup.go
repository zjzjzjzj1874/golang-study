package errgroup

import (
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
