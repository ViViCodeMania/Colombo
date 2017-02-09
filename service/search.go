package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"io"
)

const (
	lastResearch = "/Users/CodeMania/.vpm/lastSearch.dat"
	indexFile = "/Users/CodeMania/.vpm/index.dat"
)

var results []string

func Find(inputKeyword string, inputFiletype string) bool {
	inputKeyword = strings.ToLower(inputKeyword)
	inputFiletype = strings.ToLower(inputFiletype)

	inputFile, inputError := os.Open(indexFile)
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile")
		return false
	}
	defer inputFile.Close()

	resultCount := 0
	inputReader := bufio.NewReader(inputFile)
	for {
		isKeywordMacth := false

		fileInfo, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			break
		}
		fileInfoParts := strings.Split(fileInfo, " ")

		if len(fileInfoParts) != 6 {
			continue
		}

		fileName := fileInfoParts[1]
		fileType := fileInfoParts[2]

		if inputFiletype == "nil" && strings.Contains(fileName, inputKeyword) {
			isKeywordMacth = true
		} else if inputFiletype != "nil" && inputFiletype == fileType && strings.Contains(fileName, inputKeyword) {
			isKeywordMacth = true
		}

		if isKeywordMacth {
			resultCount++
			outputNum := strconv.Itoa(resultCount)
			outputString := outputNum + " " + fileInfo
			results = append(results, outputString)
			fmt.Print(outputString)
		}

	}

	if resultCount > 0 {
		return true
	}

	return false

}

func outputResults() {
	os.Remove(lastResearch)

	outputFile, outputError := os.OpenFile(lastResearch, os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	for i := 0; i < len(results); i++ {
		outputWriter.WriteString(results[i])
	}
	outputWriter.Flush()
	return
}

func main() {
	//argNum := len(os.Args)
	inputKeyword := "Java"
	inputFiletype := "pdf"

	if Find(inputKeyword, inputFiletype) {
		outputResults()
	} else {
		fmt.Println("Not Found")
	}
	return
}


