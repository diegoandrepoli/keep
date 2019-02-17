package main

import (
   "testing"
)

/**
 * Test welcome message
 */
func TestWelcomeMessage(t *testing.T) {
   
  //get values
  expected := "Welcome to keep!"
  message := getIndexMessage()
  
  //test result
  if message != expected {
     t.Fatalf("expected err to be %v, but got %v", message, expected)
  }

}


