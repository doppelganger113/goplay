// Sample benchmarks to test which function is better for converting an
// integer into a string. First using the fmt.Sprintf function, then
// the strconv. Formatint function and then strconv.Itoa
/*
	Running:
		- go test -v -run="none" -bench=. -benchtime="3s"
		- go test -v -run="none" -bench=. -benchtime="3s" -benchmem

	Results:

	BenchmarkSprintf-4      50000000               119 ns/op
	BenchmarkFormat-4       1000000000             4.14 ns/op
	BenchmarkItoA-4         1000000000             6.11 ns/op
 */
package one_test

import (
	"testing"
	"fmt"
	"strconv"
)

// BenchmarkSprintf provides performance numbers for the fmt.Sprintf function
func BenchmarkSprintf(b *testing.B) {
	number := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", number)
	}
}

// BenchMarkFormat provides performance numbers for the strconv.FormatInt function
func BenchmarkFormat(b *testing.B) {
	number := int64(10)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.FormatInt(number, 10)
	}
}

// BenchMarkItoA provides performance numbers for the strconv.Itoa function
func BenchmarkItoA(b *testing.B) {

	number := 10

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.Itoa(number)
	}
}
