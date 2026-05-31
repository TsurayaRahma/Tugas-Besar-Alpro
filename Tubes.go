package main

import "fmt"

const MAX int = 100
type Workout struct {
	id      int
	nama    string
	jenis   string
	durasi  int
	kalori  int
	tanggal string
	aktif   bool
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

//Menu 1
func tambahWorkout() {
	if n >= MAX {
		fmt.Println("Data penuh")
		return
	}

	var w Workout

	w.id = nextID
	nextID++

	fmt.Print("Nama latihan        : ")
	fmt.Scanln(&w.nama)

	var pilihanJenis int
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

	fmt.Print("Durasi (menit)      : ")
	fmt.Scanln(&w.durasi)

	fmt.Print("Jumlah kalori       : ")
	fmt.Scanln(&w.kalori)

	fmt.Print("Tanggal             : ")
	fmt.Scanln(&w.tanggal)

	w.aktif = true

	data[n] = w
	n++

	fmt.Println("Workout berhasil ditambahkan")
}

//Menu 2
func tampilkanWorkout() {
	var i int

	fmt.Println("==============================================================")
	fmt.Println("ID | Nama | Jenis | Durasi | Kalori | Tanggal")
	fmt.Println("==============================================================")

	for i = 0; i < n; i++ {
		if data[i].aktif {
			fmt.Printf("%d | %s | %s | %d | %d | %s\n",
				data[i].id,
				data[i].nama,
				data[i].jenis,
				data[i].durasi,
				data[i].kalori,
				data[i].tanggal)
		}
	}

	fmt.Println("==============================================================")
}

// Menu 3
func ubahWorkout() {
	var id int
	var i int
	var ketemu bool

	fmt.Print("Masukkan ID workout : ")
	fmt.Scanln(&id)

	ketemu = false

	for i = 0; i < n; i++ {
		if data[i].id == id && data[i].aktif {
			ketemu = true

			fmt.Print("Nama baru           : ")
			fmt.Scan(&data[i].nama)

			fmt.Print("Jenis baru          : ")
			fmt.Scan(&data[i].jenis)

			fmt.Print("Durasi baru         : ")
			fmt.Scan(&data[i].durasi)

			fmt.Print("Kalori baru         : ")
			fmt.Scan(&data[i].kalori)

			fmt.Print("Tanggal baru        : ")
			fmt.Scan(&data[i].tanggal)

			fmt.Println("Data berhasil diubah")
		}
	}

	if !ketemu {
		fmt.Println("Data tidak ditemukan")
	}
}

// Menu 4
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

			for j = i; j < n-1; j++ {
				data[j] = data[j+1]
			}

			n--
			fmt.Println("Data berhasil dihapus")
		}
	}

	if !ketemu {
		fmt.Println("Data tidak ditemukan")
	}
}


// Menu 5
func CariJenisOlahraga() {
	var jenis string
	var i int
	var ketemu bool

	fmt.Print("Cari jenis olahraga : ")
	fmt.Scanln(&jenis)

	ketemu = false

	for i = 0; i < n; i++ {
		if data[i].jenis == jenis && data[i].aktif {
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


// Menu 6
func sortJenis() {
	var i, j int
	var temp Workout

	for i = 0; i < n-1; i++ {
		for j = 0; j < n-1-i; j++ {
			if data[j].jenis > data[j+1].jenis {
				temp = data[j]
				data[j] = data[j+1]
				data[j+1] = temp
			}
		}
	}
}


// Menu 7
func binarySearch() {
	var jenis string
	var kiri, kanan, tengah int
	var ketemu bool

	fmt.Print("Cari jenis olahraga : ")
	fmt.Scanln(&jenis)

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

		fmt.Printf("%d | %s | %s | %d | %d | %s\n",
			data[tengah].id,
			data[tengah].nama,
			data[tengah].jenis,
			data[tengah].durasi,
			data[tengah].kalori,
			data[tengah].tanggal)

	} else {
		fmt.Println("Data tidak ditemukan")
	}
}


// Menu 8
func selectionSortDurasi() {
	var i, j, idx int
	var temp Workout

	for i = 0; i < n-1; i++ {
		idx = i

		for j = i + 1; j < n; j++ {
			if data[j].durasi < data[idx].durasi {
				idx = j
			}
		}

		temp = data[i]
		data[i] = data[idx]
		data[idx] = temp
	}

	fmt.Println("Data berhasil diurutkan berdasarkan durasi")
}


// Menu 9
func insertionSortKalori() {
	var i, j int
	var temp Workout

	for i = 1; i < n; i++ {
		temp = data[i]
		j = i - 1

		for j >= 0 && data[j].kalori > temp.kalori {
			data[j+1] = data[j]
			j--
		}

		data[j+1] = temp
	}

	fmt.Println("Data berhasil diurutkan berdasarkan kalori")
}


// Menu 10
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

func totalKalori() {
	var i int
	var total int

	total = 0

	for i = 0; i < n; i++ {
		if data[i].aktif {
			total = total + data[i].kalori
		}
	}

	fmt.Println("Total kalori terbakar =", total)
}

func menu() {
    fmt.Println()
    fmt.Println("===== APLIKASI WORKOUT =====")
    fmt.Println("1. Tambah Workout")
    fmt.Println("2. Tampilkan Workout")
    fmt.Println("3. Ubah Workout")
    fmt.Println("4. Hapus Workout")
    fmt.Println("5. Sequential Search (Cari Jenis)")
    fmt.Println("6. Urutkan Berdasarkan Jenis Olahraga") 
    fmt.Println("7. Binary Search (Cari Jenis Terurut)")
    fmt.Println("8. Sort Durasi")
    fmt.Println("9. Sort Kalori")
    fmt.Println("10. 10 Aktivitas Terakhir")
    fmt.Println("11. Total Kalori")
    fmt.Println("0. Keluar")
    fmt.Println("============================")
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
            CariJenisOlahraga() 
        } else if pilih == 6 {
            sortJenis()        
            fmt.Println("Data berhasil diurutkan berdasarkan jenis")
        } else if pilih == 7 {
            binarySearch()      
        } else if pilih == 8 {
            selectionSortDurasi()
        } else if pilih == 9 {
            insertionSortKalori()
        } else if pilih == 10 {
            aktivitasTerakhir()
        } else if pilih == 11 {
            totalKalori()
        } 
	}
}
