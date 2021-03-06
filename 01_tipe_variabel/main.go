package main

import (
	"fmt"
	"reflect"
)


func main() {
	// konstanta tidak dapat diubah
	const phi = 3.14
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
	old := 26

	// Reassign data dari variable `username`
	username = "haykal"

	fmt.Println("Hello\tDanil\tSyah")
	fmt.Printf("Exercise:\n%s\n%s\n%s\n%s\n%.2f\n%d", username, email, address, fullname, phi,old)

	// inisialisasi multiple variable
	var message string = "hello world"
	name, umur, isMarried := "danil", 26, true
	fmt.Printf("Person :\n%s\n%s\n%d\n%t\n",message, name, umur, isMarried)

	// black identifier
	var a, _ string = "stored in a","stored in b"
	fmt.Println(a)

	// display tipe data
	fmt.Printf("%T\n",a)
	fmt.Println(reflect.TypeOf(name))
	fmt.Println(reflect.TypeOf(umur))
	fmt.Println(reflect.TypeOf(isMarried))
	fmt.Println(reflect.TypeOf(phi))

}

// nilai default tipe data
// boolean = false
// integer = 0
// float = 0.0
// string = ""
// pointer,functions,interfaces,slices,channels,maps = nil
