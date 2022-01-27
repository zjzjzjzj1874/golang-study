package base

import (
	"fmt"
	"testing"
	"time"
)

func TestWarpError_Error(t *testing.T) {
	t.Run("#warp error test case:", func(t *testing.T) {
		if err := genError(); err != nil {
			fmt.Printf("%s\n", err.Error())
		}
	})
}

func genError() error {
	return WarpError{
		When: time.Now(),
		What: "This is an error test",
	}
}
