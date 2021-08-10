package runner

import (
	"fmt"
	"os"
)

func showVersion() {
	fmt.Printf("httpixy %s\n", version)
	os.Exit(2)
}
