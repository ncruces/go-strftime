package strftime_test

import (
	"testing"
	"time"

	"github.com/ncruces/go-strftime"
)

const benchfmt = `%A %a %B %b %d %H %I %M %m %p %S %Y %y %Z`

func BenchmarkFormat(b *testing.B) {
	var t time.Time
	for i := 0; i < b.N; i++ {
		strftime.Format(benchfmt, t)
	}
}
