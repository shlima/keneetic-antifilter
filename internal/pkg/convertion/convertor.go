package convertion

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"keneetic-antifilter/internal/pkg/cidr"
	"keneetic-antifilter/internal/pkg/list_reader"
)

type Convertor struct {
	reader       *list_reader.ListReader
	output       io.Writer
	skipPrefixes []string
}

func New(input io.Reader, output io.Writer) *Convertor {
	reader := list_reader.NewListReader(input)
	return &Convertor{
		reader:       reader,
		output:       output,
		skipPrefixes: []string{"#", "//"},
	}
}

func (c *Convertor) Next() error {
	line, err := c.reader.Next()
	switch {
	case errors.Is(err, io.EOF):
		return io.EOF
	case err != nil:
		return err
	case line == "":
		return nil
	}

	for i := range c.skipPrefixes {
		if strings.HasPrefix(line, c.skipPrefixes[i]) {
			return nil
		}
	}

	ip, err := cidr.Parse(line)
	if err != nil {
		return fmt.Errorf("failed to parse CIDR: %w", err)
	}

	message := fmt.Sprintf("route ADD %s MASK %s 0.0.0.0", ip.IP.String(), ip.Mask.String())
	_, err = c.output.Write([]byte(message + "\n"))
	if err != nil {
		return fmt.Errorf("failed to write to output: %w", err)
	}

	return nil
}
