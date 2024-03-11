package common

import (
	"io"
	"mime/multipart"
	"os"
)

type FileSystem struct {
	File *multipart.FileHeader
}

func (fs *FileSystem) CheckPath(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	} else {
		return err
	}
	return nil
}

func (fs *FileSystem) Upload(dst string) error {
	src, err := fs.File.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
