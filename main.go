package main

import (
	"fmt"
	"os"
)


func main() {
	_, err := os.Stat("proto/auth/auth.proto")
	if err != nil{
		fmt.Println(err)
	}
}