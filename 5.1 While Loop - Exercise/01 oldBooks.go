package main

import "fmt"
import "os"


func main() {
	var book string
	fmt.Scanln(&book)


	var countBooks = 0

	for {
		var currentBook string
	fmt.Scanln(&currentBook)
	

		if currentBook == "NoMoreBooks" {
			break
		}
		if book == currentBook {
			fmt.Println("You checked", countBooks, "books and found it.")
			os.Exit(0)
		}
		countBooks++

	}

	fmt.Println("The book you search is not here!\nYou checked", countBooks, "books.")
}

/*
name   :oldBooks
input  :
Troy
Stronger
LifeStyle
Troy
output :You checked 2 books and found it.
*/
