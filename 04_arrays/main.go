package main

import "fmt"

func main() {
	// penulisan array dengan menentukan panjang array
	var names [3]string
	names[0] = "Danil"
	names[1] = "Haykal"
	names[2] = "Fika"

	fmt.Println(names[0], names[1], names[2])

	// penulisan array dengan mendeklarasikan nilai nya
	var numbers = [3]int{23, 44, 55}
	fmt.Println(numbers[0], numbers[1], numbers[2])
	fmt.Println("panjang dari array numbers : ", len(numbers))

	// penulisan array dengan panjang nya tidak ditentukan
	var items = [...]string{"indomie", "coca cola", "mie sedap", "kacang"}
	fmt.Println(len(items))
	fmt.Println(items)

	// array 2 dimensi
	var nilai = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nilai)
	fmt.Println(nilai[0][1])
	fmt.Println(nilai[1][2])

	// looping terhadap array
	// menggunakan for loop biasa untuk looping array
	for i := 0; i < len(names); i++ {
		fmt.Printf("Nama %s, index ke %d; \n ", names[i], i)
	}

	// Menggunakan for range untuk looping array
	for index, value := range names {
		fmt.Printf("Nama %s, index ke %d;", value, index)
	}

	fmt.Println(" ")

	// bentuk alternatif for range; kita bisa skip variabel dengan underscore
	for index := range names {
		fmt.Printf("Nama %s, index ke %d; ", names[index], index)
	}

	for _, value := range names {
		fmt.Printf("nama : %s ", value)
	}

}
