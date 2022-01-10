package upfdata

import (
	"fmt"
	"strings"

	"github.com/muthuramanecs03g/nfcli/lib"
)

func ExecutorData(in string, promptConfig *lib.Prompt) {
	if strings.HasPrefix(in, "exit") {
		Exit()
		return
	}
	fmt.Println("Given: ", in)
}
