// Ini adalah program all untuk patch belajar bahasa Go, dan sumber referensi dari YT = Progammer Zaman Now
// Link =
// Dan ini saya buat untuk pembelajaran dari saya sendiri dan juga bisa mempermudah teman - teman dalam belajar bahasa Go!
// Jangan lupa subscribe dan juga bisa ikuti Github saya
// Link Github Sendy =

package main

import (
	"fmt"
)

func main() {
	// fmt.Println("satu =", 1)
	// fmt.Println("dua =", 2)
	// fmt.Println("tiga koma lima =", 3.5)
	// fmt.Printf("satu = %d\n", 1)

	fmt.Println("Nama saya adalah", "Sendy Prismana Nurferian")
	fmt.Println(len("Sendy"))
	fmt.Println("Sendy Prismana"[0])
	fmt.Println("Sendy Prismana Nurferian"[1]) // Jika keluar angka semua dari 2 bawah ini karena bentuknya 1 string ke byte dan bisa dilihat hasilnya.

	name := "Sendy Prismana Nurferian"
	fmt.Println(name)
	name = "Sendy Prisma Andarini"
	fmt.Println(name)
	// Sebenarnya var ini bisa digantikan dengan tanda ":=" tapi khusus di deklarasi awal jika diganti ke deklarasi kedua dari var tersebut tidak bisa dan tetap memakai "=". Dab metode ini paling banyak digunakan!

	// Bab deklarasi Multiple Variables //
	// var (
	// 	firstName  = "Sendy"
	// 	middleName = "Prismana"
	// 	lastName   = "Nurferian"
	// )

	// fmt.Println(firstName)
	// fmt.Println(middleName)
	// fmt.Println(lastName)

	// Bab Constant //
	// const (
	// 	firstName string = "Sendy"
	// 	lastName         = "Nurferian"
	// )

	// Akan error karena value di const ini tidak bisa diubah maka seperti contoh dibawah ini akan error
	// firstName = "Sendy Prismana"
	// lastName = "Nurferi"

	// Deklarasi multiple const //
	const (
		firstName string = "Sendy"
		lastName         = "Nurferian"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)

	// Bab konversi Tipe Data 1 //
	var nilai32 int32 = 32767 // maks dari int 16 adalah 32767
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)
	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	// Bab konversi Tipe Data 2 //
	var nama = "Sherlin Nilam"
	var e uint8 = nama[0]
	var eString = string(e)
	fmt.Println(nama)
	fmt.Println(e)
	fmt.Println(eString)

	// Bab Type Declaration //
	type NumKTP string
	var ktpSendy NumKTP = "1234567890"
	fmt.Println(ktpSendy)
	fmt.Println(NumKTP("0987654321"))
	// atau
	var contoh string = "1238937231"
	var contohKTP NumKTP = NumKTP(contoh)
	fmt.Println(ktpSendy)
	fmt.Println(contohKTP)

	// Bab operasi matematika di Go //
	var a = 10
	var b = 10
	var d = 5
	var f = 2
	var c = a + b - d*f
	fmt.Println(c)

	// Bab operasi matematika augmented //
	// di bab ini berisi cheatsheet seperti ini:
	// a = a + 10 => a+= 10; a = a - 10 => a-= 10; a = a * 10 => a*= 10; a = a / 10 => a/= 10; a = a % 10 => a%= 10.
	var yz = 10
	yz += 10
	fmt.Println("hasil perhitungan =", yz)
	yz /= 10
	fmt.Println("hasil perhitungan =", yz)

	// Bab Unary Operator //
	// ++ => a = a + 1; -- a = a - 1; - => negative; + => positive; ! => negasi boolean(boolean kebalikan).
	var yh = 10
	yh++
	fmt.Println("hasil yh=", yh)
	yh--
	fmt.Println("hasil yh=", yh)

	// Bab Operator Perbandingan //
	// ==(sama dengan); !=(tidak sama dengan); >(lebih dari); <(kurang dari); >=(lebih dari sama dengan); <=(kurang dari sama dengan).
	var name1 = "sendy"
	var name2 = "sen"
	var result bool = name1 < name2
	fmt.Println("hasilnya", result)

	// Bab oprasi Boolean //
	// && (Dan); || (Atau); ! (Kebalikan)
	var NilaiAkhir = 95
	var Absensi = 85
	var lulusNilaiAkhir bool = NilaiAkhir > 80
	var lulusAbsensi bool = Absensi > 80
	var lulus bool = lulusNilaiAkhir && lulusAbsensi
	fmt.Println("lulus?", lulus)

	// Bab Index di Array //
	var names [3]string
	names[0] = "Sendy"
	names[1] = "Prismana"
	names[2] = "Nurferian"
	fmt.Println("Index 0 =", names[0])
	fmt.Println("Index 1 =", names[1])
	fmt.Println("Index 2 =", names[2])

	// Membuat Array Langsung di Deklarasi variable
	var values = [3]int{
		90,
		80,
		70,
	}
	fmt.Println("VALUES =", values)
	// Atau bisa di index seperti diatas outputnya
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	// Bab Function Array //
	// Didalam bab ini ada len(array) => U/ mendapatkan panjang array; array[index] => U/ mendapatkan value(data) di posisi index array;
	// array[index] = value => U/ mengubah data di posisi index array; array1 == array2 => U/ membandingkan apakah array1 sama dengan array2.
	var val = [...]int{
		90,
		80,
		95,
		100,
		67,
	}
	fmt.Println(val)
	fmt.Println(len(val))
	val[0] = 100
	fmt.Println(val)

	// Tipe data Slice array //
	// Slice array ini seperti array tapi tidak memiliki panjang yang pasti dan bisa diubah-ubah. Dan berikut ini cara membuat Slice dari Array
	// array[low:high] => Membuat slice dari array dimulai index low sampai ke index sebelum high.
	// array[low:] => Membuat slice dari array dimulai index low sampai ke index terakhir di array.
	// array[high:] => Membuat slice dari array dimulai index 0 sampai ke index sebelum high.
	// array[:] => Membuat slice dari array dimulai index 0 sampai ke index terakhir di array.

	namaa := [...]string{"Sendy", "Prismana", "Nurferian", "Choirul", "Isa", "kembar", "laki"}
	slice1 := namaa[2:7]
	fmt.Println(slice1)
	// fmt.Println(slice1[2])
	// fmt.Println(slice1[3])
	slice2 := namaa[:3] // atau bisa ditulis secara manual menjadi var slice2 []string = namaa[:3], dan begitu juga untuk lainnya! Alasan mengapa slice tidak perlu di tulis pada [], karena slice ini memang tidak perlu membutuhkan pointer seperti array. Dan slice ini memang tidak memiliki panjang yang pasti dan bisa diubah-ubah.
	fmt.Println(slice2)
	slice3 := namaa[3:]
	fmt.Println(slice3)

	// Function Slice //
	// len(slice) => U/ mendapatkan panjang slice; cap(slice) => U/ mendapatkan kapasitas slice; append(slice, data) => U/ menambah data ke dalam slice baru ke posisi terakhir slice, jika kapasitas sudah penuh maka akan membuat array baru; copy(dest, src) => U/ menyalin data dari slice src(source) ke slice dest(destination). make([]TypeData, length, capacity) => U/ membuat slice baru dengan tipe data-TypeData, panjang-length, dan kapasitas-capacity.
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daySlice1 := days[5:] // Sabtu dan Minggu
	fmt.Println(daySlice1)
	fmt.Println(days)
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1)
	fmt.Println(days)
	daySlice2 := append(daySlice1, "Libur Baru")
	fmt.Println(daySlice2)
	daySlice2[0] = "ups"
	fmt.Println(daySlice2)
	fmt.Println(days)

	// Program Make Slice //
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Sendy"
	newSlice[1] = "Prismana"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	// jika kita tambah di newslice dengan array baru, program diatas akan error, seharusnya memakai append. Maka seperti berikut:
	newSlice2 := append(newSlice, "Isa")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2)) // Di capacity tetap karena menggunakan array yang sama. Bisa dibuktikan seperti dibawah ini:
	newSlice2[0] = "Baru"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	// Program Copy Slice //
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	//PERINGATAN! Be Careful saat membuat Array dan Slice. Karena jelas berbeda saat di implementasikan!
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)

	// Tipe data Map //
	// Map ini seperti array tapi indexnya bisa diubah-ubah dan tidak harus berupa angka. Dapat diartikan Map adalah tipe data lain yang berisikan kumpulan data yang sama, namun kita bisa menentukan jenis tipe data index yang digunakan. Map ini tipe data kumpulan key-valuenya yang digunakan bersifat unik dan tidak boleh sama. Berbeda dengan Array dan Slice di Map jumlah data yang dimasukkan boleh sebanyak-banyaknya asal kata kunci nya berbeda, jika kata kunci sama maka data yang terdahulu akan diganti dengan data yang baru.
	person := map[string]string{
		"nama":   "Sendy",
		"alamat": "Kediri",
	}
	fmt.Println(person["nama"])
	fmt.Println(person["alamat"])
	fmt.Println(person)
	//Penulisan biasa dalam array
	// var person map[string]string = map[string]string{}
	// person["name"] = "Sendy"
	// person["alamat"] = "Kediri"

	// Function Map //
	// len(map) => U/ mendapatkan jumlah data di map; map[key] => U/ mendapatkan data di map berdasarkan key; map[key] = value => U/ mengubah data di map berdasarkan key; make(map[TypeKey]TypeValue) => U/ membuat map baru dengan tipe data TypeKey untuk key dan TypeValue untuk value; delete(map, key) => U/ menghapus data di map berdasarkan key.
	book := make(map[string]string)
	book["title"] = "Buku Manis"
	book["author"] = "Sendy Prismana"
	book["SALAH"] = "Lohe"
	fmt.Println(book)
	delete(book, "SALAH")
	fmt.Println(book)

	// Bab If and Else Expression //
	namaku := "Sendy"
	if namaku == "Send" {
		fmt.Println("Hello Sendy")
	} else {
		fmt.Println("Hai, Boleh Kenalan ngga?")
	} //output akan Hai, Boleh Kenalan ngga? karena namaku tidak sama dengan Send, namun jika di namaku sudah sama maka yang keluar Hello Sendy.

	// Bab Else If Expression //
	namaq := "Fulani"
	if namaq == "Fulan" {
		fmt.Println("Hello, Fulan")
	} else if namaq == "Fulana" {
		fmt.Println("Hello, Fulana")
	} else {
		fmt.Println("Hai, Boleh Kenalan ngga?")
	} //jika inputnya sesuai di namaq maka outputnya sesuai dengan parameter di if dan else if dan else jika di namaq tidak sesuai dengan di if dan else if.

	// Bab If dengan Short Statement //
	// Jika biasanya kita harus seperti :
	//length := len(namaq)
	//if length > 5 { fmt.Println("Nama Terlalu Panjang") } else { fmt.Println("Nama Sudah Benar") } maka di Go kita bisa seperti ini:
	if length := len(namaq); length > 5 {
		fmt.Println("Nama Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}

	// Bab Switch Expression //
	Name := "Prismana"
	switch Name {
	case "Sendy":
		fmt.Println("Hello, Sendy!")
	case "Prismana":
		fmt.Println("Hello, Prismana!")
	default:
		fmt.Println("Hai, Boleh Kenalan ngga?")
	}

	// Bab Switch dengan Short Statement //
	// Jika biasanya kita harus seperti :
	//lenght := len(Name)
	//switch lenght > 5 { fmt.Println("Nama Terlalu Panjang") } else { fmt.Println("Nama Sudah Benar") } maka di Go kita bisa seperti ini:
	switch length := len(Name); length > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang!!")
	case false:
		fmt.Println("Nama Sudah Benar!!")
	}

	//Bab Switch tanda kondisi //
	lenght := len(Name)
	switch {
	case lenght > 10:
		fmt.Println("Nama Terlalu Panjang")
	case lenght > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Sudah Benar")
	}

	//Bab For Loop //
	counter := 1
	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}
	fmt.Println("Selesai")

	//Bab For Loop dengan Statement //
	// Ada 2 statement yaitu Init dan Post. Untuk Init statement adalah statement sebelum for loop dieksekusi. Post Statement adalah statement yang selalu dieksekusi di akhir tiap perulangan.
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}
	fmt.Println("Selesai")

	//Bab For Range //
	// For Range digunakan untuk melakukan iterasi terhadap data-data yang kita miliki (data collection). For Range akan mengembalikan 2 data yaitu index dan data di index tersebut, Contoh dari Data Collection adalah Array, Slice, dan Map.
	//Contoh Program sebelum For Range //
	namas := []string{"Sandy", "Prisma", "Abdul"}
	for i := 0; i < len(namas); i++ {
		fmt.Println(namas[i])
	}
	//Contoh Program setelah For Range //
	for index, nama := range namas {
		fmt.Println("Index", index, "=", nama)
	}
	//Contoh Program setelah For Range tanpa index //
	for _, nama := range namas {
		fmt.Println(nama)
	}

	//Bab Break & Continue //
	//Break digunakan untuk menghentikan seluruh perulangan. Continue digunakan untuk menghentikan perulangan yang berjalan dan langsung melanjutkan ke perulangan berikutnya.
	//BREAK //
	for i := 0; i < 10; i++ {
		if i == 9 {
			break
		}
		fmt.Println("Perulangan Break ke-", i)
	}
	//CONTINUE //
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Perulangan Continue ke-", i)
	}

	// Bab FUNCTION //
	// Function adalah blok kode yang sengaja dibuat untuk melakukan tugas tertentu. Function dapat digunakan berulang kali sesuai dengan kebutuhan. Function dapat menerima input dan mengembalikan output.
	// Func sayHello() {
	// 	fmt.Println("Hello")
	// }
	// func main() {
	// 	sayHello()
	// }

	// Bab Function Parameter //
	// Parameter adalah data yang kita butuhkan untuk memproses sebuah function. Kadang kita membutuhkan data dari luar dan biasa disebut Parameter. Dalam function parameter bisa ditambahkan lebih dari 1. Parameter memang tidak wajib, jadi kita bisa membuat function tanpa parameter seperti sebelumnya. Namun jika menambahkan parameter di function maka ketika memanggilnya kita wajib memasukkan data ke parameternya.
	// func sayHelloTo(firstName string, lastName string){
	// 	fmt.Println("Hello", firstName, lastName)
	// }
	// func main() {
	// 	sayHelloTo("Eko", "Khannedy")
	// }

	// Bab Function Return Value //
	// Function ini bisa mengembalikan data. U/ memberitahu bahwa function mengembalikan data, kita harus menuliskan tipe data kembalian dari function tsb. Jika Function kita deklarasikan dgn tipe data pengembalian, maka wajib di dalam function nya kita harus mengembalikan data. U/ mengembalikan data dari function, kita bisa menggunakan kata kunci return diikuti datanya.
	// func getHello(name string) string{
	// 	return "Hello" + name
	// }
	// func main(){
	// 	result := getHello("Eko")
	// 	fmt.Println(result)
	//  fmt.Println(getHello("Eko"))
	// }

	// Bab Returning Multiple Values //
	// Function tidak hanya dapat mengembalikan 1 value, tapi juga bisa multiple value. U/ memberitahu jika function mengembalikan multiple value, kita harus menuliskan semua tipe data return valuenya di function tsb.
	// func getFullName() (string, string) {
	// 	return "Eko", "Khannedy"
	// }
	// func main() {
	// 	firstName, lastName := getFullName()
	// fmt.Println(firstName,lastName)
	// }
	// Dari diatas kita bisa juga menghiraukan Return Value. Multiple Return Value wajib ditangkap semua value-nya. Jika kita ingin menghiraukan salah satu value, kita bisa menggunakan tanda underscore atau garis bawah (_).
	// func getFullName() (string, string) {
	// return "Eko", "Khannedy"
	// }
	// func main() {
	// firstName, _ := getFullName()
	// fmt.Println(firstName)
	// }

	// Bab Named Return Values //
	// Maksud dari bab ini adalah kita bisa mengembalikan data dari function dengan cara langsung menuliskan data yang ingin dikembalikan. Dengan cara ini kita tidak perlu menuliskan return value di functionnya. Atau dapat dikatakan biasanya kita memberi tahu bahwa sebuah funtion mengembalikan value, maka kita hanya mendeklarasikan tipe data return value di function. Namun kita juga bisa membuat variable secara langsung di tipe data return function nya.
	// func getCompleteName() (firstName, middleName, lastName string) {
	// firstName = "Eko"
	// middleName = "Khannedy"
	// lastName = "Fatriansyah"
	// return firstName, middleName, lastName
	// }
	// func main() {
	// firstName, middleName, lastName := getCompleteName()
	// fmt.Println(firstName, middleName, lastName)
	// }

	// Bab Variadic Function //
	// Parameter yang berada di posisi terakhir, memiliki kemampuan dijadikan sebuah varargs. Varargs artinya datanya bisa menerima lebih dari satu input, atau anggap saja semacam array. Untuk membuat variadic function, kita bisa menambahkan tanda elipsis (...) sebelum tipe data parameter terakhir. Jika kita lihat di function main, kita bisa memasukkan data sebanyak-banyaknya ke parameter numbers. Kemudian secara otomatis data yang dimasukkan akan dikonversi menjadi sebuah array. Lalu apa bedanya dengan parameter biasa dan tipe data Array? Jika parameter tipe Array, kita wajib membuat Array terlebih dahulu sebelum mengirimkan ke function. Jika parameter menggunakan varargs, kita bisa langsung mengirim datanya, jika lebih dari satu, cukup gunakan tanda koma.
	// func sumAll(numbers ...int) int {
	// 	total := 0
	// 	for _, value := range numbers {
	// 		total += value
	// 	}
	// 	return total
	// }
	// func main() {
	// 	total := sumAll(10, 20, 30, 40, 50)
	// 	fmt.Println(total)
	// }

	// Bab Slice Parameter //
	// Dimana ada kasus kita harus menggunakan Varadic Function, namun memiliki variable berupa Slice. Nah kita bisa menjadikan Slice sebagai varargs parameter. Sebagai contoh kita gunakan program sebelumnya namun terdapat perubahan pada algoritma di function. Maka dapat dituliskan seperti berikut:
	// func main() {
	//	total := sumAll(10, 20, 30, 40, 50)
	//	fmt.Println(total)
	//	numbers := []int{10, 20, 20, 30, 40}
	//	total = sumAll(numbers...)
	//	fmt.Println(total)
	// }

	// Bab Function as Value //
	// Function adalah first class citizen. Function juga merupakan tipe data, dan bisa disimpan di dalam variable.
	// func getGoodBye(name string) string {
	//	return "Good Bye " + name
	// }
	// func main() {
	// 	sayGoodBye := getGoodBye
	// 	fmt.Println(sayGoodBye("Sendy"))
	// }

	// Bab Function as Parameter //
	// Function juga bisa kita gunakan sebagai parameter untuk function lain. Hal ini biasa disebut dengan callback. Contoh kasusnya adalah kita memiliki function untuk melakukan filter data. Function filter ini menerima 2 parameter, yaitu data berupa Slice dan function yang digunakan untuk melakukan filter. Function filter ini akan mengembalikan data Slice yang sudah difilter. Function tidak hanya bisa kita simpan di dalam variable sebagai value.
	// Function as Parameter (1) //
	// func sayHelloWithFilter(name string, filter func(string) string) {
	// 	filteredName := filter(name)
	//	fmt.Printf("Hello %s\n", filter(name))
	// }
	// func spamFilter(name string) string {
	// 	if name == "Kuda" {
	// return "..."
	// 	} else {
	// 		return name
	// 	}
	// }
	// Function as Parameter (2) = cara memanggil di (1) //
	// func main() {
	//	sayHelloWithFilter("Eko",spamFilter)
	//	filter := spamFilter
	//	sayHelloWithFilter("Kuda",spamFilter)
	// }

	// Bab Function Type Declarations //
	// Jika function terlalu panjang biasanya agak ribet U/ menuliskannya di dalam parameter. Type Declarations juga bisa digunakan U/ membuat alias function sehingga akan mempermudah kita menggunakan function sebagai parameter.
	// type Filter func(string) string
	// func sayHelloWithFilter(name string, filter Filter) {
	// 	fmt.println("Hello", filter(name))
	// }
	// func spamFilter(name string) string {
	//	if name == "Kuda" {
	// return "..."
	//		} else {
	//			return name
	// 		}
	// }

	//  Bab Anonymous Function //
	// Sebelumnya setiap membuat function, kita akan selalu memberikan sebuah nama pada function tersebut. Namun kadang ada kalanya lebih mudah membuat function secara langsung di variable atau parameter tanpa harus membuat nama functionnya. Hal ini biasa disebut dengan Anonymous Function. Anonymous Function adalah function tanpa nama. Anonymous Function biasanya dibuat ketika kita ingin membuat function secara langsung di variable atau parameter.
	// Anonymous Function (1) //
	// type Blacklist func(string) bool
	// func registerUser(name string, blacklist Blacklist) {
	// 	if blacklist(name) {
	// 		fmt.Println("You are Blocked", name)
	//	} else {
	// 		fmt.Println("Welcome", name)
	//	}
	// }
	// Anonymous Function (2) = cara memanggil di (1) //
	// func main() {
	// 	blacklist := func(name string) bool {
	// 		return name == "anjing"
	//	}
	// 	registerUser("Sendy", blacklist)
	// 	registerUser("Anjing", func(name string) bool {
	//	return name == "Anjing"
	//	})
	// }

	// Bab Recursive Function //
	// Recursive Function adalah function yang memanggil dirinya sendiri. Recursive Function biasanya digunakan untuk perulangan yang jumlahnya tidak bisa ditentukan. Contoh kasus yang biasanya menggunakan Recursive Function adalah factorial. Dan kadang dalam pekerjaan, kita sering menemui kasus dimana menggunakan Recursive Function lebih mudah dan lebih cepat daripada menggunakan perulangan biasa.
	// func factorialLoop(value int) int{
	// 	result := 1
	// 	for i := value; i > 0; i-- {
	// 		result *= i
	//	}
	// 	return result
	// }
	// func main() {
	// 	fmt.Println(factorialLoop(10))
	// }

	// Dari Program diatas kita bisa mengubahnya menjadi seperti berikut:
	// func factorialRecursive(value int) int{
	// 	result := 1
	// 	for i := value; i > 0; i-- {
	// 		result *= i
	//	}
	// 	return result
	// }
	// func main() {
	//	result := 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1
	//	fmt.Println(result)
	// 	fmt.Println(factorialRecursive(10))
	// 	fmt.Println(factorialLoop(10))
	// }

	// Bab Closure //
	// Closure adalah kemampuan sebuah function berinteraksi dengan data-data sekitarnya dalam scope yang sama. Diharapkan menggunakan fitur ini secara bijak saay kita membuat aplikasi!
	// func main() {
	// 	counter := 0
	// 	increment := func() {
	// 		fmt.Println("Increment")
	// 		counter++
	// 	}
	// 	increment()
	// 	increment()
	// 	fmt.Println(counter)
	// }

	// Bab Defer, Panic, dan Recover //
	// Defer Function adalah Function yang bisa kita jadwalkan U/ dieksekusi setelah sebuah function selesai di eksekusi. Dan akan selalu dieksekusi walaupun terjadi error di function yang dieksekusi. Defer Function sering digunakan U/ clean up code, seperti menutup file setelah file tersebut dibuka.
	// func logging(){
	// 	fmt.Println("Selesai memanggil function")
	// }
	// func runApplication(){
	// 	defer logging()
	// 	fmt.Println("Run Application")
	// }
	// func main() {
	// 	runApplication()
	// }
	// Panic Function adalah Function yang bisa kita gunakan U/ menghentikan program. Panic Function biasanya dipanggil ketika terjadi error pada saat program kita berjalan. Saat panic function dipanggil, program akan terhenti, namun defer function tetap akan dieksekusi.
	// func endApp(){
	// 	fmt.Println("Aplikasi Selesai")
	// }
	// func runApp(error bool){
	// 	defer endApp()
	// 	if error {
	// 		panic("APLIKASI ERROR")
	// 	}
	// }
	// func main() {
	// 	runApp(true)
	// }

	// Recover Function adalah function yang bisa kita gunakan U/ menangkap data panic. Dengan begini program kita tidak akan berhenti secara langsung ketika panic terjadi. Namun tetap akan mengeksekusi defer function. Atau bisa dibilang recover function proses panic akan terhenti, sehingga program tetap bisa berjalan.
	// func runApp(error bool){
	// 	defer endApp()
	// 	if error {
	// 		panic("APLIKASI ERROR")
	// 	}
	// 	message := recover()
	// 	fmt.Println("Terjadi Error bang", message)
	// }
	// func main() {
	// 	runApp(true)
	// }
	// Program yang benar adalah berikut:
	// func endApp(){
	// 	fmt.Println("Aplikasi Selesai")
	//	mesaage := recover()
	// 	fmt.Println("Terjadi Error bang", message)
	// }
	// func runApp(error bool){
	// 	defer endApp()
	// 	if error {
	// 		panic("APLIKASI ERROR")
	// 	}
	// }
	// func main() {
	// 	runApp(true)
	// 	fmt.Println("Menjalankan Aplikasi") // buat uji coba
	// }

	// KOMENTAR ATAU COMMENT => selain memakai // juga bisa memakai /* .... */ dan juga bisa memakai // di awal dan /* .... */ di akhir. Dan juga bisa memakai // di awal dan akhirnya. Dan juga bisa memakai /* .... */ di awal dan akhirnya.

	// Bab Struct //
	// Struct adalah sebuah template data yang digunakan U/ menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan. Struct biasanya representasi data dalam program aplikasi yang kita buat. Contoh kasusnya adalah ketika kita ingin membuat data seorang pelanggan. Kita bisa membuatnya dengan menggunakan struct. Struct biasanya mendefinisikan sebuah tipe data baru, namun sebenarnya kita juga bisa menggunakan tipe data yang sudah ada seperti int, string, dan lainnya. Data di Struct disimpan di dalam field. Field di Struct memiliki tipe data dan nama. Kita bisa mengakses data di dalam field dengan menggunakan dot (.) setelah nama variable struct nya. Sederhananya Struct adalah kumpulan dari field. Struct juga bisa dikatakan template data atau prototype data, tidak bisa digunakan langsung. Namun kita bisa membuat data/object dari struct yang telah kita buat.
	// type Customer struct {
	// 	Name, Address string
	// 	Age           int
	// }
	// func main() {
	// 	var sendy Customer
	// 	fmt.Println(sendy)
	// 	sendy.Name = "Sendy Prismana Nurferian"
	// 	sendy.Address = "Indonesia"
	// 	sendy.Age = 21
	// 	fmt.Println(sendy)
	// 	fmt.Println(sendy.Name)
	// 	fmt.Println(sendy.Address)
	// 	fmt.Println(sendy.Age)
	// }

	// Struct Literals //
	// Struct Literals adalah cara membuat data struct tanpa harus mengisi semua field nya. Kita bisa mengisinya secara sebagian saja. Cara ini sangat berguna ketika struct yang kita buat memiliki banyak field. Kita tidak perlu mengisi semua field, namun hanya field yang kita butuhkan saja.
	// Dari program diatas maka:
	// isa := Customer{
	// 	Name:    "Isa",
	// 	Address: "Indonesia",
	// 	Age:     21,
	// }
	// fmt.Println(isa)
	// fmt.Println(isa.Name)
	// fmt.Println(isa.Address)
	// fmt.Println(isa.Age)

	// Dan bisa seperti:
	// isat := Customer{"Isat", "Indonesia", 22}
	// fmt.Println(isa)

	// Struct Method //
	// Struct adalah tipe data yang seperti tipe data lainnya, bisa digunakan sebagai parameter U/ function. Namun jika ingin menambahkan method ke dalam struct, sehingga seakan-akan sebuah struc memiliki function, maka kita harus menggunakan pointer di parameter functionnya. Dan Method adalah Function.
	// Dari program diatas maka:
	// func (customer Customer) sayHello() {
	// 	fmt.Println("Hello, My Name is", customer.Name)
	// }
	// func main() {
	// 	nilam := Customer{"Nilam"}
	// 	nilam.sayHello()
	// 	sendy.sayHello("Suha")
	// 	isa.sayHello("Ibra")
	// 	isat.sayHello("Tessy")
	// }

	// Bab Interfaces //
	// Merupakan tipe data Abstract, tidak memiliki implementasi langsung. Sebuah interface berisikan definisi-definisi method. Biasanya interface digunakan sebagai kontrak. Artinya jika ada struct yang memiliki method-method sesuai definisi di interface, maka struct tersebut otomatis dianggap sebagai implementasi dari interface tersebut. Biasanya interface digunakan sebagai kontrak. Artinya jika ada struct yang memiliki method-method sesuai definisi di interface, maka struct tersebut otomatis dianggap sebagai implementasi dari interface tersebut. Biasanya interface digunakan sebagai kontrak. Artinya jika ada struct yang memiliki method-method sesuai definisi di interface, maka struct tersebut otomatis dianggap sebagai implementasi dari interface tersebut.
	// type HasName interface {
	// 	GetName() string
	// }
	// func sayHello(hasName HasName) {
	// 	fmt.Println("Hello", hasName.GetName())
	// }

	// Implementasi dari interface (1) //
	// type Person struct {
	// 	Name string
	// }
	// func (person Person) GetName() string {
	// 	return peson.Name
	// }
	// func main() {
	// person := Person{Name: "Sendy"}
	// 		sayHello(person)
	// }

	// Implementasi dari interface (2) //
	// type Animal struct {
	// 	Name string
	// }
	// func (animal Animal) GetName() string {
	// 	return animal.Name
	// }
	// func main() {
	// 	animal := Animal{Name: "Kucing"}
	// 	sayHello(animal)
	// }

	// Bab Interface Kosong //
	// Go-Lang ini bukanlah bahasa pemrograman yang berorientasi dengan objek. Untuk menangani kasus ini kita bisa menggunakan Interface Kosong. Interface Kosong adalah interface yang tidak memiliki deklarasi method satupun. Biasanya interface kosong digunakan sebagai tipe data untuk parameter function yang bisa menerima semua tipe data. Atau bisa dibilang interface kosong adalah induk dari semua tipe data di Go-Lang. Hal ini membuat secara otomatis semua tipe data akan menjadi implementasinya. Interface Kosong juga memiliki type alias bernama any.
	// Contoh penggunaan interface kosong :
	// - fmt.Println(a ...interface[])
	// - fmt.Println(a ...any)
	// - panic(v interface{})
	// - recover() interface{} dan lain-lain
	// Contoh programnya :
	// func Ups() interface{} {
	// 		//return 1
	// 		//return true
	// 		return "Ups"
	// }
	// func main() {
	// 		kosong := Ups()
	// 		fmt.Println(kosong)
	// }

	// Atau di Contoh Program Lain:
	// func Ups() any {
	// 		//return 1
	// 		//return true
	// 		return "Ups"
	// }
	// func main() {
	// 		var kosong any := Ups()
	// 		fmt.Println(kosong)
	// }

	// Bab Nil
	// Dalam bahasa pemograman lain,biject yang belum diinisialisasi maka secara otomatis nilainya adalah Null atau Nil. Dengan Go-Lang kita bisa membuat variable dengan tipe data tertentu, maka secara otomatis akan dibuatkan secara default value-nya. Namun di Go-lang ada data Nil atau data kosong. Nil adalah nilai default dari sebuah tipe data interface. Nil sendiri sebenarnya bukan sebuah keyword, namun sebuah konstanta. Nil sendiri bisa digunakan di beberapa tipe data seperti interface, function, map, slice, pointer dan channel.
	// Contoh Programnya :
	// func Contoh(name string) string {
	// 	if name == "" {
	// 		return nil
	// 	} else {
	// 		return name
	// 	}
	// }
	// func NewMap(name string) map[string]string {
	// 	 if name == "" {
	// 		 return nil
	// } else {
	// 		return map[string]string{
	// 			"name": name,
	// 	   }
	// 	 }
	// }
	// func main() {
	// 		data := Contoh("") //Sesuaikan Contoh atau NewMap()
	// 		if data == nil {
	// 			fmt.Println("Data Kosong")
	// 		} else {
	// 			fmt.Println(data)
	// 	 	}
	// }

	// Bab Type Assertions //
	// Type Assertions adalah kemampuan merubah tipe data menjadi tipe data lain atau yang diinginkan. Biasanya hal ini digunakan ketika kita berinteraksi dengan data kosong (interface{}). Type Assertions juga sering digunakan ketika berinteraksi dengan package-package eksternal yang belum kita ketahui secara pasti tipe datanya. Type Assertions mirip dengan Type Casting di bahasa pemrograman lain. Namun perbedaannya adalah Type Assertions hanya bisa dilakukan kepada tipe data interface{}.
	// Contoh Programnya :
	// func random() interface{} {
	// 	return "OK"
	// }
	// func main() {
	// 	var result interface{} = random() // Atau ditulis result := random()
	// 	var resultString string = result.(string) // Atau ditulis resultString := result.(string)
	// 	fmt.Println(resultString)
	// resultInt := result.(int) // Akan error karena result bukan tipe data int (panic)
	// fmt.Println(resultInt)
	// }

	// Bab Type Assertions dengan menggunakan Switch //
	// Saat salah melakukan Type Assertions, maka program akan panic. Untuk menghindari panic, kita bisa menggunakan switch expression. Switch expression ini mirip seperti switch biasa, namun case-case nya adalah tipe data yang bisa diassign ke variable yang kita inginkan.
	// func main() {
	// 	var result interface{} = random() // Atau ditulis result := random()
	// 	switch value := result.(type) {
	// 	case string:
	// 		fmt.Println("Value", value, "is string")
	// 	case int:
	// 		fmt.Println("Value", value, "is int")
	// 	default:
	// 		fmt.Println("Unknown Type")
	// 	}
	// }

	// Bab Pointer //
	// Pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama, tanpa menduplikasi data yang sudah ada. Biasanya pointer digunakan ketika membuat variabel yang ukurannya besar, misalnya ketika membuat variabel bertipe data array atau slice. Pointer juga digunakan ketika kita ingin sebuah data bisa diubah nilainya di scope yang berbeda. Pointer mirip seperti reference di bahasa pemrograman lain. Pointer di Go-Lang tidak seperti pointer di bahasa pemrograman lain yang bisa melakukan aritmatika pointer. Di Go-Lang kita tidak bisa melakukan aritmatika pointer.
	// Pass by Value //
	// Secara default di Go-Lang semua variable itu di-pass by value. Artinya ketika kita mengirimkan data ke function, maka akan dibuatkan data baru di memory. Jadi jika kita mengubah data di function, data yang asli tidak akan berubah. Hal ini berlaku untuk semua tipe data, baik tipe data primitif maupun tipe data reference. Atau kata lain jika kita mengirim sebuah variable ke dalam function, method atau variable lain, sebenarnya yang dikirim adalah duplikasi value-nya.
	// func main() {
	// 	address1 := Address{"Kediri", "Jawa Timur", "Indonesia"}
	// 	address2 := address1
	// 	address2.City = "Surabaya"
	// 	fmt.Println(address1) // Disini Address tidak berubah
	// 	fmt.Println(address2)
	// }

	// Operator &
	// Operator & digunakan untuk membuat pointer. Operator & juga digunakan untuk mengambil data di memory yang sama dengan pointer yang kita buat. Operator & hanya bisa digunakan di tipe data pointer.
	// func main() {
	// 	address1 := Address{"Kediri", "Jawa Timur", "Indonesia"}
	// 	address2 := &address1
	// jika kita menulis dengan var biasa maka:
	//  var address1 = Address{"Kediri", "Jawa Timur", "Indonesia"}
	// 	var address2 *Address = &address1
	// 	address2.City = "Surabaya"
	// 	fmt.Println(address1) // Disini Address berubah
	// 	fmt.Println(address2)
	// }

	// Operator * (Asterisk)
	// Operator * digunakan untuk membuat pointer. Operator * juga digunakan untuk mengambil data di memory yang sama dengan pointer yang kita buat. Operator * hanya bisa digunakan di tipe data pointer. Saat kita mengubah variable pointer maka yang berubah hanya variable pointer tersebut, bukan variable aslinya. Jika kita ingin mengubah seluruh variablenya, maka kita bisa menggunakan operator * diikuti dengan nama variablenya.
	// type Address struct {
	// 	City, Province, Country string
	// }
	// Kode 1:
	// func main() {
	// 	address1 := Address{"Kediri", "Jawa Timur", "Indonesia"}
	// 	address2 := &address1
	// 	address2.City = "Surabaya"
	// 	address2 = &address{"Jakarta", "DKI", "Indonesia"}
	// 	fmt.Println(address1) // Disini Address berubah
	// 	fmt.Println(address2)
	// }
	// Kode 2:
	// func main() {
	// 	address1 := Address{"Kediri", "Jawa Timur", "Indonesia"}
	// 	address2 := &address1
	// 	address2.City = "Surabaya"
	// 	*address2 = Address{"Jakarta", "DKI", "Indonesia"}
	// 	fmt.Println(address1) // Disini Address berubah
	// 	fmt.Println(address2)
	// }

	// Operator New
	// Operator new digunakan untuk membuat pointer secara otomatis. Kita tidak perlu lagi menggunakan operator & diawal. Operator new akan mengembalikan pointer yang menunjuk ke data kosong dari tipe data pointer yang kita buat. Jadi kita tidak perlu lagi mengubah data kosong tersebut menjadi data yang kita inginkan.
	// type Address struct {
	// 	City, Province, Country string
	// }
	// func main() {
	// 	alamat1 := new(Address)
	// 	alamat2 := alamat1
	// 	alamat2.City = "Kediri"
	// 	fmt.Println(alamat1) // alamat1 berubah
	// 	fmt.Println(alamat2)
	// }
	// Atau bisa ditulis:
	// func main() {
	// var alamat1 *Address = new(Address)
	// var alamat2 *Address = alamat1
	// alamat2.City = "Kediri"
	// fmt.Println(alamat1) // alamat1 berubah
	// fmt.Println(alamat2)
	// }

	// Bab Pointer di Function //
	// Saat membuat function, secara default semua variable di function tersebut di-pass by value. Artinya ketika kita mengirimkan data ke function, maka akan dibuatkan data baru di memory. Jadi jika kita mengubah data di function, data yang asli tidak akan berubah. Hal ini berlaku untuk semua tipe data, baik tipe data primitif maupun tipe data reference. Namun jika kita ingin mengubah data asli, kita bisa menggunakan pointer di function parameter. Jadi jika kita mengubah data di function, data asli juga akan berubah. Namun kadang kita tidak ingin mengubah data asli, namun hanya ingin mengubah data di function saja. Maka kita bisa menggunakan operator * di parameter functionnya.
	// type Address struct {
	// 	City, Province, Country string
	// }
	// func ChangeAddressToIndonesia(address *Address) {
	// 	address.Country = "Indonesia"
	// }
	// func main() {
	// 	address := Address{"Kediri", "Jawa Timur", ""}
	// 	ChangeAddressToIndonesia(address)
	// 	fmt.Println(address) // Tidak berubah
	// }
	// Jika ubah pointer:
	// func main() {
	// 	var address *Address = &Address{}
	// 	ChangeAddressToIndonesia(address)
	// 	fmt.Println(address) // Berubah
	// }

	// Bab Pointer di Method //
	// Saat membuat method, secara default semua variable di method tersebut di-pass by value. Artinya ketika kita mengirimkan data ke method, maka akan dibuatkan data baru di memory. Jadi jika kita mengubah data di method, data yang asli tidak akan berubah. Hal ini berlaku untuk semua tipe data, baik tipe data primitif maupun tipe data reference. Namun jika kita ingin mengubah data asli, kita bisa menggunakan pointer di method parameter. Jadi jika kita mengubah data di method, data asli juga akan berubah. Namun kadang kita tidak ingin mengubah data asli, namun hanya ingin mengubah data di method saja. Maka kita bisa menggunakan operator * di parameter methodnya. Walaupun method akan menempel di struct, tapi sebenarnya data struct yang diakses di dalam method adalah pass by value. Sangat direkomendasikan menggunakan pointer di method, sehingga tidak boros memory karena harus selalu di duplikasi ketika memanggil method.
	// type Man struct {
	// 	Name string
	// }
	// func (man *Man) Married() {
	// man.Name = "Mr." + man.Name
	// }
	// func main() {
	//	eko := Man("Eko")
	// 	eko.Married()
	// 	fmt.Println(eko.Name) // Tidak berubah
	// }

	// Bab Package & Import //
	// Package adalah tempat yang digunakan untuk mengorganisir kode program Go-Lang. Package biasanya digunakan untuk membuat library/library yang bisa digunakan kembali. Package di Go-Lang mirip seperti namespace di bahasa pemrograman lain. Package di Go-Lang dibuat dengan cara membuat folder baru di dalam folder src. Nama folder tersebut adalah nama package yang kita inginkan. Di dalam folder tersebut kita bisa membuat file-file Go-Lang. Nama file tersebut tidak harus sama dengan nama package. Namun biasanya nama file diusahakan sama dengan nama package. Jika kita membuat package, maka package tersebut bisa digunakan di package lain. Namun ada syaratnya, yaitu nama package harus berbeda. Jika kita membuat package, maka package tersebut bisa digunakan di package lain. Namun ada syaratnya, yaitu nama package harus berbeda. Jika kita membuat package, maka package tersebut bisa digunakan di package lain. Namun ada syaratnya, yaitu nama package harus berbeda.
	// Contoh Programnya :
	// package helper
	// func SayHello(name string) {
	// 	fmt.Println("Hello", name)
	// }

	// Import adalah cara kita menggunakan package lain yang sudah dibuat. Untuk menggunakan package lain, kita bisa menggunakan kata kunci import diikuti dengan nama package yang akan kita gunakan. Jika package yang kita gunakan berada di package yang sama, maka kita bisa menggunakan nama package . (titik) untuk menggantikan nama package. Jika kita menggunakan package yang berada di package lain, maka kita harus menggunakan nama package yang benar.
	// Contoh Programnya :
	// package main
	// import (
	// 	"fmt"
	// 	"belajar-golang-dasar/helper"
	// )
	// func main() {
	// 	result := helper.SayHello("Sendy")
	// 	fmt.Println(result)
	// }

	// Bab Acces Modifier //
	// Access Modifier adalah kemampuan sebuah function, method, struct, interface, dan properti diakses dari package lain atau tidak. Di Go-Lang Access Modifier hanya ada 2, yaitu Public dan Private. Public artinya bisa diakses dari package lain, sedangkan Private artinya hanya bisa diakses dari package tempat dia dibuat. Jika kita membuat function, method, struct, interface, dan properti dengan huruf awal kapital, maka artinya dia bisa diakses dari package lain (Public). Jika kita membuat function, method, struct, interface, dan properti dengan huruf awal kecil, maka artinya dia hanya bisa diakses dari package tempat dia dibuat (Private). Di Go-Lang untuk menentukan Access Modifier cukup dengan nama function atau variable.
	// Contoh Programnya :
	// package helper
	// var version = "1.0.0" // Package version dan ini termasuk private atau tidak bisa diakses dari luar
	// var Application = "Belajar Go-Lang" // Package Application dan ini termasuk public atau bisa diakses dari luar
	// TIDAK BISA DIAKSES DARI LUAR PACKAGE
	// func sayGoodBye(name string) string {
	// 	return "Good Bye" + name
	// }
	// BISA DIAKSES DARI LUAR PACKAGE
	// func SayHello(name string) string {
	// 	return "Hello" + name
	// }
	// func main() {
	// 	fmt.Println(helper.Application)
	// 	fmt.Println(helper.sayGoodBye("Sendy"))
	// 	fmt.Println(helper.SayHello("Sendy"))
	// }

	// Bab Package Initialization //
	// Saat kita membuat package, kita bisa membuat function init di dalamnya. Function init ini akan dieksekusi ketika package diakses. Biasanya function ini digunakan untuk inisialisasi package. Jika kita membuat package, kita bisa membuat function init di dalamnya. Function init ini akan dieksekusi ketika package diakses. Biasanya function ini digunakan untuk inisialisasi package. Jika kita membuat package, kita bisa membuat function init di dalamnya. Function init ini akan dieksekusi ketika package diakses. Biasanya function ini digunakan untuk inisialisasi package.
	// Contoh Programnya bagian 1:
	// package database
	// var connection string
	// func init() {
	// 	connection = "MySQL"
	// }
	// func GetDatabase() string {
	// 	return connection
	// }
	// func main() {
	// 	fmt.Println(database.GetDatabase())
	// }
	// Bagian 2:
	// package main
	// import (
	// 	"fmt"
	// 	"belajar-golang-dasar/database"
	// )
	// func main() {
	// 	fmt.Println(database.GetDatabase())
	// }

	// Bab Blank Identifier //
	// Blank Identifier adalah kemampuan untuk menampung data yang tidak digunakan ke dalam variabel. Biasanya hal ini digunakan ketika kita import package, namun tidak menggunakannya. Jika kita tidak menggunakan package yang sudah diimport, maka program akan error. Untuk menghindari error tersebut, kita bisa menggunakan blank identifier. Kadang kita hanya ingin menjalankan init function di package tanpa harus mengeksekusi salah satu function nya yang ada di package. Secara default Go-Lang akan komplen ketika ada package yang di import namun tidak digunakan. Untuk menangani hal tersebut kita bisa menggunakan black identifier (...) sebelum packange ketika melakukan import

	// Bab Error //
	// Error Interface
	// Go-Lang memiliki interface yang digunakan sebagai kontrak untuk membuat error, dan nama interface nya adalah 'error'
	// Contoh Implementasi:
	//type error interface {
	// Error() string
	// }
	// Contoh Program 1:
	// import (
	//		"errors"
	//		"fmt"
	// )
	// func Pembagian(nilai int, pembagi int) (int, error) {
	//		if pembagi == 0 {
	//			return 0, errors.New("PEMBAGIAN DGN NOL")
	//		} else {
	//			return nilai / pembagi, nil
	//		}
	// func main() {
	//		hasil, err := Pembagian(100, 10)
	//		if err = nil {
	// 			fmt.println("Hasil =", hasil)
	//		} else {
	//			fmt.Println("Error", err.Error())
	//		}
	// }

	// Bab Custom Error //
	// Karena Error adalah sebuah interface, jadi jika kita ingin membuar error sendiri, kita bisa membuat struct yang mengikuti kontrak dari interface error.
	// Contoh Implementasi 1:
	// type valiidationError struct {
	// 	Message string
	// }
	// func (v *validationError) Error() string {
	// 	return v.Message
	// }
	// Contoh Implementasi 2:
	// type notFoundError struct {
	// 	Message string
	// }
	// func (v *notFoundError) Error() string {
	// 	Message String
	// }
	// func (v *notFoundError) Error() string {
	// 	return v.Message
	// }
	// Program Menggunakan Custom Error:
	// func SaveData(id string, data any) error {
	//	if id == "" {
	//		return &validationError{Message: "validation error"}
	// 	}
	//	if id != "Sendy" {
	//		return &notFoundError{Message: "data not found"}
	//	}
	//	return nil
	// }
	// Program cek Jenis Error dari Custom Error yang dibuat:
	// func main() {
	// err := SaveData("", nil)
	//	if err != nil {
	//		if validationError, ok := err.(*validationError); ok {
	//			fmt.Println("validation error", validationErr.Message)
	//		} else if notFoundErr, ok := err.(*notFoundError); ok {
	//		} else {
	//			fmt.Println("unknown error:", err.Error())
	//		}
	//	}
	// }

	// Untuk selanjutnya ada di FILE beda!!
}
