/*
This is free and unencumbered software released into the public domain.
Anyone is free to copy, modify, publish, use, compile, sell, or
distribute this software, either in source code form or as a compiled
binary, for any purpose, commercial or non-commercial, and by any
means.
In jurisdictions that recognize copyright laws, the author or authors
of this software dedicate any and all copyright interest in the
software to the public domain. We make this dedication for the benefit
of the public at large and to the detriment of our heirs and
successors. We intend this dedication to be an overt act of
relinquishment in perpetuity of all present and future rights to this
software under copyright law.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
OTHER DEALINGS IN THE SOFTWARE.
For more information, please refer to <http://unlicense.org/>
*/

package binarysearch

import (
	"math/rand"
	"sort"
	"testing"
)

var a []int

func randArray(ln int) []int {
	var a []int
	for i := 0; i < ln; i++ {
		a = append(a, rand.Intn(10))
	}
	return a
}

func benchmarkSearch(i int, b *testing.B) {
	a := randArray(i)
	sort.Ints(a)

	BinarySearch(a[rand.Intn(i)], a)
}

func BenchmarkBinarySearch10(b *testing.B) {
	benchmarkSearch(10, b)
}

func BenchmarkBinarySearch100(b *testing.B) {
	benchmarkSearch(100, b)
}

func BenchmarkBinarySearch1000(b *testing.B) {
	benchmarkSearch(1000, b)
}

func BenchmarkBinarySearch10000(b *testing.B) {
	benchmarkSearch(10000, b)
}

func BenchmarkBinarySearch100000(b *testing.B) {
	benchmarkSearch(100000, b)
}

func BenchmarkBinarySearch1000000(b *testing.B) {
	benchmarkSearch(1000000, b)
}

func BenchmarkBinarySearch10000000(b *testing.B) {
	benchmarkSearch(10000000, b)
}
