package main

import "time"

func main() {
	go hello()
	time.Sleep(1 * time.Second)
	aaa()
}

func hello() {
	panic("adas")
}

func aaa() {
	panic("adsa")
}
