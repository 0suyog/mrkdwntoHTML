package main

import (
	"fmt"
	"log"
	"os"

	"github.com/0suyog/mrkdwntoHTML/tokenizer"
)

func main() {

	fileData, err := os.ReadFile("./markdown.md")
	if err != nil {
		log.Fatal(err.Error())
	}

	plain_markdown := string(fileData)

	tokens := tokenizer.Tokenize(plain_markdown)

	fmt.Println(tokens.String())

}
