package utstorage

import (
	"encoding/base64"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strings"
)

func UploadFileToLocal(encodedString string, path string) error {
	encodeStringSplit := strings.Split(encodedString, ",")
	if len(encodeStringSplit) > 1 {
		encodedString = encodeStringSplit[1]
	}

	dec, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		return err
	}
	if err := f.Sync(); err != nil {
		return err
	}
	return nil
}

func CreateFolder(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
	}
	return err
}

func CopyFile(src, dst string) error {
	input, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(dst, input, 0755)
	if err != nil {
		return err
	}
	return nil
}

func CopyFileFromMultipart(filename string, originalfile multipart.File) error {
	destFile, err := os.Create(filename)
	if err != nil {
		return err
	}

	if _, err = io.Copy(destFile, originalfile); err != nil {
		return err
	}
	return err
}
