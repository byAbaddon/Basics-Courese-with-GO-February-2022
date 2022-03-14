package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	inputReader := bufio.NewReader(os.Stdin)

	text, _ := inputReader.ReadString('\n')
	filmName := text[:len(text)-1]

	textTwo, _ := inputReader.ReadString('\n')
	salonType := textTwo[:len(textTwo)-1]

	var ticketsCount float64
	fmt.Scan(&ticketsCount)

	tariffDict := map[string]map[string]float64{
		"A Star Is Born":    {"normal": 7.50, "luxury": 10.50, "ultra luxury": 13.50},
		"Bohemian Rhapsody": {"normal": 7.35, "luxury": 9.45, "ultra luxury": 12.75},
		"Green Book":        {"normal": 8.15, "luxury": 10.25, "ultra luxury": 13.25},
		"The Favourite":     {"normal": 8.75, "luxury": 11.55, "ultra luxury": 13.95},
	}

	fmt.Printf("%v -> %.2f lv.", filmName, tariffDict[filmName][salonType]*ticketsCount)
}

/*
name   :03. Oscars week in cinema
input  :
A Star Is Born
luxury
42
output :A Star Is Born -> 441.00 lv.
*/

// function oscarsWeekInCinema(...arg) {
//     let [filmName, salonType, ticketsCount] = [...arg]

//     const tariffDict = {
//         "A Star Is Born": {"normal" :7.50, "luxury" : 10.50, "ultra luxury" : 13.50},
//         "Bohemian Rhapsody": {"normal" :7.35, "luxury" : 9.45, "ultra luxury" : 12.75},
//         "Green Book": {"normal" : 8.15, "luxury" : 10.25, "ultra luxury" : 13.25},
//         "The Favourite": {"normal" : 8.75, "luxury" : 11.55, "ultra luxury" : 13.95 },
//     }

//     return `${filmName} -> ${(tariffDict[filmName][salonType] * (ticketsCount)).toFixed(2)} lv.`
// }

// //console.log(oscarsWeekInCinema("The Favourite", "ultra luxury", 34))
