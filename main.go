package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var (
	zipFilePath string
	rootPath    string
)

func init() {
	flag.StringVar(&zipFilePath, "f", "./bin/webHelpDARKCOFFIN2-all.zip", "input zip filepath ")
	flag.StringVar(&rootPath, "d", "./docs", "unzip root path")
}

func main() {
	flag.Parse()
	zipFilePath, _ = filepath.Abs(zipFilePath)
	// 打开ZIP文件
	//r, err := os.Open(zipFilePath)
	//if err != nil {
	//	fmt.Println("Error opening ZIP file:", err)
	//	return
	//}
	//defer r.Close()
	os.MkdirAll(rootPath, 0666)
	os.Chdir(rootPath)

	fd, _ := zip.OpenReader(zipFilePath)
	for _, f := range fd.File {
		file, err := f.Open()
		if err != nil {
			fmt.Println(err)
			continue
		}
		data, _ := io.ReadAll(file)
		os.WriteFile(f.Name, data, 0644)
		file.Close()
	}
}
