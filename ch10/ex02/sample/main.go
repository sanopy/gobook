package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/sanopy/gobook/ch10/ex02/unarchive"
	_ "github.com/sanopy/gobook/ch10/ex02/unarchive/tar"
	_ "github.com/sanopy/gobook/ch10/ex02/unarchive/zip"
)

func main() {
	name, files, err := unarchive.Unarchive(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unarchive: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stderr, "Input format =", name)

	for _, file := range files {
		path := filepath.Join(".", file.Name)
		if file.Fileinfo.IsDir() {
			err = os.MkdirAll(path, file.Fileinfo.Mode())
			if err != nil {
				fmt.Fprintf(os.Stderr, "mkdirall: %v\n", err)
				os.Exit(1)
			}
		} else {
			if err = createFile(path, file); err != nil {
				fmt.Fprintf(os.Stderr, "createfile: %v\n", err)
				os.Exit(1)
			}
		}
	}
}

func createFile(path string, file *unarchive.File) error {
	fmt.Fprintf(os.Stderr, "x %s\n", path)
	destFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, file.Content); err != nil {
		return err
	}

	return nil
}
