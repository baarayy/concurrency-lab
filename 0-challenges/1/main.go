package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string , wg *sync.WaitGroup) {
	defer wg.Done()
	msg = s
	printMessage()
}

func printMessage() {
	fmt.Println(msg)
}

func main() {


	msg = "Hello, world!"
	var wg sync.WaitGroup;
	wg.Add(3);

	go updateMessage("Hello, universe!" , &wg)

	go updateMessage("Hello, cosmos!" , &wg)

	go updateMessage("Hello, world!" , &wg)

	wg.Wait();
}
