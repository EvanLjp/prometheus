package debug

import "fmt"

var (
	debug = false
)

func LOG(str string) {
	if debug {
		fmt.Println(str)
	}
}
