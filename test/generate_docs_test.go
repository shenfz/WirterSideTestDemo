package test

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var (
	zipFileName = "webHelpDARKCOFFIN2-all.zip"
	zipStoreDir = "bin"
	unZipDstDir = "docs"

	separator = string(filepath.Separator)
)

func Test_Gen_Docs(t *testing.T) {

	rootDir, _ := os.Getwd()
	absZipStoreDir := rootDir[:strings.LastIndex(rootDir, separator)] + separator + zipStoreDir
	unZipDstDir = rootDir[:strings.LastIndex(rootDir, separator)] + separator + unZipDstDir

	os.MkdirAll(unZipDstDir, os.ModePerm)

	zipFile, err := os.Open(absZipStoreDir + separator + zipFileName)
	if err != nil {
		fmt.Println("Failed to open the zip file:", err)
		return
	}
	defer zipFile.Close()

	fs, err := zipFile.Stat()
	if err != nil {
		t.Fatal(err)
	}

	// 创建一个新的zip.Reader来读取压缩包内容
	zipReader, err := zip.NewReader(zipFile, fs.Size())
	if err != nil {
		fmt.Println("Failed to create a new zip reader:", err)
		return
	}

	// 遍历压缩包中的每个文件
	for _, f := range zipReader.File {
		if f.Mode().IsDir() {
			if err := os.MkdirAll(unZipDstDir+separator+f.Name, os.ModePerm); err != nil {
				t.Fatal(err)
			} else {
				t.Log("make dir Succeed : ", unZipDstDir+separator+f.Name, os.ModePerm)
			}
			continue
			//t.Logf("Occur Dir : %s", f.Name)
		}

		//	t.Logf("Filename : %s", f.Name)
		// 打开压缩包内的文件
		err = writeUnZipFile(f, unZipDstDir)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func writeUnZipFile(f *zip.File, parentDir string) error {
	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	// 创建或打开目标文件以写入解压的数据
	outFilePath := parentDir + separator + f.Name
	outFile, err := os.Create(outFilePath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// 将压缩文件的内容复制到目标文件
	_, err = io.Copy(outFile, rc)
	if err != nil {
		return err
	}
	fmt.Printf("Successfully make file: %s\n", f.Name)
	return nil
}
