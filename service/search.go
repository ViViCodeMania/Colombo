package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	lastResearch = "/Users/CodeMania/.vpm/lastSearch.dat"
	indexFile    = "/Users/CodeMania/.vpm/index.dat"
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
			resultNum := strconv.Itoa(resultCount)
			results = append(results, resultNum+" "+fileInfo)

			outputNum := resultNum + strings.Repeat(" ", 6-len(resultNum))

			outputFileName := ""
			if len(fileName) > 32 {
				outputFileName = fileName[:28] + "... "
			} else {
				outputFileName = fileName + strings.Repeat(" ", 32-len(fileName))
			}

			outputSuffix := fileType + strings.Repeat(" ", 8-len(fileType))

			outputSize := ""
			if len(fileInfoParts[3]) > 10 {
				outputSize = "e^10+"
			} else {
				outputSize = fileInfoParts[3] + strings.Repeat(" ", 10-len(fileInfoParts[3]))
			}

			outputPath := ""
			if len(fileInfoParts[4]) > 48 {
				outputPath = fileInfoParts[4][:44] + "... "
			} else {
				outputPath = fileInfoParts[4] + strings.Repeat(" ", 48-len(fileInfoParts[4]))
			}

			outputModDate := fileInfoParts[5]
			outputString := outputNum + outputFileName + outputSuffix + outputSize + outputPath + outputModDate
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
	inputKeyword := os.Args[1]
	inputFiletype := os.Args[2]

	if Find(inputKeyword, inputFiletype) {
		outputResults()
	} else {
		fmt.Println("Not Found")
	}
	return
}
