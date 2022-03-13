package moduleProject

import "fmt"

// function parameter

func Welcome(nama, jawaban string) {
	fmt.Println("Selamat Datang di Lapak Kami, Apakah anda ingin berbelanja(Y/N)")
	fmt.Println(&jawaban)
        if jawaban== "Y"{
	fmt.Println("Silahkan masukkan nama Anda")
	fmt.Scanln(&nama)
	fmt.Printf("Hai kak %q, silahkan untuk melihat lihat produk kami\n",nama)
		} else {
	fmt.Println("Mohon maaf kami tidak melayani peminta minta")
		}
}
//function variadic
func ItemPakaian(nama string, pilihanBaju... string){
	var bajuAsstring = strings.Join(pilihanBaju,",")
	fmt.Printf("Hello Mr/Mrs %s\n", nama)
	fmt.Printf("This is a selection of clothes in our shop %s\n", bajuAsstring)
}
//function multiple return
func ProfitSale(harga, hargaPokok float64) (float64, float64) {
	keuntungan := harga - hargaPokok
	persentaseKeuntungan := (keuntungan / hargaPokok) * 100
	return keuntungan, persentaseKeuntungan
}
// anonymousstruct
func Pembayaran() {
	Bayar := struct {
		Diskon float64
		Pay float64
	}{
		Diskon: 20000,
		Pay: 200000,
	}
	if Bayar.Pay > 100000{

		fmt.Println("Selamat Anda Mendapat Potongan Harga", Bayar.Diskon, "Jumlah yang dibayarkan", (Bayar.Pay-Bayar.Diskon))
	} else {
		fmt.Println("Maaf anda tidak mendapat ptongan harga, Jumlah yang dibayarkan", Bayar.Pay)
	}
}



