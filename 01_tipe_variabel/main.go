package main

import "fmt"

func main() {
	// konstanta tidak dapat diubah
	const phi float64 = 3.14
	// error jika berusaha reassign constant
	// phi = 22 / 7

	// Deklarasi variable tanpa initial value
	var username string

	// Deklarasi variable dengan initial value
	var email string = "myemail@example.com"

	// Deklarasi variable dengan initial value tanpa definisi tipe data
	var address = "Work Street"

	/* Fitur shorthand untuk deklarasi
	variable dari golang dan tanpa perlu menentukan tipe data, akan otomatis terdeteksi */
	fullname := "danil syah"

	// Reassign data dari variable `username`
	username = "haykal"

	fmt.Printf("Exercise:\n%s\n%s\n%s\n%s\n%.2f", username, email, address, fullname, phi)

}
