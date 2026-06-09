package convertion

import (
	"errors"
	"io"
	"strings"
	"testing"
)

func runAll(t *testing.T, c *Convertor) {
	t.Helper()
	for {
		line, err := c.Next()
		if errors.Is(err, io.EOF) {
			return
		}
		if err != nil {
			t.Fatalf("Next returned error: %v", err)
		}
		err = c.Write(line)
		if err != nil {
			t.Fatalf("Write returned error: %v", err)
		}
	}
}

func TestConvertAppendsRemComment(t *testing.T) {
	input := strings.NewReader("91.108.4.0/22\n")
	var out strings.Builder

	c := New(input, &out, "telegram")
	runAll(t, c)

	want := "route ADD 91.108.4.0 MASK 255.255.252.0 0.0.0.0 :: rem telegram\n"
	if out.String() != want {
		t.Errorf("got %q, want %q", out.String(), want)
	}
}

func TestConvertWithoutCommentHasNoRem(t *testing.T) {
	input := strings.NewReader("91.108.4.0/22\n")
	var out strings.Builder

	c := New(input, &out, "")
	runAll(t, c)

	want := "route ADD 91.108.4.0 MASK 255.255.252.0 0.0.0.0\n"
	if out.String() != want {
		t.Errorf("got %q, want %q", out.String(), want)
	}
}

func TestCommentFromOutputPath(t *testing.T) {
	cases := map[string]string{
		"routes/telegram-ipv4.bat":   "telegram",
		"routes/cloudflare-ipv4.bat": "cloudflare",
		"telegram-ipv4.bat":          "telegram",
		"routes/all-ipv4-aa.bat":     "all-ipv4-aa",
		"something.bat":              "something",
	}

	for path, want := range cases {
		if got := CommentFromOutputPath(path); got != want {
			t.Errorf("CommentFromOutputPath(%q) = %q, want %q", path, got, want)
		}
	}
}
