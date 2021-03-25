package main

import "fmt"

func main() {
	/*map adalah kumpulan data yang tidak berurutan dengan sifat reference dan
	memiliki tipe key: value, dengan syarat key harus bersifat unik dan tidak dapat berupa slice / map*/

	//map deklarasi singkat
	salam := map[string]string{
		"danil":  "hai danil",
		"haykal": "hai haykal",
		"fika":   "hai fika",
	}

	fmt.Println(salam)
	fmt.Println(salam["danil"])

	// menambahkan data ke dalam map
	salam["ahmed"] = "hai ahmad"
	fmt.Println(salam)

	// mendeklarasikan map 2
	var bulan = map[int]string{}
	bulan[1] = "januari"
	bulan[2] = "februari"
	bulan[3] = "maret"

	fmt.Println(bulan)

	age:= new([]int) //mengembalikan pointer
	fmt.Println(age)
	fmt.Println(*age)

	bday := make([]int, 10, 20) // make hanya dipergunakan untuk slice, map dan channel
	fmt.Println(bday)

	var names1 = make(map[int]string)
	names1[1] = "jan"
	names1[2] = "februari"
	names1[3] = "maret"

	fmt.Println(names1)
	fmt.Println(len(names1))

	for key, value := range names1{
		fmt.Println("data ke", key, " : ", value)
	}

	m := map[string]string{
		"sapi":"moooo",
	}
	changeMap(m)
	fmt.Println(m)

}

func changeMap(m map[string]string){
	m["kucing"] = "miaw"
}
