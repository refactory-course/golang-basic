package main

import "fmt"

// function tanpa argument parameter
func hello() {
	fmt.Println("hello world")
}

// function dengan parameter
func cetakNama(nama string, alamat string, umur int, merried bool){
	fmt.Println(nama)
	fmt.Println(alamat)
	fmt.Println(umur)

	if merried{
		fmt.Println("menikah")
	}else{
		fmt.Println("single")
	}
}


// function dengan nilai kembali
func tambah(bil1, bil2 int) (string,int){
	operator := "tambah"
	return operator, bil1 + bil2
}

func gretting(first, last string) string{
	return "hai" + first + last
}

// function variadic
func average(nilai ... float32)float32{
	var total float32 = 0.0
	for _, value := range nilai{
		total += value
	}
	return total / float32(len(nilai))
}

// fungsi sebagai return value
func generatorBilanganGenap() func() int{
	i := 0
	return func() int{
		i += 2
		return i
	}
}

func main() {
	hello()
	cetakNama("danil","tanjung uban",26, true)
	operator,result := tambah(4, 4)
	fmt.Println(operator, result)

	intro := gretting("danil","syah")
	fmt.Println(intro)

	rata := average(7.5, 8.4, 7.7, 9.5, 8.3, 7.6)
	fmt.Println(rata)

	// dengan menggunakan slice
	data := []float32{10,21,23,12,11,18}
	n2 := average(data ...)
	fmt.Println(n2)

	fmt.Println()

	// function closure
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println()
	bilanganGenap := generatorBilanganGenap()
	fmt.Println(bilanganGenap())
	fmt.Println(bilanganGenap())
	fmt.Println(bilanganGenap())
}
