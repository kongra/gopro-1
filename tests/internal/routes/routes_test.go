package routes

import (
	"log/slog"
	"testing"

	"github.com/kongra/gopro-1/internal/routes"
)

type fibonacciTestData struct {
	name     string
	n        uint16
	expected uint64
}

func TestSomeExample(t *testing.T) {
	tests := []fibonacciTestData{
		{"fibonacci of 0", 0, 0},
		{"fibonacci of 1", 1, 1},
		{"fibonacci of 2", 2, 1},
		{"fibonacci of 3", 3, 2},
		{"fibonacci of 4", 4, 3},
		{"fibonacci of 5", 5, 5},

		// The last possible value within unit64 range
		{"fibonacci of 93", 93, 12200160415121876738},
	}

	for _, ftd := range tests {
		t.Run(ftd.name, func(t *testing.T) {
			result, _ := routes.Fibonacci(ftd.n)
			if result != ftd.expected {
				t.Errorf("Fibonacci(%d) = %d, expected %d", ftd.n, result, ftd.expected)
			}
		})
	}

	slog.Info("End of Fibonacci tests")
}

type fibonacciBenchData struct {
	name string
	n    uint16
}

func BenchmarkFibonacci(b *testing.B) {
	// Benchmark different Fibonacci numbers
	benchmarks := []fibonacciBenchData{
		{"Fib-5", 5},
		{"Fib-10", 10},
		{"Fib-20", 20},
		{"Fib-30", 30},
		{"Fib-40", 40},
		{"Fib-50", 50},
		{"Fib-93", 93}, // Maximum value for uint64
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for b.Loop() {
				_, _ = routes.Fibonacci(bm.n)
			}
		})
	}
}

func BenchmarkFibonacci20(b *testing.B) {
	for b.Loop() {
		_, _ = routes.Fibonacci(93)
	}
}

// Memory allocation benchmark
func BenchmarkFibonacciAllocs(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		_, _ = routes.Fibonacci(93)
	}
}
