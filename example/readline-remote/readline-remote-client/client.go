package main

import "github.com/acatton/goreadline-ng"

func main() {
	if err := readline.DialRemote("tcp", ":12344"); err != nil {
		println(err.Error())
	}
}
