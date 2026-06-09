package convertion

import (
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"keneetic-antifilter/internal/pkg/cidr"
	"keneetic-antifilter/internal/pkg/list_reader"
)

type Convertor struct {
	reader       *list_reader.ListReader
	output       io.Writer
	skipPrefixes []string
	comment      string
}

func New(input io.Reader, output io.Writer, comment string) *Convertor {
	reader := list_reader.NewListReader(input)
	return &Convertor{
		reader:       reader,
		output:       output,
		skipPrefixes: []string{"#", "//"},
		comment:      comment,
	}
}

// CommentFromOutputPath derives a route comment (service name) from the output
// file path: the basename with a trailing "-ipv4.bat" removed, falling back to
// the basename without the ".bat" extension.
func CommentFromOutputPath(outputPath string) string {
	base := filepath.Base(outputPath)
	if trimmed := strings.TrimSuffix(base, "-ipv4.bat"); trimmed != base {
		return trimmed
	}
	return strings.TrimSuffix(base, ".bat")
}

func (c *Convertor) Next() (*Line, error) {
	line, err := c.reader.Next()
	switch {
	case errors.Is(err, io.EOF):
		return nil, io.EOF
	case err != nil:
		return nil, err
	case line == "":
		return c.Next()
	}

	for i := range c.skipPrefixes {
		if strings.HasPrefix(line, c.skipPrefixes[i]) {
			return c.Next()
		}
	}

	ip, err := cidr.Parse(line)
	if err != nil {
		return nil, fmt.Errorf("failed to parse CIDR: %w", err)
	}

	return &Line{
		Address: ip,
		Comment: c.comment,
	}, nil
}

func (c *Convertor) Write(line *Line) (err error) {
	message := fmt.Sprintf("route ADD %s MASK %s 0.0.0.0", line.Address.IP.String(), line.Address.Mask.String())
	if comment := line.Comment; comment != "" {
		message += " :: rem " + comment
	}

	_, err = c.output.Write([]byte(message + "\n"))
	if err != nil {
		return fmt.Errorf("failed to write to output: %w", err)
	}

	return nil
}
