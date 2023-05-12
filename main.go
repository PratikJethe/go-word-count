package main

import (
	"github.com/pratikjethe/go-word-count/cmd"
	"github.com/pratikjethe/go-word-count/wordcount"
)

func main() {

	inputFlags := cmd.GetInput()
	wordcount.StartSearch(inputFlags)
}
