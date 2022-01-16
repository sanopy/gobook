package tar

import (
	"archive/tar"
	"bytes"
	"io"

	"github.com/sanopy/gobook/ch10/ex02/unarchive"
)

const (
	magic1 = "ustar\x0000"
	magic2 = "ustar  \x00"
	offset = 257
)

func init() {
	unarchive.RegisterFormat("tar", magic1, offset, decode)
	unarchive.RegisterFormat("tar", magic2, offset, decode)
}

func decode(r io.Reader) ([]*unarchive.File, error) {
	var ret []*unarchive.File
	tarReader := tar.NewReader(r)
	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		buf := new(bytes.Buffer)
		if _, err := io.Copy(buf, tarReader); err != nil {
			return nil, err
		}

		f := &unarchive.File{
			Name:     header.Name,
			Fileinfo: header.FileInfo(),
			Content:  buf}
		ret = append(ret, f)
	}

	return ret, nil
}
