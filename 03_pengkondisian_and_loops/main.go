package main

import "fmt"

func halo() []string {
	return nil
}

func main() {

	// if...if else..else

	// warna := "biru"
	if warna := "biru"; warna == "merah" {
		fmt.Println("berhenti")
	} else if warna == "hijau" {
		fmt.Println("jalan")
	} else {
		fmt.Println("lainnya")
	}

	// switch case
	var warnaLampu = "ungu"
	switch warnaLampu {
	case "merah", "ungu", "jingga":
		{
			fmt.Println("Berhenti")
			fmt.Println("Stop")
		}
		fallthrough
	case "hijau":
		fmt.Println("Jalan")
	default:
		fmt.Println("lainnya")
	}

	// penulisan switch cara 2
	lampuStop := "merah"
	switch {
	case lampuStop == "merah":
		fmt.Println("Stop")
	case lampuStop == "hijau":
		fmt.Println("Jalan")
	default:
		fmt.Println("lainnya")
	}

	fmt.Println("--------------------------------------")

	// quiz
	h := halo
	fmt.Println()
	if h == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("bukan nil")
	}

	// perulangan

	fmt.Println("------------------------")

	// for
	for i := 1; i < 10; i++ {
		fmt.Println("Berhitung, ", i)
	}

	// for dengan argumen
	var i = 0
	for i < 5 {
		fmt.Println("Angka", i)
		i++
	}

	fmt.Println("--------------------------------------")

	// for tanpa argumen
	var j = 0
	for {
		fmt.Println("Angka", j)
		j++
		if j == 5 {
			break
		}
	}

	fmt.Println("--------------------------------------")

	// penggunaan break dan continue
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka Genap = ", i)
	}

	// looping label dan nested for loop
perulanganLabel:
	for i := 0; i < 999; i++ {
		for j := 0; j < 3; j++ {
			if i == 2 {
				// perintah `break` dijalankan untuk menghentikan `for` dengan label `perulanganLuar`
				// yang juga berarti perulangan paling awal dan paling luar akan diberhentikan
				break perulanganLabel
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}

	// mencari bilangan prima
	var k, l int
	for k = 2; k <= 100; k++ {
		for l = 2; l <= (k / l); l++ {
			if k%l == 0 {
				break //if factor found , not prime
			}
		}
		if l > (k / l) {
			fmt.Printf("%d is prime\n", k)
		}
	}

}
