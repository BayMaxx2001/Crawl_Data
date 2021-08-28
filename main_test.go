package main

import (
	"sync"
	"testing"
)

func Benchmark10(b *testing.B) {
	benchmark(b, make(chan int, 10))
}
func Benchmark100(b *testing.B) {
	benchmark(b, make(chan int, 100))
}
func Benchmark1000(b *testing.B) {
	benchmark(b, make(chan int, 1000))
}
func BenchmarkUnbuffered(b *testing.B) {
	benchmark(b, make(chan int))
}
func BenchmarkHeavy10(b *testing.B) {
	benchmarkHeavy(b, make(chan int, 10))
}
func BenchmarkHeavy100(b *testing.B) {
	benchmarkHeavy(b, make(chan int, 100))
}
func BenchmarkHeavy1000(b *testing.B) {
	benchmarkHeavy(b, make(chan int, 1000))
}
func BenchmarkHeavyUnbuffered(b *testing.B) {
	benchmarkHeavy(b, make(chan int))
}
func benchmark(b *testing.B, c chan int) {
	limit := 10
	var wg sync.WaitGroup
	wg.Add(limit)
	for i := 0; i < limit; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for range c {
				// handling
			}
		}(&wg)
	}
	for i := 0; i < b.N; i++ {
		c <- i
	}
	close(c)
	wg.Wait()
}
func benchmarkHeavy(b *testing.B, c chan int) {
	limit := 10
	var wg sync.WaitGroup
	wg.Add(limit)
	for i := 0; i < limit; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for range c {
				main()
			}
		}(&wg)
	}
	for i := 0; i < b.N; i++ {
		c <- i
	}
	close(c)
	wg.Wait()
}
