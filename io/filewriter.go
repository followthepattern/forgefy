package io

import "os"

type FileWriter struct {
}

func (fw FileWriter) Write(fileName string, content string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}

	_, err = file.Write([]byte(content))
	if err != nil {
		return err
	}

	return nil
}
