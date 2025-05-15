package main

import "fmt"

type nft struct {
	judulAset, namaCreator, blockChain, status, rarity, tanggalBeli string
	nilai                                                           float64
}

type koleksi [1000]nft

var jumlahData int

func main() {
	var koleksiNFT koleksi
	var totalNilaiAset float64
	var opsi int
	for opsi != 5 {
		menuAwal()
		opsi = pilihOpsi(opsi)

		switch opsi {
		case 1:
			tambahAset(&koleksiNFT)
		case 2:
			lihatAsetList(koleksiNFT)
		case 3:
			lihatTotalAset(koleksiNFT, &totalNilaiAset)
		case 4:
			hapusAset(&koleksiNFT)
		case 5:
			fmt.Println("   Terima kasih telah menggunakan aplikasi. Sampai jumpa kembali!")
		default:
			fmt.Println("   Opsi tidak valid. Silahkan input ulang lagi.")
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

	fmt.Print("Masukkan judul aset: ")
	fmt.Scan(&koleksiNFT[jumlahData].judulAset)
	fmt.Print("Masukkan nama creator: ")
	fmt.Scan(&koleksiNFT[jumlahData].namaCreator)
	fmt.Print("Masukkan blockchain: ")
	fmt.Scan(&koleksiNFT[jumlahData].blockChain)
	fmt.Print("Masukkan status (terjual/belum): ")
	fmt.Scan(&koleksiNFT[jumlahData].status)
	fmt.Print("Masukkan tanggal beli (dd-mm-yyyy): ")
	fmt.Scan(&koleksiNFT[jumlahData].tanggalBeli)
	fmt.Print("Masukkan nilai aset (IDR): Rp.")
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
		fmt.Println("Maaf anda belum memiliki aset dalam koleksi.")
		fmt.Println("Silahkan tambahkan aset terlebih dahulu!")
	} else {
		fmt.Println("====================================")
		fmt.Println("List Aset yang anda miliki saat ini:")
		for i = 0; i < jumlahData; i++ {
			fmt.Println("====================================")
			fmt.Printf("Judul Aset  : %s\n", koleksiNFT[i].judulAset)
			fmt.Printf("Creator     : %s\n", koleksiNFT[i].namaCreator)
			fmt.Printf("Blockchain  : %s\n", koleksiNFT[i].blockChain)
			fmt.Printf("Status      : %s\n", koleksiNFT[i].status)
			fmt.Printf("Tanggal Beli: %s\n", koleksiNFT[i].tanggalBeli)
			fmt.Printf("Nilai       : Rp %.2f\n", koleksiNFT[i].nilai)
			fmt.Printf("Rarity      : %s\n", koleksiNFT[i].rarity)
		}
	}
}

func lihatTotalAset(koleksiNFT koleksi, totalNilaiAset *float64) {
	var i int
	if jumlahData == 0 {
		fmt.Println("Maaf anda belum memiliki aset dalam koleksi.")
		fmt.Println("Silahkan tambahkan aset terlebih dahulu!")
	} else {
		for i = 0; i < jumlahData; i++ {
			*totalNilaiAset = *totalNilaiAset + koleksiNFT[i].nilai
		}
		fmt.Printf("Total aset dalam koleksi: %d\n", jumlahData)
		fmt.Printf("Value: %f\n", *totalNilaiAset)
	}
}

func hapusAset(koleksiNFT *koleksi) {
	var judulCari string
	var index, i int
	var ditemukan bool 
	
	ditemukan = false

	fmt.Print("Masukkan judul aset yang ingin dihapus: ")
	fmt.Scan(&judulCari)

	if jumlahData == 0 {
		fmt.Println("Maaf anda belum memiliki aset dalam koleksi.")
		fmt.Println("Silahkan tambahkan aset terlebih dahulu!")
	} else {
		for i = 0; i < jumlahData; i++ {
			if koleksiNFT[i].judulAset == judulCari {
				index = i
				ditemukan = true
			}
		}

		if ditemukan {
			for i = index; i < jumlahData-1; i++ {
				koleksiNFT[i] = koleksiNFT[i+1]
			}
			jumlahData--
			fmt.Println("Aset berhasil dihapus!")
		} else {
			fmt.Println("Aset tidak ditemukan!")
		}
	}
}
