package test

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var (
	stwRootDir   string
	stwSeparator string
	jsonFileName = "dir.json"

	rootNode Node
)

type Node struct {
	Text     string `json:"text"`
	Children []Node `json:"children"`
}

func Test_Generate_Dir(t *testing.T) {
	loadJson()
	parseNode(rootNode, "")
}

func loadJson() {
	stwSeparator = string(filepath.Separator)
	stwWorkDir, _ := os.Getwd()
	stwRootDir = stwWorkDir[:strings.LastIndex(stwWorkDir, stwSeparator)]
	fdJsonFileBytes, err := os.ReadFile(stwWorkDir + stwSeparator + jsonFileName)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fdJsonFileBytes, &rootNode)
	if err != nil {
		panic(err)
	}
}

func parseNode(node Node, stParentDir string) {
	if node.Text != "" {
		createDir(node, stParentDir)
	}

	if stParentDir != "" {
		stParentDir = stParentDir + stwSeparator
	}
	if node.Text != "" {
		stParentDir = stParentDir + node.Text
	}

	for _, child := range node.Children {
		parseNode(child, stParentDir)
	}
}

func createDir(node Node, stParentDir string) {
	stDirPath := stwRootDir + stwSeparator
	if stParentDir != "" {
		stDirPath = stDirPath + stwSeparator + stParentDir
	}
	stDirPath = stDirPath + stwSeparator + node.Text

	err := os.MkdirAll(filepath.Clean(stDirPath), os.ModePerm)
	if err != nil {
		panic(err)
	}
	log.Println("Create Succeed : " + stDirPath)
}
