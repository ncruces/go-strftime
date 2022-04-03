package strftime_test

import (
	"fmt"
	"log"

	"github.com/ncruces/go-strftime"
)

func ExampleLayout() {
	layout, err := strftime.Layout("%Y-%m-%d %H:%M:%S")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%q", layout)
	// Output:
	// "2006-01-02 15:04:05"
}
