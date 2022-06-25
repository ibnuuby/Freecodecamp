package main

import "fmt"

func main() {
	/*
		Seleksi kondisi bersarang adalah seleksi kondisi, yang berada dalam seleksi kondisi, yang mungkin
		juga berada dalam seleksi kondisi dan seterusnya. seleksi kondisi besarang bisa dilakukan pada "if else, switch
		ataupun kombinasi keduanya"
	*/

	var point = 10
	if point > 7 {
		switch point {
		case 10:
			fmt.Println("perfect")
		default:
			fmt.Println("nice")
		}
	} else {
		if point == 5 {
			fmt.Println("not bad")
		} else if point == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if point == 0 {
				fmt.Println("try harder")

			}
		}
	}
}
