package cmd

import (
	"flag"
	"io"
	"log"
	"os"
)

type InputFlags struct {
	GetWordCount      bool
	GetLineCount      bool
	GetCharacterCount bool
}

type Input struct {
	FileList  []string
	UserInput string
	InputFlags
}
// GetInput function is responsible to accept cmd line args
// It returns Input struct with appropriate flags given by user
// Incase files are not provided user is prompted to add inpput through stdinput
func GetInput() Input {
	// Define flags
	flagWord := flag.Bool("w", false, "pass -w to get word count")
	flagLines := flag.Bool("l", false, "pass -l to get line count")
	flagCharacter := flag.Bool("c", false, "pass -c to get char count")

	flag.Parse()
	if !*flagWord && !*flagLines && !*flagCharacter {
		*flagWord, *flagLines, *flagCharacter = true, true, true
	}

	args := flag.Args()

	if len(args) == 0 {
		// read user input
		data, err := io.ReadAll(os.Stdin)

		if err != nil {
			log.Fatal(err)
		}

		userInput := string(data)

		input := Input{
			UserInput: userInput,
			InputFlags: InputFlags{
				GetLineCount:      *flagLines,
				GetWordCount:      *flagWord,
				GetCharacterCount: *flagCharacter,
			},
		}
		return input

	} else {

		inputFlags := Input{
			InputFlags: InputFlags{
				GetLineCount:      *flagLines,
				GetWordCount:      *flagWord,
				GetCharacterCount: *flagCharacter,
			},
			FileList: args,
		}

		return inputFlags

	}
}
