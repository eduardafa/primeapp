package main

import "testing"

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{name: "zero", testNum: 0, expected: false, msg: "0 is not a prime number!"},
		{name: "one", testNum: 1, expected: false, msg: "1 is not a prime number!"},
		{name: "not prime", testNum: 8, expected: false, msg: "8 is not a prime number because it is divisible by 2!"},
		{name: "negative number", testNum: -1, expected: false, msg: "Negative numbers are not prime!"},
		{name: "prime", testNum: 7, expected: true, msg: "7 is a prime number!"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		// true && false
		if e.expected && !result {
			t.Errorf("%s: expected true, but got false", e.name)
		}

		// false && true
		if !e.expected && result {
			t.Errorf("%s: expected false, but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s, but got %s", e.name, e.msg, msg)
		}
	}
}
