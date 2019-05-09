package io

import (
	"io"
	"compress/gzip"
)

/*
compress 文件压缩
 */

func NewCompressReader(r io.Reader) (*gzip.Reader, error) { return gzip.NewReader(r) }
