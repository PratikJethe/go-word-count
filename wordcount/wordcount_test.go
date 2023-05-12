package wordcount

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetCounts(t *testing.T) {

	testCases := []struct {
		description       string
		line              string
		expectedCharCount int
		expectedWordCount int
	}{
		{
			description:       "test case for character and word cout",
			line:              "hello world",
			expectedCharCount: 11,
			expectedWordCount: 2,
		},
	}

	for _, testCase := range testCases {
		wc, cc := getCounts(testCase.line)
		fmt.Println(wc, cc)
		if wc != testCase.expectedWordCount || cc != testCase.expectedCharCount {
			t.Fatal("Test Failed: " + testCase.description)
		}
	}
}

func TestWalkDirectory(t *testing.T) {
	testCases := []struct {
		description    string
		dirPath        string
		expectedresult []string
	}{
		{
			description:    "test case to walk directory",
			dirPath:        "..\\test-data",
			expectedresult: []string{"..\\test-data\\All's Well That Ends Well.txt", "..\\test-data\\Antony and Cleopatra.txt", "..\\test-data\\As You Like It.txt", "..\\test-data\\Comedy of Errors.txt", "..\\test-data\\Coriolanus.txt", "..\\test-data\\Cymbeline.txt", "..\\test-data\\Hamlet.txt", "..\\test-data\\Henry IV, part 1.txt", "..\\test-data\\Henry IV, part 2.txt", "..\\test-data\\Henry V.txt", "..\\test-data\\Henry VI, part 1.txt", "..\\test-data\\Henry VI, part 2.txt", "..\\test-data\\Henry VI, part 3.txt", "..\\test-data\\Henry VIII.txt", "..\\test-data\\Julius Caesar.txt", "..\\test-data\\King John.txt", "..\\test-data\\King Lear.txt", "..\\test-data\\Love's Labour's Lost.txt", "..\\test-data\\Macbeth.txt", "..\\test-data\\Measure for Measure.txt", "..\\test-data\\Merchant of Venice.txt", "..\\test-data\\Merry Wives of Windsor.txt", "..\\test-data\\Midsummer Night's Dream.txt", "..\\test-data\\Much Ado About Nothing.txt", "..\\test-data\\Othello.txt", "..\\test-data\\Pericles.txt", "..\\test-data\\Richard II.txt", "..\\test-data\\Richard III.txt", "..\\test-data\\Romeo and Juliet.txt", "..\\test-data\\Taming of the Shrew.txt", "..\\test-data\\The Tempest.txt", "..\\test-data\\Timon of Athens.txt", "..\\test-data\\Titus Andronicus.txt", "..\\test-data\\Troiles and Cressida.txt", "..\\test-data\\Twelfth Night.txt", "..\\test-data\\Two Gentlemen of Verona.txt","..\\test-data\\protected.txt"},
		},
	}

	for _, testCase := range testCases {
		files := walkDirectory(testCase.dirPath)
		fmt.Println(files)
		fmt.Println(testCase.expectedresult)
		if !reflect.DeepEqual(files, testCase.expectedresult) {
			t.Fatal("Test Failed: " + testCase.description)
		}

	}
}


