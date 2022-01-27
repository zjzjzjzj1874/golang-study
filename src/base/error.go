package base

import (
	"fmt"
	"time"
)

// WarpError warp error
type WarpError struct {
	When time.Time // error happened time
	What string    // error's content
}

func (w WarpError) Error() string {
	return fmt.Sprintf("%v:%s", w.When.Format(time.RFC3339Nano), w.What)
}
