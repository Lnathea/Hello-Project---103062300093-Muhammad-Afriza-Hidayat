// Muhammad Afriza Hidayat
// 103062300093

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// Definisi tipe bentukan untuk Buku dan Peminjaman
type Buku struct {
	ID        int
	Judul     string
	Pengarang string
	Tahun     int
}

type Peminjaman struct {
	IDBuku          int
	TanggalPinjam   time.Time
	TanggalKembali  time.Time
	DendaPerHari    float64
	HariTerlambat   int
	TarifPeminjaman float64
}

// Array statis untuk menyimpan data buku dan peminjaman
const maxSize = 100

var buku [maxSize]Buku
var lastIndex int = -1
var peminjaman [maxSize]Peminjaman
var lastPeminjamanIndex int = -1

// Fungsi untuk menambahkan buku baru
func tambahBuku(id int, judul string, pengarang string, tahun int) {
	if lastIndex < maxSize-1 {
		lastIndex++
		buku[lastIndex] = Buku{id, judul, pengarang, tahun}
	} else {
		fmt.Println("Array buku sudah penuh.")
	}
}

// Fungsi untuk mengubah data buku
func ubahBuku(id int, judul string, pengarang string, tahun int) {
	for i := 0; i <= lastIndex; i++ {
		if buku[i].ID == id {
			buku[i].Judul = judul
			buku[i].Pengarang = pengarang
			buku[i].Tahun = tahun
			return
		}
	}
	fmt.Println("Buku dengan ID tersebut tidak ditemukan.")
}

// Fungsi untuk menghapus buku
func hapusBuku(id int) {
	for i := 0; i <= lastIndex; i++ {
		if buku[i].ID == id {
			buku[i] = buku[lastIndex]
			lastIndex--
			fmt.Println("Buku telah Dihapus.")
			fmt.Println("--------------------")
			return

		}
	}
	fmt.Println("Buku dengan ID tersebut tidak ditemukan.")
}

// Fungsi untuk menambahkan peminjaman baru
func tambahPeminjaman(idBuku int, tanggalPinjam, tanggalKembali time.Time, dendaPerHari float64) {
	if lastPeminjamanIndex < maxSize-1 {
		lastPeminjamanIndex++
		peminjaman[lastPeminjamanIndex] = Peminjaman{
			IDBuku:          idBuku,
			TanggalPinjam:   tanggalPinjam,
			TanggalKembali:  tanggalKembali,
			DendaPerHari:    dendaPerHari,
			HariTerlambat:   0, // Ini akan dihitung saat pengembalian buku
			TarifPeminjaman: 0, // Ini akan dihitung berdasarkan durasi peminjaman
		}
	} else {
		fmt.Println("Array peminjaman sudah penuh.")
	}
}

func cariBuku(id int) (*Buku, error) {
	for i := 0; i <= lastIndex; i++ {
		if buku[i].ID == id {
			return &buku[i], nil
		}
	}
	return nil, errors.New("Buku tidak ditemukan")
}

// Fungsi untuk mengubah data peminjaman
func ubahPeminjaman(idBuku int, tanggalPinjam, tanggalKembali time.Time, dendaPerHari float64) {
	for i := 0; i <= lastPeminjamanIndex; i++ {
		if peminjaman[i].IDBuku == idBuku {
			peminjaman[i].TanggalPinjam = tanggalPinjam
			peminjaman[i].TanggalKembali = tanggalKembali
			peminjaman[i].DendaPerHari = dendaPerHari
			return
		}
	}
	fmt.Println("Peminjaman dengan ID Buku tersebut tidak ditemukan.")
}

// Fungsi untuk menghapus peminjaman
func hapusPeminjaman(idBuku int) {
	for i := 0; i <= lastPeminjamanIndex; i++ {
		if peminjaman[i].IDBuku == idBuku {
			peminjaman[i] = peminjaman[lastPeminjamanIndex]
			lastPeminjamanIndex--
			return
		}
	}
	fmt.Println("Peminjaman dengan ID Buku tersebut tidak ditemukan.")
}

// Fungsi untuk mencari buku berdasarkan judul menggunakan Sequential Search
func cariBukuSequential(judul string) *Buku {
	for i := 0; i <= lastIndex; i++ {
		if buku[i].Judul == judul {
			return &buku[i]
		}
	}
	return nil // Buku tidak ditemukan
}

// Fungsi untuk mencari buku berdasarkan ID menggunakan Binary Search
// Catatan: Array buku harus sudah terurut berdasarkan ID sebelum menggunakan binary search
func cariBukuBinary(id int) *Buku {
	low := 0
	high := lastIndex

	for low <= high {
		mid := low + (high-low)/2
		if buku[mid].ID == id {
			return &buku[mid]
		} else if buku[mid].ID < id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return nil // Buku tidak ditemukan
}

// Fungsi untuk mengurutkan buku berdasarkan judul menggunakan Selection Sort
func urutkanBukuSelection() {
	for i := 0; i <= lastIndex; i++ {
		minIndex := i
		for j := i + 1; j <= lastIndex; j++ {
			if buku[j].Judul < buku[minIndex].Judul {
				minIndex = j
			}
		}
		// Tukar posisi buku[i] dengan buku[minIndex]
		buku[i], buku[minIndex] = buku[minIndex], buku[i]
	}
}

// Fungsi untuk mengurutkan buku berdasarkan tahun menggunakan Insertion Sort
func urutkanBukuInsertion() {
	for i := 1; i <= lastIndex; i++ {
		key := buku[i]
		j := i - 1
		for j >= 0 && buku[j].Tahun > key.Tahun {
			buku[j+1] = buku[j]
			j--
		}
		buku[j+1] = key
	}
}

func main() {
menuUtama:
	for {
		fmt.Print("\033[H\033[2J")
		fmt.Println("============== Menu Utama ==============")
		fmt.Println("1. Tambahkan, ubah, dan hapus data buku")
		fmt.Println("2. Tambahkan, ubah, dan hapus data peminjaman buku")
		fmt.Println("3. Lihat data buku terurut berdasarkan kategori")
		fmt.Println("4. Cari buku dengan kata kunci tertentu")
		fmt.Println("5. Perhitungan tarif peminjaman dan denda keterlambatan")
		fmt.Println("6. Tampilkan buku yang sedang dipinjam dan 5 buku terfavorit")
		fmt.Println("7. Keluar")

		var pilihan int
		fmt.Print("Pilih menu (1-7): ")
		fmt.Scan(&pilihan)
		fmt.Print("\033[H\033[2J")

		switch pilihan {
		case 1:
			fmt.Println("============== Menu Buku ==============")
			for {
				var subPilihan int
				fmt.Println("1. Tambah Buku")
				fmt.Println("2. Ubah Buku")
				fmt.Println("3. Hapus Buku")
				fmt.Println("4. Kembali ke Menu Utama")
				fmt.Print("Pilih aksi (1-4): ")
				fmt.Scan(&subPilihan)
				fmt.Print("\033[H\033[2J")

				switch subPilihan {
				case 1:
					for {
						var id, tahun int
						var judul, pengarang, lanjut string
						fmt.Print("Masukkan ID Buku: ")
						fmt.Scan(&id)
						fmt.Print("Masukkan Judul Buku: ")
						fmt.Scan(&judul)
						fmt.Print("Masukkan Pengarang Buku: ")
						fmt.Scan(&pengarang)
						fmt.Print("Masukkan Tahun Terbit: ")
						fmt.Scan(&tahun)
						tambahBuku(id, judul, pengarang, tahun)
						fmt.Print("Lanjutkan menambah buku? (y/n): ")
						fmt.Scan(&lanjut)
						fmt.Print("\033[H\033[2J")
						if lanjut != "y" {
							break
						}
					}

				case 2:
					for {
						var id, tahun int
						var judul, pengarang, lanjut string
						fmt.Print("Masukkan ID Buku yang akan diubah: ")
						fmt.Scan(&id)

						// Ambil data buku sebelum diubah
						bukuSebelum, err := cariBuku(id)
						if err != nil {
							fmt.Println("Buku tidak ditemukan")
							continue
						}

						originalJudul := bukuSebelum.Judul
						originalPengarang := bukuSebelum.Pengarang
						originalTahun := bukuSebelum.Tahun

						fmt.Print("Masukkan Judul Buku baru: ")
						fmt.Scan(&judul)
						fmt.Print("Masukkan Pengarang Buku baru: ")
						fmt.Scan(&pengarang)
						fmt.Print("Masukkan Tahun Terbit baru: ")
						fmt.Scan(&tahun)

						// Ubah data buku
						ubahBuku(id, judul, pengarang, tahun)

						// Tampilkan perubahan data buku
						fmt.Println(" ")
						fmt.Println("Data Buku Sebelum Diubah:")
						fmt.Printf("ID: %d, Judul: %s, Pengarang: %s, Tahun Terbit: %d\n", bukuSebelum.ID, originalJudul, originalPengarang, originalTahun)
						fmt.Println("--------------------")
						fmt.Println("Data Buku Setelah Diubah:")
						fmt.Printf("ID: %d, Judul: %s, Pengarang: %s, Tahun Terbit: %d\n", bukuSebelum.ID, judul, pengarang, tahun)

						fmt.Print("Lanjutkan? (y/n): ")
						fmt.Scan(&lanjut)
						fmt.Print("\033[H\033[2J")
						if lanjut != "y" {
							break
						}
					}

				case 3:
					var id int
					fmt.Print("Masukkan ID Buku yang akan dihapus: ")
					fmt.Scan(&id)
					hapusBuku(id)

				case 4:
					goto menuUtama
				default:
					fmt.Println("Pilihan tidak valid.")
				}

			}

		case 2:
			fmt.Println("========== Menu Peminjam ==========")
			for {
				var subPilihan int
				fmt.Println("1. Tambah Peminjaman")
				fmt.Println("2. Ubah Peminjaman")
				fmt.Println("3. Hapus Peminjaman")
				fmt.Println("4. Kembali ke Menu Utama")
				fmt.Print("Pilih aksi (1-4): ")
				fmt.Scan(&subPilihan)
				fmt.Print("\033[H\033[2J")

				switch subPilihan {
				case 1:
					for {
						var idBuku int
						var tanggalPinjamStr, tanggalKembaliStr, lanjut string
						var dendaPerHari float64
						var tanggalPinjam, tanggalKembali time.Time
						var err error

						reader := bufio.NewReader(os.Stdin)

						fmt.Print("Masukkan ID Buku: ")
						fmt.Scan(&idBuku)

						fmt.Print("Masukkan Tanggal Pinjam (YYYY-MM-DD): ")
						tanggalPinjamStr, _ = reader.ReadString('\n')
						tanggalPinjamStr = strings.TrimSpace(tanggalPinjamStr)
						tanggalPinjam, err = time.Parse("2006-01-02", tanggalPinjamStr)
						if err != nil {
							fmt.Println(" ")
							fmt.Println("Format tanggal pinjam tidak valid.")
							fmt.Println(" ")
							break // Keluar dari switch
						}
						fmt.Print("Masukkan Tanggal Kembali (YYYY-MM-DD): ")
						tanggalKembaliStr, _ = reader.ReadString('\n')
						tanggalKembaliStr = strings.TrimSpace(tanggalKembaliStr)
						tanggalKembali, err = time.Parse("2006-01-02", tanggalKembaliStr)
						if err != nil {
							fmt.Println("Format tanggal kembali tidak valid.")
							break // Keluar dari switch
						}

						fmt.Print("Masukkan Denda Per Hari: ")
						fmt.Scan(&dendaPerHari)

						tambahPeminjaman(idBuku, tanggalPinjam, tanggalKembali, dendaPerHari)
						fmt.Print("Lanjutkan menambah peminjam? (y/n): ")
						fmt.Scan(&lanjut)
						fmt.Print("\033[H\033[2J")
						if lanjut != "y" {
							break
						}
					}

				case 2:
					reader := bufio.NewReader(os.Stdin)
					var idBuku int
					var tanggalPinjamStr, tanggalKembaliStr string
					var dendaPerHari float64
					var tanggalPinjam, tanggalKembali time.Time
					var err error

					fmt.Print("Masukkan ID Peminjaman Buku yang akan diubah: ")
					fmt.Scan(&idBuku)

					fmt.Print("Masukkan Tanggal Pinjam baru (YYYY-MM-DD): ")
					tanggalPinjamStr, _ = reader.ReadString('\n')
					tanggalPinjamStr = strings.TrimSpace(tanggalPinjamStr)
					tanggalPinjam, err = time.Parse("2006-01-02", tanggalPinjamStr)
					if err != nil {
						fmt.Println("Format tanggal pinjam tidak valid.")
						break // Keluar dari switch
					}

					fmt.Print("Masukkan Tanggal Kembali baru (YYYY-MM-DD): ")
					tanggalKembaliStr, _ = reader.ReadString('\n')
					tanggalKembaliStr = strings.TrimSpace(tanggalKembaliStr)
					tanggalKembali, err = time.Parse("2006-01-02", tanggalKembaliStr)
					if err != nil {
						fmt.Println("Format tanggal kembali tidak valid.")
						break // Keluar dari switch
					}

					fmt.Print("Masukkan Denda Per Hari baru: ")
					fmt.Scan(&dendaPerHari)

					ubahPeminjaman(idBuku, tanggalPinjam, tanggalKembali, dendaPerHari)

				case 3:
					var idBuku int
					fmt.Print("Masukkan ID Buku yang akan dihapus: ")
					fmt.Scan(&idBuku)
					hapusPeminjaman(idBuku)
					fmt.Print("Peminjaman Buku telah Dihapus.")
					fmt.Print("\033[H\033[2J")

				case 4:
					goto menuUtama
				default:
					fmt.Println("Pilihan tidak valid.")
				}
			}

		case 3:
			fmt.Println("============== Tampilkan Data Buku ==============")
			for {
				var subPilihan int
				fmt.Println("1. Urutkan berdasarkan judul")
				fmt.Println("2. Urutkan berdasarkan tahun")
				fmt.Println("3. Kembali ke Menu Utama")
				fmt.Print("Pilih aksi (1-3): ")
				fmt.Scan(&subPilihan)
				fmt.Print("\033[H\033[2J")

				switch subPilihan {
				case 1:
					urutkanBukuSelection()
					fmt.Println("Data buku telah diurutkan berdasarkan judul.")
					for i := 0; i <= lastIndex; i++ {
						fmt.Printf("ID: %d, Judul: %s, Pengarang: %s, Tahun: %d\n", buku[i].ID, buku[i].Judul, buku[i].Pengarang, buku[i].Tahun)
						fmt.Println(" ")
					}
				case 2:
					urutkanBukuInsertion()
					fmt.Println("Data buku telah diurutkan berdasarkan tahun.")
					for i := 0; i <= lastIndex; i++ {
						fmt.Printf("ID: %d, Judul: %s, Pengarang: %s, Tahun: %d\n", buku[i].ID, buku[i].Judul, buku[i].Pengarang, buku[i].Tahun)
						fmt.Println(" ")
					}
				case 3:
					goto menuUtama
				default:
					fmt.Println("Pilihan tidak valid.")
				}
			}

		case 4:
			for {
				fmt.Println("============== Pencarian Buku ==============")
				var judul, lanjut string
				fmt.Print("Masukkan judul buku yang dicari: ")
				fmt.Scan(&judul)
				bukuFound := cariBukuSequential(judul)
				if bukuFound != nil {
					fmt.Printf("Buku dengan judul '%s' ditemukan!\n", judul)
					fmt.Printf("ID: %d, Judul: %s, Pengarang: %s, Tahun: %d\n", bukuFound.ID, bukuFound.Judul, bukuFound.Pengarang, bukuFound.Tahun)
					fmt.Print("Lanjutkan mencari buku? (y/n): ")
					fmt.Scan(&lanjut)
					fmt.Print("\033[H\033[2J")
					if lanjut != "y" {
						break
					}
				} else {
					fmt.Println("Buku tidak ditemukan.")
				}
			}

		case 5:
			for {
				fmt.Println("============== Menu Perhitungan Denda Buku ==============")
				var idBuku int
				var lanjut string
				fmt.Print("Masukkan ID Buku yang dipinjam: ")
				fmt.Scan(&idBuku)
				for i := 0; i <= lastPeminjamanIndex; i++ {
					if peminjaman[i].IDBuku == idBuku {
						tanggalKembali := peminjaman[i].TanggalKembali
						tanggalPinjam := peminjaman[i].TanggalPinjam
						dendaPerHari := peminjaman[i].DendaPerHari
						hariTerlambat := int(tanggalKembali.Sub(tanggalPinjam).Hours() / 24)
						if hariTerlambat > 7 {
							tarifPeminjaman := float64(hariTerlambat-7) * dendaPerHari
							fmt.Printf("Lama meminjam buku: %d hari\n", hariTerlambat)
							fmt.Printf("Denda keterlambatan: Rp %.2f\n", tarifPeminjaman)
							fmt.Print("Lanjutkan Perhitungan Denda Buku (y/n): ")
							fmt.Scan(&lanjut)
							fmt.Print("\033[H\033[2J")
							if lanjut != "y" {
								break
							} else {
								fmt.Println("Tidak ada denda keterlambatan.")
							}
						}
					}
				}
			}

		case 6:
			fmt.Println("Anda memilih menu menampilkan buku.")
			fmt.Println("1. Tampilkan Riwayat Peminjaman buku")
			fmt.Println("2. Tampilkan 5 buku terfavorit")
			var subPilihan int
			fmt.Print("Pilih tampilan (1-2): ")
			fmt.Scan(&subPilihan)

			switch subPilihan {
			case 1:
				fmt.Println("Buku yang sedang dipinjam:")
				for i := 0; i <= lastPeminjamanIndex; i++ {
					fmt.Printf("ID Buku: %d, Tanggal Pinjam: %s, Tanggal Kembali: %s\n", peminjaman[i].IDBuku, peminjaman[i].TanggalPinjam.Format("2006-01-02"), peminjaman[i].TanggalKembali.Format("2006-01-02"))
				}
			case 2:
				fmt.Println("5 Buku terfavorit:")
				// Implementasi untuk menampilkan 5 buku terfavorit
				// Contoh: menggunakan array untuk menyimpan jumlah peminjaman setiap buku
				var favorit [maxSize]int
				for i := 0; i <= lastIndex; i++ {
					favorit[i] = 0
				}
				for i := 0; i <= lastPeminjamanIndex; i++ {
					for j := 0; j <= lastIndex; j++ {
						if peminjaman[i].IDBuku == buku[j].ID {
							favorit[j]++
							break
						}
					}
				}
				// Urutkan array favorit menggunakan bubble sort
				for i := 0; i <= lastIndex; i++ {
					for j := i + 1; j <= lastIndex; j++ {
						if favorit[i] < favorit[j] {
							favorit[i], favorit[j] = favorit[j], favorit[i]
							buku[i], buku[j] = buku[j], buku[i]
						}
					}
				}
				// Tampilkan 5 buku terfavorit
				for i := 0; i < 5; i++ {
					fmt.Printf("ID: %d, Judul: %s, Pengarang: %s, Tahun: %d, Jumlah Peminjaman: %d\n", buku[i].ID, buku[i].Judul, buku[i].Pengarang, buku[i].Tahun, favorit[i])
				}
			default:
				fmt.Println("Pilihan tidak valid.")
			}

			var lanjut string
			fmt.Print("Lanjutkan? (y/n): ")
			fmt.Scan(&lanjut)
			if lanjut != "y" {
				break
			}

		case 7:
			fmt.Println("Anda telah keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")

		}

	}
}
