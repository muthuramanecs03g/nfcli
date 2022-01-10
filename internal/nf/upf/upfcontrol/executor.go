package upfcontrol

import (
	"fmt"
	"strings"

	"github.com/muthuramanecs03g/nfcli/lib"
)

func ExecutorControl(in string, promptConfig *lib.Prompt) {
	if strings.HasPrefix(in, "exit") {
		Exit()
		return
	}

	fmt.Println("Given: ", in)
}
