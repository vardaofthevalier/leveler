package util

import (
	"io"
	"os"
	"errors"
)

func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	_, err = os.Stat(dst)
	if err == nil {
		return errors.New("Destination file already exists!")
	}

	out, err := os.Create(dst) 
	if err != nil {
		return err 
	}
	defer out.Close()

	_, err = io.Copy(in, out)
	if err != nil {
		return err
	}

	return nil
}