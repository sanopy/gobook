package unarchive

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"io/fs"
)

var ErrFormat = errors.New("unknown format")

type format struct {
	name   string
	magic  string
	offset int
	decode func(io.Reader) ([]*File, error)
}

var formats []format

func RegisterFormat(name, magic string, offset int, decode func(io.Reader) ([]*File, error)) {
	formats = append(formats, format{name, magic, offset, decode})
}

type File struct {
	Name     string
	Fileinfo fs.FileInfo
	Content  *bytes.Buffer
}

func Unarchive(r io.Reader) (string, []*File, error) {
	br := bufio.NewReader(r)
	f, err := sniff(br)
	if err != nil {
		return "", nil, err
	}
	file, err := f.decode(br)
	return f.name, file, err
}

func sniff(r *bufio.Reader) (*format, error) {
	for _, f := range formats {
		n := f.offset + len(f.magic)
		header, err := r.Peek(n)
		if err != nil {
			return nil, err
		}
		magic := string(header[f.offset:])
		if magic == f.magic {
			return &f, nil
		}
	}
	return nil, ErrFormat
}
