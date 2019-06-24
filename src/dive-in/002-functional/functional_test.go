package main

import (
	"fmt"
	"testing"
)

func TestYouGo(t *testing.T) {
	expected := "Goo!"
	gooStripper := stripper(youGo, " Yuth")

	if got := gooStripper; got != expected {
		t.Errorf("Returned %s instead of expected %s", got, expected)
	} else {
		fmt.Printf("Got %s and expected %s\n", got, expected)
	}
}
