package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	targetPath = "/Users/CodeMania/Documents/编程相关"
	indexFile  = "/Users/CodeMania/.vpm/index.dat"
)

var files []string

func listFunc(path string, f os.FileInfo, err error) error {
	if f == nil {
		return err
	}

	if f.IsDir() {
		return nil
	}

	nameParts := strings.Split(f.Name(), ".")
	fileName := strings.ToLower(strings.Replace(nameParts[0], " ", "_", -1))

	suffix := "nil"
	if len(nameParts) > 1 {
		suffix = nameParts[len(nameParts)-1]
	}
	size := strconv.FormatInt(f.Size(), 10)
	path = "." + strings.Replace(path, " ", "_", -1)[39:]
	modtime := f.ModTime().Format(time.RFC3339)

	fInfo := fileName + " " + suffix + " " + size + " " + path + " " + modtime + "\n"
	files = append(files, fInfo)
	return nil
}

func FileListUpdate(path string) {
	err := filepath.Walk(path, listFunc)
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}

	os.Remove(indexFile)

	outputFile, outputError := os.OpenFile(indexFile, os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)

	for i := 0; i < len(files); i++ {
		num := strconv.Itoa(i+1) + " "
		outputWriter.WriteString(num + files[i])
	}
	outputWriter.Flush()
	return
}

func main() {
	FileListUpdate(targetPath)
}
