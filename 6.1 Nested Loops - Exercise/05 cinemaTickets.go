package main

import "fmt"

func main() {
	var student, standard, kid float64

	for {
		var movie string
		var seats float64
		fmt.Scan(&movie,&seats )
    
		tickets := 0.0
		chairsType:= ""
		
		for seats > tickets && chairsType != "End" { 
			
			fmt.Scan(&chairsType)


			switch chairsType {
			case "student":
				student++
			case "standard":
				standard++
			case "kid":
				kid++
			}

			tickets++
		}

		fmt.Printf("%s - %.2f%% full.\n", movie, tickets* 100 / seats)
    
		fmt.Scan(&movie)

		if movie == "Finish" {
			break
		}

	}

	totalTickets := student + standard + kid
	fmt.Println("Total tickets:", totalTickets)
	fmt.Printf("%.2f%% student tickets.\n", student*100/totalTickets)
	fmt.Printf("%.2f%% standard tickets.\n", standard*100/totalTickets)
	fmt.Printf("%.2f%% kids tickets.\n", kid*100/totalTickets)

}





// package main

// import "fmt"

// func main() {
// 	var student, standard, kid float64

// 	var movie string
// 	fmt.Scan(&movie)

// 	for {
// 		var seats float64
// 		fmt.Scan(&seats)
// 		tickets := 0.0

// 		for i := 0; i < int(seats); i++ {
// 			var chairsType string
// 			fmt.Scan(&chairsType)

// 			if chairsType == "End"{
// 				break
// 			}

// 			switch chairsType {
// 			case "student":
// 				student++
// 			case "standard":
// 				standard++
// 			case "kid":
// 				kid++
// 			}

// 			tickets++
// 		}

// 		fmt.Printf("%s - %.2f%% full.\n", movie, tickets* 100 / seats)
    
// 		fmt.Scan(&movie)

// 		if movie == "Finish" {
// 			break
// 		}

// 	}

// 	totalTickets := student + standard + kid
// 	fmt.Println("Total tickets:", totalTickets)
// 	fmt.Printf("%.2f%% student tickets.\n", student*100/totalTickets)
// 	fmt.Printf("%.2f%% standard tickets.\n", standard*100/totalTickets)
// 	fmt.Printf("%.2f%% kids tickets.\n", kid*100/totalTickets)

// }

/*
name   :cinemaTickets

input  :
Taxi
10
standard
kid
student
student
standard
standard
End
ScaryMovie
6
student
student
student
student
student
student
Finish

output :
Taxi - 60.00% full.
ScaryMovie - 100.00% full.
Total tickets: 12
66.67% student tickets.
25.00% standard tickets.
8.33% kids tickets.

*/
