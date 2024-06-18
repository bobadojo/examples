package main

import (
	"fmt"

	storespb "github.com/bobadojo/go/stores/v1/storespb"
)

func main() {
	s := storespb.Store{Description: "Hello"}
	fmt.Printf("%+v\n", s)
	fmt.Printf("%s\n", storespb.StoreLocator_GetStore_FullMethodName)
}
