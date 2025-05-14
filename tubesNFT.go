package main

import "fmt"

type nft struct {
	judulAset, namaCreator, blockChain, status, rarity string
	tanggalBeli                                        int
	nilai                                              float64
}

type koleksi [1000]nft

var jumlahData int

func main() {
	var koleksiNFT koleksi
	var opsi int
	for opsi != 5 {
		menuAwal()
		opsi = pilihOpsi(opsi)

		switch opsi {
		case 1:
			tambahAset(&koleksiNFT)
		case 2:
			lihatAsetList(koleksiNFT)
		case 5:
			fmt.Println("Terima kasih telah menggunakan aplikasi. Selamat jumpa kembali!")
		default:
			fmt.Println("Opsi tidak valid.")
		}
	}
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

func tambahAset(koleksiNFT *koleksi) {
	if jumlahData >= len(*koleksiNFT) {
		fmt.Println("Koleksi sudah penuh!")
	}

	fmt.Println("Masukkan judul aset")
	fmt.Print(">> ")
	fmt.Scan(&koleksiNFT[jumlahData].judulAset)
	fmt.Println("Masukkan nama creator")
	fmt.Print(">> ")
	fmt.Scan(&koleksiNFT[jumlahData].namaCreator)
	fmt.Println("Masukkan blockchain")
	fmt.Print(">> ")
	fmt.Scan(&koleksiNFT[jumlahData].blockChain)
	fmt.Println("Masukkan status (terjual/belum)")
	fmt.Print(">> ")
	fmt.Scan(&koleksiNFT[jumlahData].status)
	fmt.Println("Masukkan tanggal beli")
	fmt.Print(">> ")
	fmt.Scan(&koleksiNFT[jumlahData].tanggalBeli)
	fmt.Println("Masukkan nilai aset (IDR)")
	fmt.Print(">> ")
	fmt.Scan(&koleksiNFT[jumlahData].nilai)

	koleksiNFT[jumlahData].rarity = tentukanRarity(koleksiNFT[jumlahData].nilai)

	jumlahData++

	fmt.Println("Aset berhasil ditambahkan!")
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

func lihatAsetList(koleksiNFT koleksi) {
	var i int
	if jumlahData == 0 {
		fmt.Println("Belum ada aset dalam koleksi.")
	} else {
		fmt.Println("List Aset yang anda miliki saat ini:")
		for i = 0; i < jumlahData; i++ {
			fmt.Println("====================================")
			fmt.Printf("Judul Aset: %s\n", koleksiNFT[i].judulAset)
			fmt.Printf("Creator     : %s\n", koleksiNFT[i].namaCreator)
			fmt.Printf("Blockchain  : %s\n", koleksiNFT[i].blockChain)
			fmt.Printf("Status      : %s\n", koleksiNFT[i].status)
			fmt.Printf("Tanggal Beli: %d\n", koleksiNFT[i].tanggalBeli)
			fmt.Printf("Nilai       : Rp %.2f\n", koleksiNFT[i].nilai)
			fmt.Printf("Rarity      : %s\n", koleksiNFT[i].rarity)
		}
		fmt.Println("====================================")
	}
}
