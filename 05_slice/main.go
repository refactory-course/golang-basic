package main

import "fmt"

func main() {

	// Inisialisasi variable dengan tipe data slice
	var names = []string{"danil", "haykal", "fika","naruto","antman"}

	// deklarasi variable tipe slice menggunakan fungsi make() dapat memberikan batasan ukuran variable
	var names2 = make([]string, 2) // memiliki arti buat slice dengan tipe data string maksimal 2 index

	names2[0] = "valentino"
	names2[1] = "pedrosa"

	names2 = append(names2, "Black_Panther")
	names2 = append(names2, "Captain_America")
	names2 = append(names2, "Dora")
	names2 = append(names2, "Ironman")
	names2 = append(names2, names...)

	fmt.Println(names, names2)
	fmt.Println(len(names2))


	var names3 = names2[1:3]
	fmt.Println(names3)

	fmt.Println(len(names2), cap(names2))
	fmt.Println(len(names3), cap(names3))
}
