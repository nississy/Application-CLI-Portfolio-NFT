package main

import "fmt"

func main() {
	var opsi int
	menuAwal()
	opsi = pilihOpsi(opsi)
}

func menuAwal() {
	fmt.Println("====================================")
	fmt.Println("   [ Aplikasi CLI Portfolio NFT ]   ")
	fmt.Println("====================================")
	fmt.Println("   1. Tambah Aset CFT")
	fmt.Println("   2. Lihat Aset")
	fmt.Println("   3. Lihat Total Aset")
	fmt.Println("   4. Hapus aset")
	fmt.Println("   5. Keluar")
	fmt.Println("------------------------------------")
}

func pilihOpsi(opsi int) int {
	fmt.Print("   Silahkan pilih opsi: ")
	fmt.Scan(&opsi)
	return opsi
}
