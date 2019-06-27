//this is NOT part of the bot. It's used to generate Markov chains from given text files. Just a simple utility.

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/MYKatz/gojam"
)

func main() {
	txtFilename := flag.String("file", "", "a .txt file in paragraph form")
	sep := flag.String("sep", "", "what separates each 'message'. 'punc' for punctuation (as in a book), 'newline' for new lines (songs, poems, etc) ")
	flag.Parse()

	fmt.Println(*sep)

	file, err := os.Open(*txtFilename)
	mark := gojam.NewMarkov(1, " ")

	if err != nil {
		fmt.Printf("Can't open the txt file: %s\n", *txtFilename)
		os.Exit(1)
	}

	defer file.Close()
	txt, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Can't read the txt file: %s\n", *txtFilename)
		os.Exit(1)
	}

	var examples []string
	s := strings.Replace(string(txt), "\n\n", "\n", -1)

	switch string(*sep) {
	case "punc":
		examples = strings.Split(s, ".")
	case "newline":
		examples = strings.Split(s, "\n")
	default:
		fmt.Printf("Invalid separator: %s\n", *sep)
		os.Exit(1)
	}

	fmt.Println(len(examples))

	for i := 0; i < len(examples); i++ {
		mark.TrainOnExample(examples[i])
	}

	fmt.Printf("Training Complete: \n \n")

	for i := 0; i < 10; i++ {
		fmt.Printf("%d \n \n \n", i)
		fmt.Println(mark.GenerateExample())
	}

	ioutil.WriteFile("output.txt", mark.ToJSON(), 0644)

}
