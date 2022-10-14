package main

import (
	"fmt"
	"go-unit-test/repository"
)

type any struct {
	name string
}

func main() {
	me := any{"Uji"}
	fmt.Println(me)

	cube := repository.cubeRepository{1}

}
