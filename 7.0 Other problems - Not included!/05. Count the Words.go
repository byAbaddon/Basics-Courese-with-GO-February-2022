package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	text, _ := inputReader.ReadString('\n')
	words := strings.Fields(text)
	if len(words) > 10 {
		fmt.Println("The message is too long to be send! Has", len(words), "words.")
	} else {
		fmt.Println("The message was send successfully!")
	}

}

/*
name   :countTheWords
input  :This message has exactly eleven words. One more as it's allowed!
output :The message is too long to be send!Has 11 words.
*/
