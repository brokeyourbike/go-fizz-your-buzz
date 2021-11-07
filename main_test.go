package main

import "testing"

func TestFizz(t *testing.T) {
	want := "fizz"
	if got := fizzbuzz(3); got != want {
		t.Errorf("fizzbuzz() = %q, want %q", got, want)
	}
}
func TestBuzz(t *testing.T) {
	want := "buzz"
	if got := fizzbuzz(5); got != want {
		t.Errorf("fizzbuzz() = %q, want %q", got, want)
	}
}
func TestFizzBuzz(t *testing.T) {
	want := "fizz buzz"
	if got := fizzbuzz(15); got != want {
		t.Errorf("fizzbuzz() = %q, want %q", got, want)
	}
}
func TestNumber(t *testing.T) {
	want := "1"
	if got := fizzbuzz(1); got != want {
		t.Errorf("fizzbuzz() = %q, want %q", got, want)
	}
}
