package parallel

import (
	"testing"
	"time"

	"github.com/kasari/test-parallel-test/share"
)

func Test2(t *testing.T) {
	t.Parallel()

	share.Str = t.Name()

	for i := 0; i < 3; i++ {
		time.Sleep(200 * time.Millisecond)
		t.Log(i, share.Str)
	}
}
