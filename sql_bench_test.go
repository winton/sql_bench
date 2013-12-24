package main

import(
	// "fmt"
	"testing"
)

func BenchmarkQueryAllSingleTag(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllSingleTag()
	}
}

func BenchmarkQueryAllTwoTags(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllTwoTags()
	}
}