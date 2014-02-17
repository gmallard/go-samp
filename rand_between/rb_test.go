package main

import "testing"

func BenchmarkBetween(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = randBetween(min, max)
	}
}

func BenchmarkBetween2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = randBetween2(min, max)
	}
}

func BenchmarkBetween3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = randBetween3(min, max, 1.0)
	}
}
