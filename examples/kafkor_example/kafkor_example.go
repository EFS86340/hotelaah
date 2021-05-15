package main

import (
"fmt"
"os"

"github.com/EFS86340/hotelaah"
		)

func main() {
	if len(os.Args) < 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s <broker> <group> <topics..>\n", os.Args[0])
		os.Exit(1)
	}

	k := hotelaah.NewKafkor(os.Args[1], os.Args[2], os.Args[3:])
	k.Init()
	k.Listen()
}
