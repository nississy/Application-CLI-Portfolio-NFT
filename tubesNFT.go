package main

import "fmt"

const NMAX int = 1000

type nft struct {
	judulAset, namaCreator, blockChain, status, rarity string
	tanggalBeli                                        int
	nilai                                              float64
}

type koleksi [NMAX - 1]nft

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
			fmt.Println("   Terima kasih telah menggunakan aplikasi ini. Sampai jumpa kembali!")
		default:
			fmt.Println("   Opsi tidak valid. Silahkan input ulang lagi.")
		}
	}
}

func menuAwal() {
	fmt.Println("====================================")
	fmt.Println("   [ Aplikasi CLI Portfolio NFT ]   ")
	fmt.Println("====================================")
	fmt.Println("   1. Tambah Aset NFT")
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

	var hari, bulan, tahun int
	fmt.Print("Masukkan tanggal beli (dd mm yyyy): ")
	fmt.Scan(&hari, &bulan, &tahun)
	koleksiNFT[jumlahData].tanggalBeli = tahun*10000 + bulan*100 + hari

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
	var pilihan int

	if jumlahData == 0 {
		fmt.Println("Maaf anda belum memiliki aset dalam koleksi.")
		fmt.Println("Silahkan tambahkan aset terlebih dahulu!")
	} else {
		fmt.Println("   Pilih cara menampilkan aset")
		fmt.Println("------------------------------------")
		fmt.Println("   1. Tampilkan berdasarkan nilai (termahal ke termurah)")
		fmt.Println("   2. Tampilkan berdasarkan tanggal beli (terlama ke terbaru)")
		fmt.Println("   3. Tampilkan berdasarkan rarity")
		fmt.Println("   4. Cari berdasarkan judul aset (binary search)")
		fmt.Println("------------------------------------")
		fmt.Print("Pilihan anda: ")
		fmt.Scan(&pilihan)
		fmt.Println("------------------------------------")

		switch pilihan {
		case 1:
			SelectionSortNilaiDescending(&koleksiNFT, jumlahData)
			tampilanAset(koleksiNFT)
		case 2:
			InsertionSortTanggalAscending(&koleksiNFT, jumlahData)
			fmt.Println("Aset diurutkan berdasarkan tanggal beli awal")
			tampilanAset(koleksiNFT)
		case 3:
			var rarity string
			fmt.Print("Masukkan rarity yang ingin dicari (Common/Rare/Epic/Legendary): ")
			fmt.Scan(&rarity)
			cariAsetBerdasarkanRarity(koleksiNFT, rarity)
		case 4:
			cariAsetDenganJudulBinary(koleksiNFT, jumlahData)
		default:
			fmt.Println("Opsi tidak valid. Silahkan input ulang lagi.")
		}
	}
}

func tampilanAset(koleksiNFT koleksi) {
	var i, t int

	for i = 0; i < jumlahData; i++ {
		fmt.Println("====================================")
		fmt.Printf("Judul Aset  : %s\n", koleksiNFT[i].judulAset)
		fmt.Printf("Creator     : %s\n", koleksiNFT[i].namaCreator)
		fmt.Printf("Blockchain  : %s\n", koleksiNFT[i].blockChain)
		fmt.Printf("Status      : %s\n", koleksiNFT[i].status)

		t = koleksiNFT[i].tanggalBeli
		fmt.Printf("Tanggal Beli: %02d-%02d-%04d\n", t%100, (t/100)%100, t/10000)

		fmt.Printf("Nilai       : Rp. %.2f\n", koleksiNFT[i].nilai)
		fmt.Printf("Rarity      : %s\n", koleksiNFT[i].rarity)
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

//Menerapkan Linear Search
func hapusAset(koleksiNFT *koleksi) {
	var judulCari string
	var index, i int
	var ditemukan bool

	ditemukan = false

	if jumlahData == 0 {
		fmt.Println("Maaf anda belum memiliki aset dalam koleksi untuk dihapus.")
		fmt.Println("Silahkan tambahkan aset terlebih dahulu!")
	} else {
		fmt.Print("Masukkan judul aset yang ingin dihapus: ")
		fmt.Scan(&judulCari)

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

//Mengurutkan(Descending) berdasarkan nilai aset
func SelectionSortNilaiDescending(koleksiNFT *koleksi, N int) {
	var i, idx, pass int
	var temp nft

	if N == 0 {
		fmt.Println("Belum ada data untuk diurutkan. Tambahkan aset terlebih dahulu.")
	} else {
		for pass = 0; pass < N-1; pass++ {
			idx = pass
			for i = pass + 1; i < N; i++ {
				if koleksiNFT[i].nilai > koleksiNFT[idx].nilai {
					idx = i
				}
			}
			if idx != pass {
				temp = koleksiNFT[pass]
				koleksiNFT[pass] = koleksiNFT[idx]
				koleksiNFT[idx] = temp
			}
		}
		fmt.Println("Urutan aset anda dari yang paling mahal")
	}
}

//Mengurutkan(Ascending) berdasarkan tanggal aset
func InsertionSortTanggalAscending(koleksiNFT *koleksi, N int) {
	var i, pass int
	var temp nft

	pass = 1
	for pass < N {
		i = pass
		temp = koleksiNFT[pass]
		for i > 0 && temp.tanggalBeli < koleksiNFT[i-1].tanggalBeli {
			koleksiNFT[i] = koleksiNFT[i-1]
			i = i - 1
		}
		koleksiNFT[i] = temp
		pass = pass + 1
	}
}

//Penerapan binary search untuk mencari data berdasarkan judul aset
//Nilai diurutkan terlebih dahulu pakai insertion
func InsertionSortJudulAscending(koleksiNFT *koleksi, N int) {
	var i, pass int
	var temp nft

	pass = 1
	for pass < N {
		i = pass
		temp = koleksiNFT[pass]
		for i > 0 && temp.judulAset < koleksiNFT[i-1].judulAset {
			koleksiNFT[i] = koleksiNFT[i-1]
			i = i - 1
		}
		koleksiNFT[i] = temp
		pass = pass + 1
	}
}

func BinarySearchJudul(koleksiNFT koleksi, N int, target string) int {
	var low, high, mid int
	low = 0
	high = N - 1
	for low <= high {
		mid = (low + high) / 2
		if koleksiNFT[mid].judulAset == target {
			return mid
		} else if koleksiNFT[mid].judulAset < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func cariAsetDenganJudulBinary(koleksiNFT koleksi, N int) {
	var pos int
	if jumlahData == 0 {
		fmt.Println("Belum ada aset dalam koleksi.")
		return
	}

	var target string
	fmt.Print("Masukkan judul aset yang ingin dicari: ")
	fmt.Scan(&target)

	// Urutkan terlebih dahulu berdasarkan judul
	InsertionSortJudulAscending(&koleksiNFT, N)

	// Cari menggunakan binary search
	pos = BinarySearchJudul(koleksiNFT, N, target)
	if pos == -1 {
		fmt.Println("Aset dengan judul tersebut tidak ditemukan.")
	} else {
		fmt.Println("Aset ditemukan!")
		fmt.Println("====================================")
		fmt.Printf("Judul Aset  : %s\n", koleksiNFT[pos].judulAset)
		fmt.Printf("Creator     : %s\n", koleksiNFT[pos].namaCreator)
		fmt.Printf("Blockchain  : %s\n", koleksiNFT[pos].blockChain)
		fmt.Printf("Status      : %s\n", koleksiNFT[pos].status)

		t := koleksiNFT[pos].tanggalBeli
		fmt.Printf("Tanggal Beli: %02d-%02d-%04d\n", t%100, (t/100)%100, t/10000)

		fmt.Printf("Nilai       : Rp. %.2f\n", koleksiNFT[pos].nilai)
		fmt.Printf("Rarity      : %s\n", koleksiNFT[pos].rarity)
		fmt.Println("====================================")
	}
}

//Penerapan linier search untuk list seluruh aset pada kasus rarity
func cariAsetBerdasarkanRarity(koleksiNFT koleksi, rarity string) {
	var ditemukan bool
	var i int
	for i = 0; i < jumlahData; i++ {
		if koleksiNFT[i].rarity == rarity {
			ditemukan = true
			t := koleksiNFT[i].tanggalBeli
			fmt.Println("====================================")
			fmt.Printf("Judul Aset  : %s\n", koleksiNFT[i].judulAset)
			fmt.Printf("Creator     : %s\n", koleksiNFT[i].namaCreator)
			fmt.Printf("Blockchain  : %s\n", koleksiNFT[i].blockChain)
			fmt.Printf("Status      : %s\n", koleksiNFT[i].status)
			fmt.Printf("Tanggal Beli: %02d-%02d-%04d\n", t%100, (t/100)%100, t/10000)
			fmt.Printf("Nilai       : Rp. %.2f\n", koleksiNFT[i].nilai)
			fmt.Printf("Rarity      : %s\n", koleksiNFT[i].rarity)
		}
	}
	if !ditemukan {
		fmt.Printf("Tidak ditemukan aset dengan rarity: %s\n", rarity)
	}
}
