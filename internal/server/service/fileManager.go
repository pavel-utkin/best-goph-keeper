package service

import (
	"io"
	"os"
	"path/filepath"
	"strconv"
)

func CreateStorageUser(dirPath string, id int64) error {
	userId := strconv.Itoa(int(id))
	path := filepath.Join(dirPath, userId)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func CreateStorageNotExistsUser(dirPath string, id int64) error {
	userId := strconv.Itoa(int(id))
	path := filepath.Join(dirPath, userId)
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

func UploadFile(dirPath string, id int64, name string, data []byte) error {
	userId := strconv.Itoa(int(id))
	path := filepath.Join(dirPath, userId, "/", name)
	// Write data to file
	err := os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func DownloadFile(dirPath string, id int64, name string) ([]byte, error) {
	userId := strconv.Itoa(int(id))
	path := filepath.Join(dirPath, userId, "/", name)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func RemoveFile(dirPath string, id int64, name string) error {
	userId := strconv.Itoa(int(id))
	path := filepath.Join(dirPath, userId, "/", name)
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}
