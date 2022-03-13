package moduleProject

import "fmt"
import "strings"


func (struk Strukpembayaran) GrandOpening(toko string) {
	fmt.Println("Thank You for shopping at", toko," Shop. Hope you are satisfied with our service Mr/Mrs", struk.Name)
}
// function parameter

func Welcome (nama string) {
	fmt.Println("Hai Mr/Mrs")
	fmt.Println("Please enter your name")
	fmt.Scanln(&nama)
	fmt.Printf("Please see our products Mr/Mrs %s\n",nama)

}
//function variadic
func ItemPakaian (nama string, pilihanBaju... string){
	var bajuAsstring = strings.Join (pilihanBaju,",")
	fmt.Printf("Hello Mr/Mrs %s\n", nama)
	fmt.Printf("This is a selection of clothes in our shop %s\n", bajuAsstring)
}
//function multiple return
func ProfitSale (harga, hargaPokok int) (int, int) {
	keuntungan := harga - hargaPokok
	diskon := (10/100) * hargaPokok
	return keuntungan, diskon
}
// anonymousstruct
func AlamatToko () {
	Addrres := struct {
		NamaToko string
		AddrresToko string
	}{
		NamaToko: "Berkah",
		AddrresToko: "Jl.Soekarno Hatta No.1" ,
	}
		fmt.Println("Welcome to ", Addrres.NamaToko,"Shop which is located at", Addrres.AddrresToko)	
}

// defer
func Selesai (){
	defer fmt.Println("terimakasih sudah berbelanja")
	fmt.Println("==================================")
	fmt.Println("==================================")
	fmt.Println("==================================")	
}

//Struct
func CustomerData(){
	Pelanggan1 := DataPelanggan {
		Nama:         "Muhammad Iqbal",
		Alamat : "Padang",
		Umur:         25,
		Pekerjaan:        "Staf Operasional dan Pembayaran",
	}
	fmt.Println("Data Pelanggan>>>>>>>>>>>>>>>>>>>>")
	fmt.Println("Nama		:",Pelanggan1.Nama)
	fmt.Println("Alamat		:",Pelanggan1.Alamat)
	fmt.Println("Umur		:",Pelanggan1.Umur)
	fmt.Println("Pekerjaan	:",Pelanggan1.Pekerjaan)

}