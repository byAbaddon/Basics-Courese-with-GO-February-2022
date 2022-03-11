package main

import "fmt"

func main() {
	var floor, rooms int
	fmt.Scan(&floor, &rooms)

	var roomList []string

	for i := 1; i <= floor; i++ {
		for j := rooms; j > 0; j-- {
			if i == floor {
				roomList = append(roomList, fmt.Sprintf("L%d%d", i, j-1))
				continue
			}
			if i%2 != 0 {
				roomList = append(roomList, fmt.Sprintf("A%d%d", i, j-1))
			} else {
				roomList = append(roomList, fmt.Sprintf("O%d%d", i, j-1))
			}
		}
	}

	for i := 0; i < floor; i++ {
		var template string
		for j := 0; j < rooms; j++ {
			slice := roomList[len(roomList)-1]    // Copy last element
			roomList[len(roomList)-1] = ""        // Erase last element (write zero value).
			roomList = roomList[:len(roomList)-1] // Truncate slice.
			template += string(slice) + " "
		}

		fmt.Println(template)
	}

}

/*
name   :06 building
input  : 6 4
output :
L60 L61 L62 L63
A50 A51 A52 A53
O40 O41 O42 O43
A30 A31 A32 A33
O20 O21 O22 O23
A10 A11 A12 A13
*/
