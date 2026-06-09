package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"keneetic-antifilter/internal/pkg/convertion"
)

func main() {
	inputPath := os.Args[1]
	outputPath := os.Args[2]

	fmt.Println("Input Path:", inputPath)
	fmt.Println("Output Path:", outputPath)

	input, err := os.Open(inputPath)
	if err != nil {
		panic(fmt.Errorf("failed to open input file: %v", err))
	}
	defer input.Close()

	// open output file
	output, err := os.Create(outputPath)
	if err != nil {
		panic(fmt.Errorf("failed to create output file: %v", err))
	}
	defer output.Close()

	comment := convertion.CommentFromOutputPath(outputPath)
	convertor := convertion.New(input, output, comment)

	var line *convertion.Line

LOOP:
	for {
		line, err = convertor.Next()
		switch {
		case errors.Is(err, io.EOF):
			fmt.Println("")
			break LOOP
		case err != nil:
			panic(fmt.Errorf("failed to read line: %v", err))
		case line.Address.IP.IsLoopback() || line.Address.IP.IsPrivate():
			fmt.Printf("> found local address: %s<", line.Address.IP.String())
			continue LOOP
		default:
			if err = convertor.Write(line); err != nil {
				panic(fmt.Errorf("failed to write line: %v", err))
			}
			fmt.Print(".")
		}
	}
}
