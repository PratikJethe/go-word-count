package wordcount

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/pratikjethe/go-word-count/cmd"
)

var (
	overallWordCount      int
	overallLineCount      int
	overallCharacterCount int
	printTotal            bool
)

type searchResult struct {
	wordCount      int
	lineCount      int
	characterCount int
	inputFlags     cmd.InputFlags
	fileName       string
	err            error
}

// StartSearch acts like and entry point for word count. It decides to call serachFile or searchUserInpuy based on input params
// It does file searching concurrently while supporting max open file descriptors
func StartSearch(input cmd.Input) {

	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup
	var outputChannel = make(chan searchResult)
	var doneChannel = make(chan bool)
	var maxOpenFiles int = 100
	var limitMaxOpenFileChannel = make(chan int, maxOpenFiles)

	wg1.Add(1)
	go handleOutput(outputChannel, doneChannel, &wg1)
	if len(input.FileList) != 0 {
		for _, filename := range input.FileList {
			limitMaxOpenFileChannel <- 1
			wg2.Add(1)
			go func(filepath string) {
				serachFile(filepath, input.InputFlags, outputChannel)
				<-limitMaxOpenFileChannel
				wg2.Done()

			}(filename)

		}
	}

	if len(input.UserInput) != 0 {
		searchUserInput(input.UserInput, input.InputFlags, outputChannel)

	}

	wg2.Wait()

	if len(input.FileList) > 1 || printTotal {
		outputChannel <- searchResult{
			wordCount:      overallWordCount,
			lineCount:      overallLineCount,
			characterCount: overallCharacterCount,
			inputFlags:     input.InputFlags,
			fileName:       "total",
		}
	}
	doneChannel <- true
	wg1.Wait()

}

// serachFile is responsible to open a file and get required counts
// It also checks if a given path is of an directory and walks through that directory performing an recursive serach on underlying files
// It also has error hadling and concurrecy support
func serachFile(filename string, inputFlags cmd.InputFlags, outputChannel chan searchResult) {

	var mu sync.Mutex
	var wg sync.WaitGroup
	fileInfo, err := os.Stat(filename)
	if err != nil {
		outputChannel <- searchResult{
			fileName: filename,
			err:      errors.New("wc: " + filename + " open: " + err.Error()),
		}

		return
	}
	if fileInfo.IsDir() {

		filesInDirectory := walkDirectory(filename)

		if len(filesInDirectory) > 1 {
			printTotal = true
		}

		for _, file := range filesInDirectory {
			wg.Add(1)
			go func(filePath string) {
				serachFile(filePath, inputFlags, outputChannel)
				wg.Done()
			}(file)

		}
		wg.Wait()
		return
	}

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		outputChannel <- searchResult{
			fileName: filename,
			err:      errors.New("wc: " + filename + " open: " + err.Error()),
		}

		return
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)

	lineCount := 0
	totalWordCount := 0
	totalCharCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		wg.Add(1)
		go func() {
			wordCount, charCount := getCounts(line)

			mu.Lock()
			totalWordCount += wordCount
			totalCharCount += charCount
			lineCount++
			mu.Unlock()

			wg.Done()

		}()
	}

	wg.Wait()
	overallCharacterCount += totalCharCount
	overallLineCount += lineCount
	overallWordCount += totalWordCount

	if err := scanner.Err(); err != nil {
		outputChannel <- searchResult{
			fileName: filename,
			err:      errors.New("wc: " + filename + " read: " + err.Error()),
		}

		return
	}
	outputChannel <- searchResult{
		wordCount:      totalWordCount,
		characterCount: totalCharCount,
		lineCount:      lineCount,
		fileName:       filename,
		inputFlags:     inputFlags,
	}
}

// searchUserInput is responsible to get required counts from user provided inputs

func searchUserInput(inputString string, inputFlags cmd.InputFlags, outputChannel chan searchResult) {

	userInputLines := strings.Split(inputString, "\n")

	lineCount := len(userInputLines)
	if lineCount > 0 && userInputLines[lineCount-1] == "" {
		lineCount--
		userInputLines = userInputLines[:len(userInputLines)-1]
	}
	totalWordCount := 0
	totalCharCount := 0

	for _, line := range userInputLines {
		wc, cc := getCounts(line)
		totalWordCount += wc
		totalCharCount += cc
	}

	outputChannel <- searchResult{
		wordCount:      totalWordCount,
		lineCount:      lineCount,
		characterCount: totalCharCount,
		inputFlags:     inputFlags,
	}

	overallCharacterCount += totalCharCount
	overallLineCount += lineCount
	overallWordCount += totalWordCount

}

// getCounts returns wordcount and character count for a given string(line)
func getCounts(line string) (int, int) {

	wordCount, charCount := 0, 0

	wordCount = len(strings.Fields(line))

	charCount = len(line) - strings.Count(line, "\n")

	return wordCount, charCount
}

// handleOutput keeps running in an goroutine and keeps printing result of file as they are pushed to outputChan
// doneChannel is used to indicate that all files are done and now handleOutput can exit

func handleOutput(outputChan <-chan searchResult, doneChannel <-chan bool, wg *sync.WaitGroup) {
	for {
		select {
		case result := <-outputChan:
			printOutput(result)

		case done := <-doneChannel:
			if done {
				wg.Done()
				return
			}
		}
	}

}

// printOutput is responsible to print message to stdoutput in a formatted way. It decides what parameter to include in output message based on inputFlags
func printOutput(searchResult searchResult) {

	if searchResult.err != nil {

		fmt.Println(searchResult.err.Error())
	} else {

		message := ""

		if searchResult.inputFlags.GetLineCount {
			message += fmt.Sprintf("%8d", searchResult.lineCount)

		}
		if searchResult.inputFlags.GetWordCount {
			message += fmt.Sprintf("%8d", searchResult.wordCount)

		}
		if searchResult.inputFlags.GetCharacterCount {
			message += fmt.Sprintf("%8d", searchResult.characterCount)
		}
		if searchResult.fileName != "" {
			message += " " + searchResult.fileName
		}

		fmt.Println(message)

	}

}

// walkDirectory recursively finds all the files in a given directroy
func walkDirectory(dirPath string) []string {
	filePathList := []string{}

	err := filepath.Walk(dirPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {

				filePathList = append(filePathList, path)
			}
			return nil
		})

	if err != nil {
		log.Fatal(err.Error())

	}

	return filePathList
}
