package time

import (
	"testing"
	"time"
)

func TestVariable(t *testing.T) {
	time := time.Now()
	t.Logf("%#v", time)
}
