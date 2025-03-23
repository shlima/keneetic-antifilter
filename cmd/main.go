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

	convertor := convertion.New(input, output)

LOOP:
	for {
		err = convertor.Next()
		switch {
		case errors.Is(err, io.EOF):
			break LOOP
		case err != nil:
			panic(fmt.Errorf("failed to read line: %v", err))
		default:
			fmt.Print(".")
		}
	}
}
