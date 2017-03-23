package parallel

import (
	"testing"
	"time"

	"github.com/kasari/test-parallel-test/share"
)

func Test1(t *testing.T) {
	t.Parallel()

	share.Str = t.Name()

	for i := 0; i < 3; i++ {
		time.Sleep(300 * time.Millisecond)
		t.Log(i, share.Str)
	}
}

func Test1_2(t *testing.T) {
	share.Str = t.Name()

	for i := 0; i < 3; i++ {
		time.Sleep(300 * time.Millisecond)
		t.Log(i, share.Str)
	}
}
