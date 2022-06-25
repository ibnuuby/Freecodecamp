package main

import "fmt"

func main() {
	/*
		fallthrough digunakan untuk memaksa proses pengecekan diteruskan ke satu 'case' selanjutnya
		dengan "tanpa menghiraukan nilai kondisinya"
		jadi satu 'case' dilakukan pengecekan selanjutnya tersebut dianggap selalu benar "meskipun aslinya adalah salah"

	*/
	var point = 6

	switch {
	case point == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("awesome")
		fallthrough
	case point < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}

	}

}
