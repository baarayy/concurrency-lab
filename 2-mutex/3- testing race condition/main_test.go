package main

import "testing"

func Test_updateMessage(t *testing.T) {
	msg = "Hello, House!"
	wg.Add(2)
	go updateMessage("Goodbye, House!")
	go updateMessage("Wait, House!")
	wg.Wait()

	if msg != "Goodbye, House!" {
		t.Errorf("updateMessage() = %v, want %v", msg, "Goodbye, House!")
	}
}
