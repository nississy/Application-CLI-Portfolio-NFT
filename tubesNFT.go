package main

import "fmt"

type nft struct {
	judulAset, namaCreator, blockChain, status, rarity string
	tanggalBeli                                        int
	nilai                                              float64
}

var koleksi [1000]nft
var jumlahData int

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
	fmt.Println("------------------------------------")
	return opsi
}

func tambahAset() {
	var judul, creator, chain, status string
	var tanggal int
	var nilai float64

	fmt.Print("Masukkan judul aset: ")
	fmt.Scan(&judul)
	fmt.Print("Masukkan nama creator: ")
	fmt.Scan(&creator)
	fmt.Print("Masukkan blockchain: ")
	fmt.Scan(&chain)
	fmt.Print("Masukkan status (terjual/belum): ")
	fmt.Scan(&status)
	fmt.Print("Masukkan tanggal beli (misal 20250101): ")
	fmt.Scan(&tanggal)
	fmt.Print("Masukkan nilai aset (IDR): ")
	fmt.Scan(&nilai)

	koleksi[jumlahData].judulAset = judul
	koleksi[jumlahData].namaCreator = creator
	koleksi[jumlahData].blockChain = chain
	koleksi[jumlahData].status = status
	koleksi[jumlahData].tanggalBeli = tanggal
	koleksi[jumlahData].nilai = nilai
	koleksi[jumlahData].rarity = tentukanRarity(nilai)

	jumlahData++

	fmt.Println("Aset berhasil ditambahkan!\n")
}

func tentukanRarity(nilai float64) string {
	if nilai > 15000000 {
		return "Legendary"
	} else if nilai >= 1000000 {
		return "Epic"
	} else if nilai >= 100000 {
		return "Rare"
	} else {
		return "Common"
	}
}
