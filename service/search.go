package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

const (
	lastResearch     = "/Users/CodeMania/.vpm/lastSearch.dat"
	indexFile        = "/Users/CodeMania/.vpm/index.dat"
	numGripLength    = 6
	nameGripLength   = 60
	suffixGripLength = 8
	sizeGripLength   = 12
	pathGripLength   = 60
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

			outputNum := resultNum + strings.Repeat(" ", numGripLength-len(resultNum))

			outputFileName := ""
			outputFileNameLength := getStringLengthInTerminal(fileName)
			if outputFileNameLength > nameGripLength {
				outputFileName = fileName[:nameGripLength-6] + "... "
				outputFileName = outputFileName + strings.Repeat(" ", nameGripLength-getStringLengthInTerminal(outputFileName))
			} else {
				outputFileName = fileName + strings.Repeat(" ", nameGripLength-outputFileNameLength)
			}

			outputSuffix := fileType + strings.Repeat(" ", suffixGripLength-len(fileType))

			outputSize := ""
			if len(fileInfoParts[3]) > sizeGripLength {
				outputSize = "e^10+"
			} else {
				outputSize = fileInfoParts[3] + strings.Repeat(" ", sizeGripLength-len(fileInfoParts[3]))
			}

			outputPath := ""
			outputPathLength := getStringLengthInTerminal(fileInfoParts[4])
			if outputPathLength > pathGripLength {
				outputPath = fileInfoParts[4][:pathGripLength-6] + "... "
				outputPath = outputPath + strings.Repeat(" ", pathGripLength-getStringLengthInTerminal(outputPath))
			} else {
				outputPath = fileInfoParts[4] + strings.Repeat(" ", pathGripLength-outputPathLength)
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

func getStringLengthInTerminal(s string) int {
	width := 0
	for _, c := range s {
		if utf8.RuneLen(c) >= 2 {
			width += 2
		} else {
			width += 1
		}
	}
	return width
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
	inputKeyword, inputFiletype := "", "nil"
	if len(os.Args) > 2 {
		inputKeyword = os.Args[1]
		inputFiletype = os.Args[2]
	} else {
		inputKeyword = os.Args[1]
	}

	if Find(inputKeyword, inputFiletype) {
		outputResults()
	} else {
		fmt.Println("Not Found")
	}
	return
}
