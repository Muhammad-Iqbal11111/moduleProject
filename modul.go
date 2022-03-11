package moduleProject

import "fmt"

// function parameter

func Welcome(nama, jawaban string) {
	fmt.Println("Selamat Datang di Lapak Kami, Apakah anda ingin berbelanja(Y/N)")
        if jawaban== Y{
	fmt.Println("Silahkan masukkan nama Anda")
	fmt.Scanln(&nama)
	fmt.Printf("Hai kak %q, silahkan untuk melihat lihat produk kami\n",nama)
		} else {
	fmt.Println("Mohon maaf kami tidak melayani peminta minta")
		}
}
//function variadic
func ItemPakaian(nama string, pilihanBaju... string){
	var bajuAsstring = string.Join(pilihanBaju,",")
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
		Diskon int
		Pay int
	}{
		Diskon: 10/100 * harga,
		Pay: harga,
	}
	if Pembayaran.Pay > 100000{

		fmt.Println("Selamat Anda Mendapat Potongan Harga", Pembayaran.Diskon, "Jumlah yang dibayarkan", (Pembayaran.Pay-Pembayaran.Diskon))
	} else {
		fmt.Println("Maaf anda tidak mendapat ptongan harga, Jumlah yang dibayarkan", Pembayaran.Pay)
	}
}