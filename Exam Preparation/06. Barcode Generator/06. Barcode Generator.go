package main

import "fmt"

func main() {
	var start, end string
	fmt.Scan(&start, &end)

	var collection string

	for i := start[0]; i <= end[0]; i++ {
		for j := start[1]; j <= end[1]; j++ {
			for k := start[2]; k <= end[2]; k++ {
				for m := start[3]; m <= end[3]; m++ {
					if int(i) % 2 != 0 && int(j) % 2 != 0 && int(k) % 2 != 0&& int(m) % 2 != 0 {
						collection += string(i) + string(j) + string(k) + string(m) + " " 
					}

				}
			}
		}
	}

	fmt.Println(collection)
}

/*
name   :06. Barcode Generator
input  :2345 6789 
output :
3355 3357 3359 3375 3377 3379 3555 3557
3559 3575 3577 3579 3755 3757 3759 3775
3777 3779 5355 5357 5359 5375 5377 5379
5555 5557 5559 5575 5577 5579 5755 5757
5759 5775 5777 5779
*/