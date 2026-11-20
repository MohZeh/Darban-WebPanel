package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestVersion(t *testing.T) {
	var out bytes.Buffer
	if code := run([]string{"-version"}, &out); code != 0 {
		t.Fatalf("exit code %d, want 0", code)
	}
	if got := strings.TrimSpace(out.String()); got != version {
		t.Fatalf("printed %q, want %q", got, version)
	}
}

func TestNotReadyWithoutVersion(t *testing.T) {
	var out bytes.Buffer
	if code := run(nil, &out); code != 1 {
		t.Fatalf("exit code %d, want 1", code)
	}
}

func TestUnknownFlag(t *testing.T) {
	var out bytes.Buffer
	if code := run([]string{"-no-such-flag"}, &out); code != 2 {
		t.Fatalf("exit code %d, want 2", code)
	}
}
