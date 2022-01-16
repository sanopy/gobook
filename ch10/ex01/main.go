package main

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		printUsage()
	}

	switch os.Args[1] {
	case "-gif":
		if err := toGIF(os.Stdin, os.Stdout); err != nil {
			fmt.Fprintf(os.Stderr, "gif: %v\n", err)
			os.Exit(1)
		}
	case "-jpeg":
		if err := toJPEG(os.Stdin, os.Stdout); err != nil {
			fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
			os.Exit(1)
		}
	case "-png":
		if err := toPNG(os.Stdin, os.Stdout); err != nil {
			fmt.Fprintf(os.Stderr, "png: %v\n", err)
			os.Exit(1)
		}
	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Fprintf(os.Stderr, "usage: $ %s [format]\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  format: -gif | -jpeg | -png\n")
	os.Exit(1)
}

func readImg(in io.Reader) (image.Image, error) {
	img, kind, err := image.Decode(in)
	if err != nil {
		return nil, err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return img, nil
}

func toGIF(in io.Reader, out io.Writer) error {
	img, err := readImg(in)
	if err != nil {
		return err
	}
	return gif.Encode(out, img, &gif.Options{NumColors: 256})
}

func toJPEG(in io.Reader, out io.Writer) error {
	img, err := readImg(in)
	if err != nil {
		return err
	}
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

func toPNG(in io.Reader, out io.Writer) error {
	img, err := readImg(in)
	if err != nil {
		return err
	}
	return png.Encode(out, img)
}
