package main

import "testing"

func Benchmark_Mutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := SumSquaresMutex(2, 4, 6, 8, 10)
		if res != 220 {
			b.Fatal("not equal 220")
		}
	}
}

func Benchmark_Channel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := SumSquaresChannel(2, 4, 6, 8, 10)
		if res != 220 {
			b.Fatal("not equal 220")
		}
	}
}

func Benchmark_ChannelForRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := SumSquaresChannelForRange(2, 4, 6, 8, 10)
		if res != 220 {
			b.Fatal("not equal 220")
		}
	}
}

func Benchmark_Atomic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := SumSquaresAtomic(2, 4, 6, 8, 10)
		if res != 220 {
			b.Fatal("not equal 220")
		}
	}
}
