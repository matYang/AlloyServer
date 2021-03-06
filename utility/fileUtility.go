package utility

import (
	"io"
	"os"
)

//check if the file exists, return true if not exist
func FileNotExist(filename string) bool {
	_, err := os.Stat(filename)
	return os.IsNotExist(err)
}

func CreateDirectoryIfNotExist(path string) error {
	if DirectoryNotExist(path) {
		return CreateDirectory(path)
	}
	return nil
}

func DirectoryNotExist(path string) bool {
	src, err := os.Stat(path)
	if os.IsNotExist(err) {
		return true
	}
	if !src.IsDir() {
		panic("Fatal Semantic, Given Path: " + path + " is not directory")
	}
	return false
}

func CreateDirectory(path string) (err error) {
	err = os.Mkdir(path, 0777)
	return
}

func RemoveDirectory(path string) (err error) {
	err = os.RemoveAll(path)
	return
}

func MoveFile(src, dest string) (err error) {
	err = os.Rename(src, dest)
	return
}

func RemoveFile(filename string) (err error) {
	err = os.RemoveAll(filename)
	return
}

func DeepCopyFile(src, dest string) (err error) {
	// open files r and w
	r, err := os.Open(src)
	if err != nil {
		return
	}
	defer r.Close()

	w, err := os.Create(dest)
	if err != nil {
		return
	}
	defer w.Close()

	// do the actual work
	_, err = io.Copy(w, r)
	return
}
