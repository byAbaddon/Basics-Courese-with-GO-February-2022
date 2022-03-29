package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
)

func main() {
	var charCheck = regexp.MustCompile(`[aeiouyAEIOUY]`)
	var sumWordChar, totalSum float64
	var powerWord string
	
  for {
		inputReader := bufio.NewReader(os.Stdin)
		text, _ := inputReader.ReadString('\n')
		word := text[:len(text)-1]

		if word == "End of words" {
			fmt.Printf("The most powerful word is %v - %v", powerWord, totalSum)
			break
		}

		for _, char := range word {
			sumWordChar += float64(int(char))
		}

		if charCheck.MatchString(word[:1]) {
			sumWordChar *= float64(len(word))
		} else {
			sumWordChar = math.Round(float64(sumWordChar) / float64(len(word)))
		}
    
    if totalSum < sumWordChar {
      totalSum = sumWordChar
      powerWord = word
    }

    sumWordChar= 0


	}
}

/*
name   :6. TheMostPowerfulWord
input  :
The
Most
Powerful
Word
Is
Experience
End of words

output :The most powerful word is Experience - 10320
*/

// function theMostPowerfulWord(arg) {
//   let sumWordChar = ''
//   let totalSum = 0
//   let powerWord = {}

//   for (let i = 0; i < arg.length; i++) {
//     if (arg[i] === 'End of words') {
//       break
//     }
//     word = arg[i]
//     sumWordChar = word.split('').map(el => el.charCodeAt(0)).reduce((a, b) => +a + +b)
//     wordFirstChar = word.slice(0, 1)

//     if (/[aeiouyAEIOUY]/gm.test(wordFirstChar)) {
//       totalSum = sumWordChar * word.length
//     } else {
//       totalSum = sumWordChar / word.length
//     }

//     powerWord[word] = Math.round(totalSum)
//     totalSum = 0
//   }

//  const mostStrongWord =  Object.keys(powerWord).reduce((a , b) => powerWord[a] > powerWord[b] ? a : b )
//  console.log(`The most powerful word is ${mostStrongWord} - ${powerWord[mostStrongWord]}`);
// }

// //theMostPowerfulWord(['The', 'Most', 'Powerful', 'Word', 'Is', 'Experience', 'End of words', ])
// //theMostPowerfulWord(["But", "Some", "People", "Say", "It's", "LOVE", "End of words", ])
