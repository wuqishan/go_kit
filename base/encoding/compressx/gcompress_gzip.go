package compressx

import (
	"bytes"
	"compress/gzip"
	"github.com/wuqishan/go_kit/base/file"
	"io"
)

// Gzip compresses `data` using gzip algorithm.
// The optional parameter `level` specifies the compression level from
// 1 to 9 which means from none to the best compression.
//
// Note that it returns error if given `level` is invalid.
func Gzip(data []byte, level ...int) ([]byte, error) {
	var (
		writer *gzip.Writer
		buf    bytes.Buffer
		err    error
	)
	if len(level) > 0 {
		writer, err = gzip.NewWriterLevel(&buf, level[0])
		if err != nil {
			return nil, err
		}
	} else {
		writer = gzip.NewWriter(&buf)
	}
	if _, err = writer.Write(data); err != nil {
		return nil, err
	}
	if err = writer.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// GzipFile compresses the file `src` to `dst` using gzip algorithm.
func GzipFile(srcFilePath, dstFilePath string, level ...int) (err error) {
	dstFile, err := file.Create(dstFilePath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	return GzipPathWriter(srcFilePath, dstFile, level...)
}

// GzipPathWriter compresses `filePath` to `writer` using gzip compressing algorithm.
//
// Note that the parameter `path` can be either a directory or a file.
func GzipPathWriter(filePath string, writer io.Writer, level ...int) error {
	var (
		gzipWriter *gzip.Writer
		err        error
	)
	srcFile, err := file.Open(filePath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	if len(level) > 0 {
		gzipWriter, err = gzip.NewWriterLevel(writer, level[0])
		if err != nil {
			return err
		}
	} else {
		gzipWriter = gzip.NewWriter(writer)
	}
	defer gzipWriter.Close()

	if _, err = io.Copy(gzipWriter, srcFile); err != nil {
		return err
	}
	return nil
}

// UnGzip decompresses `data` with gzip algorithm.
func UnGzip(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	reader, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	if _, err = io.Copy(&buf, reader); err != nil {
		return nil, err
	}
	if err = reader.Close(); err != nil {
		return buf.Bytes(), err
	}
	return buf.Bytes(), nil
}

// UnGzipFile decompresses srcFilePath `src` to `dst` using gzip algorithm.
func UnGzipFile(srcFilePath, dstFilePath string) error {
	srcFile, err := file.Open(srcFilePath)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	dstFile, err := file.Create(dstFilePath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	reader, err := gzip.NewReader(srcFile)
	if err != nil {
		return err
	}
	defer reader.Close()

	if _, err = io.Copy(dstFile, reader); err != nil {
		return err
	}
	return nil
}
