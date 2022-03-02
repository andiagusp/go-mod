package sayhello

import (
	"fmt"
	"strconv"
)

func SayHello(name string, year int) string {
	return "Hello my name " + name + " age " + strconv.Itoa(year)
}

func SayHay(name string) string {
	return "Hai " + name
}

func HelloGo() {
	fmt.Println("Hello Go")
}

func HelloGoV2() {
	fmt.Println("Hello Go V2")
}
