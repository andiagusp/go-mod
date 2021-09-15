package sayhello

import "strconv"

// SayHello is exported
func SayHello(name string, year int) string {
	return "Hello my name " + name + " age " + strconv.Itoa(year)
}

// SayHay is exported
func SayHay(name string) string {
	return "Hai " + name
}
