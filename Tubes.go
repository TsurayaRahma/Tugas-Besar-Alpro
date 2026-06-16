package main

import "fmt"

const MAX int = 100
type Workout struct {
	id, durasi, kalori int
	nama, jenis,tanggal string
}

var listJenis = [5]string{
    "",
    "Cardio",
    "Strength",
    "Flexibility",
    "HIIT",
}

var data [MAX]Workout
var n int
var nextID int = 1

// ================= MENU 1 =================
func tambahWorkout() {
	if n >= MAX {
		fmt.Println("Data penuh")
		return
	}

	var w Workout

	w.id = nextID
	nextID++

	// Input nama
	fmt.Print("Nama latihan        : ")
	fmt.Scanln(&w.nama)

	// Pilih jenis olahraga (validasi)
	var pilihanJenis int
	pilihanJenis = 0

	for pilihanJenis < 1 || pilihanJenis > 4 {
		fmt.Println("\nJenis Olahraga:")
		fmt.Println("1. Cardio")
		fmt.Println("2. Strength")
		fmt.Println("3. Flexibility")
		fmt.Println("4. HIIT")
		fmt.Print("Masukkan angka pilihan (1-4): ")
		fmt.Scanln(&pilihanJenis)

		if pilihanJenis < 1 || pilihanJenis > 4 {
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
	w.jenis = listJenis[pilihanJenis]

	// Input durasi workout (harus > 0)
	w.durasi = 0
	for w.durasi <= 0 {
		fmt.Print("Durasi (menit) > 0 : ")
		fmt.Scanln(&w.durasi)

		if w.durasi <= 0 {
			fmt.Println("Durasi harus lebih dari 0.")
		}
	}

	// Input kalori (harus > 0)
	w.kalori = 0
	for w.kalori <= 0 {
		fmt.Print("Jumlah kalori > 0  : ")
		fmt.Scanln(&w.kalori)

		if w.kalori <= 0 {
			fmt.Println("Kalori harus lebih dari 0.")
		}
	}

	// Input tanggal
	fmt.Print("Tanggal             : ")
	fmt.Scanln(&w.tanggal)

	data[n] = w
	n++

	fmt.Println("Workout berhasil ditambahkan")
}

//==========================Menu 2=======================
func tampilkanWorkout() {
	var i int

	if n == 0 {
		fmt.Println("Belum ada data workout")
		return
	}

	fmt.Println("==============================================================")
	fmt.Println("ID | Nama | Jenis | Durasi | Kalori | Tanggal")
	fmt.Println("==============================================================")

	for i = 0; i < n; i++ {
		fmt.Printf("%d | %s | %s | %d | %d | %s\n",
			data[i].id,
			data[i].nama,
			data[i].jenis,
			data[i].durasi,
			data[i].kalori,
			data[i].tanggal)
	}

	fmt.Println("==============================================================")
}

// =================================Menu 3==============================
func ubahWorkout() {
	var id int
	var i int
	var ketemu bool

	fmt.Print("Masukkan ID workout : ")
	fmt.Scanln(&id)

	ketemu = false

	//SEQUENTIAL SEARCH
	for i = 0; i < n; i++ {
		if data[i].id == id {
			ketemu = true

			// Nama
			fmt.Print("Nama baru           : ")
			fmt.Scanln(&data[i].nama)

			// Jenis (validasi seperti tambah)
			var pilihanJenis int
			pilihanJenis = 0

			for pilihanJenis < 1 || pilihanJenis > 4 {
				fmt.Println("\nJenis Olahraga:")
				fmt.Println("1. Cardio")
				fmt.Println("2. Strength")
				fmt.Println("3. Flexibility")
				fmt.Println("4. HIIT")
				fmt.Print("Masukkan angka pilihan (1-4): ")
				fmt.Scanln(&pilihanJenis)

				if pilihanJenis < 1 || pilihanJenis > 4 {
					fmt.Println("Pilihan tidak valid, silakan coba lagi.")
				}
			}
			data[i].jenis = listJenis[pilihanJenis]

			// Durasi
			data[i].durasi = 0
			for data[i].durasi <= 0 {
				fmt.Print("Durasi baru > 0     : ")
				fmt.Scanln(&data[i].durasi)

				if data[i].durasi <= 0 {
					fmt.Println("Durasi harus lebih dari 0.")
				}
			}

			// Kalori
			data[i].kalori = 0
			for data[i].kalori <= 0 {
				fmt.Print("Kalori baru > 0     : ")
				fmt.Scanln(&data[i].kalori)

				if data[i].kalori <= 0 {
					fmt.Println("Kalori harus lebih dari 0.")
				}
			}

			// Tanggal
			fmt.Print("Tanggal baru        : ")
			fmt.Scanln(&data[i].tanggal)

			fmt.Println("Data berhasil diubah")
		}
	}

	if !ketemu {
		fmt.Println("Data tidak ditemukan")
	}
}

// ===========================Menu 4==========================
func hapusWorkout() {
	var id int
	var i int
	var j int
	var ketemu bool

	fmt.Print("Masukkan ID workout : ")
	fmt.Scanln(&id)

	ketemu = false

	for i = 0; i < n; i++ {
		if data[i].id == id {
			ketemu = true

			// geser array ke kiri
			for j = i; j < n-1; j++ {
				data[j] = data[j+1]
			}

			n--
			
			//hentikan loop
			i = n
		}
	}

	if ketemu {
		fmt.Println("Data berhasil dihapus")
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}


// =================Menu 5==========================
func CariJenisOlahragaSequential() {
	var jenis string
	var i int
	var ketemu bool

	fmt.Print("Cari jenis olahraga : ")
	fmt.Scanln(&jenis)

	ketemu = false

	//SEQUENTIAL SEARCH
	for i = 0; i < n; i++ {
		if data[i].jenis == jenis {
			if !ketemu {
				fmt.Println("Data ditemukan :")
			}

			fmt.Printf("%d | %s | %s | %d | %d | %s\n",
				data[i].id,
				data[i].nama,
				data[i].jenis,
				data[i].durasi,
				data[i].kalori,
				data[i].tanggal)

			ketemu = true
		}
	}

	if !ketemu {
		fmt.Println("Data tidak ditemukan")
	}
}


// =========================Menu 6===========================
func sortKalori() {
	var i, j int
	var key Workout
	var pilihan int

	fmt.Println("1. Ascending (kecil -> besar)")
	fmt.Println("2. Descending (besar -> kecil)")
	fmt.Print("Pilih: ")
	fmt.Scanln(&pilihan)

	// validasi dikit biar aman
	if pilihan != 1 && pilihan != 2 {
		fmt.Println("Pilihan tidak valid")
		return
	}

	// INSERTION SORT
	for i = 1; i < n; i++ {
		key = data[i]
		j = i - 1

		if pilihan == 1 { // ascending
			for j >= 0 && data[j].kalori > key.kalori {
				data[j+1] = data[j]
				j--
			}
		} else { // descending
			for j >= 0 && data[j].kalori < key.kalori {
				data[j+1] = data[j]
				j--
			}
		}

		data[j+1] = key
	}

	fmt.Println("Data berhasil di-sort berdasarkan kalori")

	tampilkanWorkout()
}

// SORTING JENIS (Selection Sort - Ascending) (INI BUKAN MENU)
func sortJenis() {
	var i, j, idxMin int
	var temp Workout

	for i = 0; i < n-1; i++ {
		idxMin = i // anggap elemen ke-i adalah yang paling kecil

		// cari index dengan jenis paling kecil di sisa array
		for j = i + 1; j < n; j++ {
			if data[j].jenis < data[idxMin].jenis {
				idxMin = j
			}
		}

		// tukar posisi
		temp = data[i]
		data[i] = data[idxMin]
		data[idxMin] = temp
	}

	fmt.Println("Data berhasil di-sort berdasarkan jenis (Selection Sort)")
}

// ===================Menu 7=========================
func CariJenisOlahragaBinary() {
	var jenis string
	var kiri, kanan, tengah int
	var ketemu bool

	fmt.Print("Cari jenis olahraga : ")
	fmt.Scanln(&jenis)

	// PENTING: sort dulu sebelum binary search
	sortJenis()

	kiri = 0
	kanan = n - 1
	ketemu = false

	for kiri <= kanan && !ketemu {
		tengah = (kiri + kanan) / 2

		if data[tengah].jenis == jenis {
			ketemu = true
		} else if data[tengah].jenis < jenis {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	if ketemu {
		fmt.Println("Data ditemukan :")

		// tampilkan SEMUA yang jenisnya sama
		
		// ke kiri
		i := tengah
		for i >= 0 && data[i].jenis == jenis {
			fmt.Printf("%d | %s | %s | %d | %d | %s\n",
				data[i].id,
				data[i].nama,
				data[i].jenis,
				data[i].durasi,
				data[i].kalori,
				data[i].tanggal)
			i--
		}

		// ke kanan
		i = tengah + 1
		for i < n && data[i].jenis == jenis {
			fmt.Printf("%d | %s | %s | %d | %d | %s\n",
				data[i].id,
				data[i].nama,
				data[i].jenis,
				data[i].durasi,
				data[i].kalori,
				data[i].tanggal)
			i++
		}

	} else {
		fmt.Println("Data tidak ditemukan")
	}
}



// ==========================Menu 8=====================
func aktivitasTerakhir() {
	var i int
	var mulai int

	if n <= 10 {
		mulai = 0
	} else {
		mulai = n - 10
	}

	fmt.Println("10 Aktivitas Terakhir")

	for i = mulai; i < n; i++ {
		fmt.Printf("%d | %s | %s | %d | %d | %s\n",
			data[i].id,
			data[i].nama,
			data[i].jenis,
			data[i].durasi,
			data[i].kalori,
			data[i].tanggal)
	}
}

// ==========================Menu 9=====================
func totalKalori() {
	var i int
	var total int

	total = 0

	for i = 0; i < n; i++ {
		total = total + data[i].kalori
	}

	fmt.Println("Total kalori terbakar =", total)
}

func menu() {
    fmt.Println()
    fmt.Println("===== APLIKASI TRACKING WORKOUT =====")
    fmt.Println("1. Tambah Workout")
    fmt.Println("2. Tampilkan Workout")
    fmt.Println("3. Ubah Workout")
    fmt.Println("4. Hapus Workout")
    fmt.Println("5. Cari Jenis Olahraga (Sequential)")
    fmt.Println("6. Sorting Kalori") 
    fmt.Println("7. Cari Jenis Olahraga (Binary)")
    fmt.Println("8. Aktivitas Terakhir")
    fmt.Println("9. Total Kalori")
    fmt.Println("0. Keluar")
    fmt.Println("======================================")
}

func main() {
	var pilih int

	for {
		menu()

		fmt.Print("Pilih menu : ")
		fmt.Scanln(&pilih)

		if pilih == 1 {
            tambahWorkout()
        } else if pilih == 2 {
            tampilkanWorkout()
        } else if pilih == 3 {
            ubahWorkout()
        } else if pilih == 4 {
            hapusWorkout()
        } else if pilih == 5 {
            CariJenisOlahragaSequential() 
        } else if pilih == 6 {
            sortKalori()
        } else if pilih == 7 {
            CariJenisOlahragaBinary()      
        } else if pilih == 8 {
            aktivitasTerakhir()
        } else if pilih == 9 {
            totalKalori()
        } else if pilih == 0 {
			fmt.Println("Program Selesai")
			return
		} else {
			fmt.Println("Pilihan tidak valid")
		}
	}
}
