package main

import (
	"fmt"

	"github.com/mtagstyle/cineplex/pkg/theatres"
)

func main() {
	c := theatres.NewTheatresAPIClient()
	output, err := c.GetTheatres(nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", output)
}
