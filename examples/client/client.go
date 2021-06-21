package main

import (
				"fmt"

				pb "github.com/EFS86340/hotelaah"
	)

func printCities(client pb.HaalClient, prov *pb.PullRequest) {
		cities, err := client.PullCities(prov)
		for _, c := range cities {
				fmt.Printf("[rpc client] got %v", c)
		}
}

func main() {
}
