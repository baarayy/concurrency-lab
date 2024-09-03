package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

var wg sync.WaitGroup

func Test_updateMessage(t *testing.T) {
	wg.Add(1)
	go updateMessage("Hello, universe!", &wg)
	wg.Wait()
	if msg != "Hello, universe!" {
		t.Errorf("updateMessage() = %v, want %v", msg, "Hello, universe!")
	}
}

func Test_printMessage(t *testing.T) {
	msg = "Hello, universe!"
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	printMessage()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdout

	if !strings.Contains(output, "Hello, universe!") {
		t.Errorf("printMessage() = %v, want %v", output, "Hello, universe!")
	}
}

func Test_main(t *testing.T) {
	msg = "Hello, world!"
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdout

	if !strings.Contains(output, "Hello, universe!") {
		t.Errorf("main() = %v, want %v", output, "Hello, universe!")
	}
	if !strings.Contains(output, "Hello, cosmos!") {
		t.Errorf("main() = %v, want %v", output, "Hello, cosmos!")
	}
	if !strings.Contains(output, "Hello, world!") {
		t.Errorf("main() = %v, want %v", output, "Hello, world!")
	}
}