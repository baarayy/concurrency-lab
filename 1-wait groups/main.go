package main

import (
	"fmt"
	"sync"
)

func PrintSomething(s string , wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s);
}

func main() {

	var wg sync.WaitGroup;

	symbols := []string {
		"alpha",
		"beta",
		"gamma",
		"delta",
		"epsilon",
		"zeta",
		"eta",
		"theta",
		"iota",
		"kappa",
	}

	wg.Add(len(symbols));

	for _, symbol := range symbols {
		go PrintSomething(symbol, &wg);
	}
	wg.Wait();
	wg.Add(1);
	PrintSomething("runs inside main routine", &wg);
}