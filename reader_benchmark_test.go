package main

import (
	"context"
	"testing"
)

func BenchmarkReadAsync(b *testing.B) {
	ctx, cancel := context.WithCancel(context.Background())
	otr := NewOmdbTsvReader(WithContext(ctx), WithCancel(cancel))

	for n := 0; n < b.N; n++ {
		otr.ReadAsync()
	}
}
