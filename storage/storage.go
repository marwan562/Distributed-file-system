package storage

import (
	"io"
	"os"
)

type PathTransformFunc func(string) string

type StorageOpts struct {
	PathTransformFunc PathTransformFunc
}

func DefaultPathTransformFunc(s string) string {
	return s
}

type Storage struct {
	StorageOpts
}

func NewStorage(opts StorageOpts) *Storage {
	return &Storage{
		StorageOpts: opts,
	}
}

func (s *Storage) writeStream(key string, r io.Reader) error {
	pathname := s.PathTransformFunc(key)

	if err := os.MkdirAll(pathname, os.ModePerm); err != nil {
		return err
	}

	return nil
}
