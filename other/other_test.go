package other

import (
	"testing"
	"time"

	"github.com/kasari/test-parallel-test/share"
)

func TestOther(t *testing.T) {
	t.Parallel()

	share.Str = t.Name()

	for i := 0; i < 3; i++ {
		time.Sleep(300 * time.Millisecond)
		t.Log(i, share.Str)
	}
}
