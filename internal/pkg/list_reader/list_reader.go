package list_reader

import (
	"bufio"
	"io"
)

type ListReader struct {
	scanner *bufio.Scanner
}

func NewListReader(reader io.Reader) *ListReader {
	scanner := bufio.NewScanner(reader)
	return &ListReader{
		scanner: scanner,
	}
}

// Next returns next string
func (f *ListReader) Next() (out string, err error) {

	switch {
	case f.scanner.Scan():
		return f.scanner.Text(), nil
	case f.scanner.Err() != nil:
		return "", f.scanner.Err()
	default:
		return "", io.EOF
	}
}
