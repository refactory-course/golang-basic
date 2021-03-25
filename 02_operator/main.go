package main

import "fmt"

func main() {
	// aritmatika
	// logika
	// perbandingan

// aritmatika
	bil1, bil2 := 5, 2
	var tambah int
	var bagi float32
	var kurang int
	var kali int
	var modulus int

	tambah = bil1 + bil2
	fmt.Println("jumlah = ", tambah)

	bagi = float32(bil1) / float32(bil2)
	fmt.Println("bagi : ",bagi)

	kurang = bil1 - bil2
	fmt.Println("kurang : ", kurang)

	kali = bil1 * bil2
	fmt.Println("Kali : ", kali)

	modulus = bil1 % bil2
	fmt.Println("Modulus : ", modulus)

	fmt.Println("-----------------------------------------")

	// operator Logika
	var a bool = true
	var b bool = false

	fmt.Println(a && b)
	fmt.Println(a || b)
	fmt.Println(!a)

	fmt.Println("------------------------------------------")

	// operator Perbandingan
	// == : sama dengan
	// != : tidak sama dengan
	// < : kurang dari
	// <= : kurang dari sama dengan
	// > : lebih besar
	// >= : lebih besar sama dengan

	var nilai1, nilai2 int
	nilai1 = 5
	nilai2 = 3

	println(nilai1 > nilai2)
	println(nilai1 == nilai2)
	println(nilai1 >= nilai2)
	println(nilai1 < nilai2)
	println(nilai1 <= nilai2)
	println(nilai1 != nilai2)
}
