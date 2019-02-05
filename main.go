package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"time"
)

var args = os.Args[1:]
var sourceFilePath = args[0]
var destinationDirectory = args[1]
var fileSeparator = getFileSpearator()
var pathBlocks = strings.Split(sourceFilePath, fileSeparator)
var fileName = pathBlocks[len(pathBlocks)-1]
var targetFile = destinationDirectory + fileSeparator + fileName

func main() {
	defer os.Exit(0)
	zipFilePath := destinationDirectory + fileSeparator + newFileName() + ".zip"

	var funcs []func() (error, string)
	funcs = append(funcs, func() (error, string) { return copy(sourceFilePath, targetFile), "File copy successful" })
	funcs = append(funcs, func() (error, string) { return archive(zipFilePath, targetFile), "Zipping complete" })
	funcs = append(funcs, func() (error, string) { return remove(targetFile), "Removed leftovers" })

	for i := 0; i < len(funcs); i++ {
		err, successMsg := funcs[i]()
		if err != nil {
			panic(err)
		} else {
			fmt.Println(successMsg)
		}
	}
	fmt.Println("All done!")
}

func getFileSpearator() string {
	if runtime.GOOS == "windows" {
		return "\\"
	}
	return "/"
}

func newFileName() string {
	currentTime := time.Now().Local()
	return currentTime.Format("2006-01-02")
}

func copy(src string, targetFile string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	file, err := os.Create(targetFile)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, source)
	return err
}

func archive(zipFilePath, targetFile string) error {
	newZipFile, err := os.Create(zipFilePath)
	if err != nil {
		return err
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	zipFile, err := os.Open(targetFile)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	info, err := zipFile.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Method = zip.Deflate // Compression level 8

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, zipFile)
	if err != nil {
		return err
	}
	return nil
}

func remove(file string) error {
	err := os.Remove(file)
	if err != nil {
		return err
	}
	return nil
}
