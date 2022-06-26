package main

import "fmt"

func main() {
	/*
		Fungsi "Copy()" digunakan untuk mengcopy element slice pada "src (parameter ke 2)" ke "dst" (parameter pertama)

	*/

	dst := make([]string, 3)
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	n := copy(dst, src)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(n)
}
