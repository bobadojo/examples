package main

import (
	"fmt"

	storespb "github.com/bobadojo/go/pkg/stores/v1/storespb"
)

func main() {
	s := storespb.Store{Description: "Hello"}
	fmt.Printf("%+v\n", s)
	fmt.Printf("%s\n", storespb.Stores_GetStore_FullMethodName)
}
