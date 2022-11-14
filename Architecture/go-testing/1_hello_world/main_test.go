package main

import "testing"

func TestHello(t *testing.T) {
    got := Hello("Mario")
    //want := "Hello, world"
    want := "Hello, Mario"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
