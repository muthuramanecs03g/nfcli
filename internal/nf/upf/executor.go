package upf

import (
	"fmt"
	"strings"
)

// Executor parse CLI
func Executor(in string) {
	if strings.HasPrefix(in, "stats") {
		fmt.Println("Handled: ", in)
	} else {
		fmt.Println("UnHandled: ", in)
	}
}
