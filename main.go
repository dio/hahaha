package main

import (
	"fmt"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/ok/ok/haha"
)

const json = "{\"haha\": \"HTTP\"}"

func main() {
	ok := haha.OK{}
	if err := jsonpb.UnmarshalString(json, &ok); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%v\n", ok.Haha)
}
