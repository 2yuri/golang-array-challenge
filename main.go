package main

import (
	"fmt"

	"github.com/hyperyuri/interview-test/service"
)

func main() {
	array := []int{1, 2}

	service := service.NewArrayService()
	fmt.Printf("%v\n", service.ArrayChallenge(array))
}
