package main

import "testing"

func TestMessage(t *testing.T) {
    expected := "Hello from my automated Jenkins!"
    // This is a simple placeholder test
    if expected == "" {
        t.Errorf("Expected a message, but got empty string")
    }
}
