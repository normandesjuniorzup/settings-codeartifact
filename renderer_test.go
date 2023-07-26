package main_test

import (
	"bytes"
	"io"
	. "settings"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestRender(t *testing.T) {
	buf := bytes.Buffer{}
	token := "some-token"

	settingsRenderer, err := NewSettingsRenderer()
	if err != nil {
		t.Fatal(err)
	}

	if err := settingsRenderer.Render(&buf, token); err != nil {
		t.Fatal(err)
	}

	approvals.VerifyString(t, buf.String())
}

func BenchmarkRender(b *testing.B) {
	token := "bench-token"

	settingsRenderer, err := NewSettingsRenderer()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		settingsRenderer.Render(io.Discard, token)
	}
}
