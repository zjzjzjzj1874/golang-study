package errgroup

import "testing"

func TestErrGroup(t *testing.T) {
	t.Run("err group", func(t *testing.T) {
		errGroup()
	})

	t.Run("err group with ctx", func(t *testing.T) {
		cancelGoroutineInErrWithCtx()
	})
}
