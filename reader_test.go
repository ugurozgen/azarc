package main

import (
	"context"
	"testing"
)

func TestNewOmdbTsvReader(t *testing.T) {
	otr := NewOmdbTsvReader()

	if otr == nil {
		t.Error("OmdbTsvReader should be created.")
	}
}

func TestWithGoroutineCountOption(t *testing.T) {
	expectedValue := 2
	otr := NewOmdbTsvReader(WithGoroutineCount(expectedValue))

	if otr.goroutineCount != expectedValue {
		t.Errorf("Goroutine count should be %d.", expectedValue)
	}
}

func TestFilenameOption(t *testing.T) {
	expectedValue := "xyz"
	otr := NewOmdbTsvReader(Filename(expectedValue))

	if otr.fileName != expectedValue {
		t.Errorf("Filename should be %s.", expectedValue)
	}
}

func TestWithContextOption(t *testing.T) {
	expectedValue := context.Background()
	otr := NewOmdbTsvReader(WithContext(expectedValue))

	if otr.ctx != expectedValue {
		t.Errorf("Context should be provided.")
	}
}

func TestWithCancelOption(t *testing.T) {
	_, expectedValue := context.WithCancel(context.Background())
	otr := NewOmdbTsvReader(WithCancel(expectedValue))

	if otr.cancel == nil {
		t.Errorf("Cancel should be provided.")
	}
}

func TestReadAsync(t *testing.T) {
	_, expectedValue := context.WithCancel(context.Background())
	otr := NewOmdbTsvReader(WithCancel(expectedValue))

	outputCh, err := otr.ReadAsync()
	if err != nil {
		t.Errorf("Error should be nil.")
	}
	if outputCh == nil {
		t.Errorf("outputCh should not be nil")
	}
}
