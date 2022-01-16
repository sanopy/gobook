package tar

import (
	"archive/zip"
	"bytes"
	"io"
	"io/ioutil"
	"os"

	"github.com/sanopy/gobook/ch10/ex02/unarchive"
)

const (
	magic  = "PK\x03\x04"
	offset = 0
)

func init() {
	unarchive.RegisterFormat("zip", magic, offset, decode)
}

func decode(r io.Reader) ([]*unarchive.File, error) {
	var readerAt io.ReaderAt
	var size int64
	switch v := r.(type) {
	case *os.File:
		readerAt = v
		fileInfo, err := v.Stat()
		if err != nil {
			return nil, err
		}
		size = fileInfo.Size()
	default:
		buf, err := ioutil.ReadAll(r)
		if err != nil {
			return nil, err
		}
		readerAt = bytes.NewReader(buf)
		size = int64(len(buf))
	}

	zipReader, err := zip.NewReader(readerAt, size)
	if err != nil {
		return nil, err
	}

	var ret []*unarchive.File
	for _, file := range zipReader.File {
		r, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer r.Close()

		buf := new(bytes.Buffer)
		if _, err := io.Copy(buf, r); err != nil {
			return nil, err
		}

		f := &unarchive.File{
			Name:     file.Name,
			Fileinfo: file.FileInfo(),
			Content:  buf}
		ret = append(ret, f)
	}

	return ret, nil
}
