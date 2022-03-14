package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var voucher int
	fmt.Scan(&voucher)

	var countTickets, countPurchase int

	for {
		inputReader := bufio.NewReader(os.Stdin)
		text, _ := inputReader.ReadString('\n')
		purchase := text[:len(text)-1]

		if purchase == "End" {
			break
		}

		currentSum := 0

		if len(purchase) > 8 {
			currentSum += int(purchase[0:1][0] + purchase[1:2][0])

			if voucher >= currentSum {
				countTickets++
				voucher -= currentSum
			}
		} else {
			currentSum += int(purchase[0:1][0])
			if voucher >= currentSum {
				countPurchase++
				voucher -= currentSum
			} else {
				break
			}
		}
	}

	fmt.Printf("%v\n%v", countTickets, countPurchase)
}

/*
name   :04. Cinema Voucher
input  :
300
Captain Marvel
popcorn
Pepsi

output :
1
1
*/
