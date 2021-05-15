package main

import (
		"fmt"
		"os"
		
		"github.com/EFS86340/hotelaah"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <broker> <topic>\n", os.Args[0])
		os.Exit(1)
	}
	k := hotelaah.NewKafkee(os.Args[2], os.Args[1])
	defer k.Disconnect()
	k.Init()

	sampleMsg := hotelaah.StringPair{
		First: "example_1st",
		Second: "example_2nd",
	}

		 for i := 0; i < 5; i++ {
			 k.Publish(&sampleMsg)
		 }


}
