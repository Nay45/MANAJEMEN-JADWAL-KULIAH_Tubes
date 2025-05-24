package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Mahasiswa struct {
	NIM   int
	Nama  string
	Prodi string
}

type MataKuliah struct {
	KodeMK string
	NamaMK string
	SKS    int
	Hari   string
	Jam    string
}

type JadwalKuliah struct {
	NIM    int
	KodeMK string
	Hari   string
	Jam    string
}

var mahasiswa []Mahasiswa
var mataKuliah []MataKuliah
var jadwalKuliah []JadwalKuliah

var reader = bufio.NewReader(os.Stdin)

// =====================
// MENU UTAMA
// =====================
func main() {
	seedData()
	for {
		fmt.Println("\nManajemen Jadwal")
		fmt.Println("=======")
		fmt.Println("1. Data Mahasiswa")
		fmt.Println("2. Data Mata Kuliah")
		fmt.Println("3. Data Jadwal Kuliah")
		fmt.Println("4. Keluar")
		fmt.Println("=======")
		fmt.Print("Input: ")

		choice := readLine()
		switch choice {
		case "1":
			menuMahasiswa()
		case "2":
			menuMataKuliah()
		case "3":
			menuJadwalKuliah()
		case "4":
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// =====================
// SEEDING DATA DUMMY
// =====================
func seedData() {
	mahasiswa = []Mahasiswa{
		{103072400101, "Andi", "Informatika"},
		{103072400102, "Budi", "Sistem Informasi"},
		{103072400103, "Citra", "Teknik Komputer"},
		{103072400104, "Dewi", "Informatika"},
		{103072400105, "Eko", "Sistem Informasi"},
		{103072400106, "Fajar", "Teknik Komputer"},
		{103072400107, "Gita", "Informatika"},
		{103072400108, "Hadi", "Sistem Informasi"},
		{103072400109, "Intan", "Teknik Komputer"},
		{103072400110, "Joko", "Informatika"},
		{103072400111, "Kiki", "Sistem Informasi"},
		{103072400112, "Lina", "Teknik Komputer"},
		{103072400113, "Maya", "Informatika"},
		{103072400114, "Nina", "Sistem Informasi"},
		{103072400115, "Oka", "Teknik Komputer"},
		{103072400116, "Putu", "Informatika"},
		{103072400117, "Qori", "Sistem Informasi"},
		{103072400118, "Rina", "Teknik Komputer"},
		{103072400119, "Sari", "Informatika"},
		{103072400120, "Tono", "Sistem Informasi"},
		{103072400121, "Uli", "Teknik Komputer"},
		{103072400122, "Vina", "Informatika"},
		{103072400123, "Wawan", "Sistem Informasi"},
		{103072400124, "Xena", "Teknik Komputer"},
		{103072400125, "Yuni", "Informatika"},
		{103072400126, "Zaki", "Sistem Informasi"},
		{103072400127, "Adi", "Teknik Komputer"},
		{103072400128, "Bella", "Informatika"},
		{103072400129, "Caca", "Sistem Informasi"},
		{103072400130, "Dedi", "Teknik Komputer"},
	}

	mataKuliah = []MataKuliah{
		{"MK001", "Pemrograman Dasar", 3, "Senin", "08:00"},
		{"MK002", "Basis Data", 3, "Senin", "10:00"},
		{"MK003", "Algoritma", 3, "Selasa", "08:00"},
		{"MK004", "Jaringan Komputer", 3, "Selasa", "10:00"},
		{"MK005", "Sistem Operasi", 3, "Rabu", "08:00"},
		{"MK006", "Rekayasa Perangkat Lunak", 3, "Rabu", "10:00"},
		{"MK007", "Matematika Diskrit", 3, "Kamis", "08:00"},
		{"MK008", "Pemrograman Lanjut", 3, "Kamis", "10:00"},
		{"MK009", "Keamanan Komputer", 3, "Jumat", "08:00"},
		{"MK010", "Sistem Informasi", 3, "Jumat", "10:00"},
		{"MK011", "Data Mining", 3, "Senin", "13:00"},
		{"MK012", "Kecerdasan Buatan", 3, "Selasa", "13:00"},
		{"MK013", "Pemrograman Web", 3, "Rabu", "13:00"},
		{"MK014", "Grafika Komputer", 3, "Kamis", "13:00"},
		{"MK015", "Basis Data Lanjut", 3, "Jumat", "13:00"},
		{"MK016", "Pemrograman Mobile", 3, "Senin", "15:00"},
		{"MK017", "Pengolahan Citra", 3, "Selasa", "15:00"},
		{"MK018", "Komputasi Awan", 3, "Rabu", "15:00"},
		{"MK019", "Pemrograman Game", 3, "Kamis", "15:00"},
		{"MK020", "Internet of Things", 3, "Jumat", "15:00"},
		{"MK021", "Big Data", 3, "Senin", "17:00"},
		{"MK022", "Cloud Computing", 3, "Selasa", "17:00"},
		{"MK023", "Analisis Data", 3, "Rabu", "17:00"},
		{"MK024", "Manajemen Proyek", 3, "Kamis", "17:00"},
		{"MK025", "Sistem Terdistribusi", 3, "Jumat", "17:00"},
		{"MK026", "Robotika", 3, "Senin", "19:00"},
		{"MK027", "Pengembangan Web", 3, "Selasa", "19:00"},
		{"MK028", "Keamanan Jaringan", 3, "Rabu", "19:00"},
		{"MK029", "Basis Data NoSQL", 3, "Kamis", "19:00"},
		{"MK030", "Data Science", 3, "Jumat", "19:00"},
	}

	jadwalKuliah = []JadwalKuliah{
		{103072400101, "MK001", "Senin", "08:00"},
		{103072400101, "MK004", "Selasa", "10:00"},
		{103072400102, "MK002", "Senin", "10:00"},
		{103072400103, "MK003", "Selasa", "08:00"},
		{103072400104, "MK004", "Selasa", "10:00"},
		{103072400105, "MK005", "Rabu", "08:00"},
		{103072400106, "MK006", "Rabu", "10:00"},
		{103072400107, "MK007", "Kamis", "08:00"},
		{103072400108, "MK008", "Kamis", "10:00"},
		{103072400109, "MK009", "Jumat", "08:00"},
		{103072400110, "MK010", "Jumat", "10:00"},
		{103072400111, "MK011", "Senin", "13:00"},
		{103072400112, "MK012", "Selasa", "13:00"},
		{103072400113, "MK013", "Rabu", "13:00"},
		{103072400114, "MK014", "Kamis", "13:00"},
		{103072400115, "MK015", "Jumat", "13:00"},
		{103072400116, "MK016", "Senin", "15:00"},
		{103072400117, "MK017", "Selasa", "15:00"},
		{103072400118, "MK018", "Rabu", "15:00"},
		{103072400119, "MK019", "Kamis", "15:00"},
		{103072400120, "MK020", "Jumat", "15:00"},
		{103072400121, "MK021", "Senin", "17:00"},
		{103072400122, "MK022", "Selasa", "17:00"},
		{103072400123, "MK023", "Rabu", "17:00"},
		{103072400124, "MK024", "Kamis", "17:00"},
		{103072400125, "MK025", "Jumat", "17:00"},
		{103072400126, "MK026", "Senin", "19:00"},
		{103072400127, "MK027", "Selasa", "19:00"},
		{103072400128, "MK028", "Rabu", "19:00"},
		{103072400129, "MK029", "Kamis", "19:00"},
		{103072400130, "MK030", "Jumat", "19:00"},
	}
}

// =====================
// MENU MAHASISWA
// =====================
func menuMahasiswa() {
	for {
		fmt.Println("\n--- Data Mahasiswa ---")
		fmt.Println("a. Lihat Mahasiswa")
		fmt.Println("b. Tambah Mahasiswa")
		fmt.Println("c. Ubah Mahasiswa")
		fmt.Println("d. Hapus Mahasiswa")
		fmt.Println("e. Cari Mahasiswa")
		fmt.Println("f. Lihat Mahasiswa Secara Terurut")
		fmt.Println("g. Kembali")
		fmt.Print("Input: ")
		input := readLine()
		switch input {
		case "a":
			tampilMahasiswa()
		case "b":
			addMahasiswa()
		case "c":
			updateMahasiswa()
		case "d":
			deleteMahasiswa()
		case "e":
			searchMahasiswa()
		case "f":
			for {
				fmt.Print("Urutan (asc/desc): ")
				order := strings.ToLower(readLine())
				if order != "asc" && order != "desc" {
					fmt.Println("Input salah! Harap masukkan 'asc' atau 'desc'.")
					continue
				}
				sortMahasiswa(order == "asc")
				tampilMahasiswa()
				break
			}
		case "g":
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tampilMahasiswa() {
	fmt.Printf("\n%-20s %-20s %-20s\n", "NIM", "Nama", "Prodi")
	fmt.Println(strings.Repeat("-", 50))
	for _, m := range mahasiswa {
		fmt.Printf("%-20d %-20s %-20s\n", m.NIM, m.Nama, m.Prodi)
	}
}

func addMahasiswa() {
	for {
		fmt.Print("Masukkan NIM: ")
		nimStr := readLine()
		nim, err := strconv.Atoi(nimStr)
		if err != nil {
			fmt.Println("NIM harus berupa angka.")
			continue
		}

		duplikat := false
		for _, m := range mahasiswa {
			if m.NIM == nim {
				duplikat = true
				break
			}
		}
		if duplikat {
			fmt.Println("NIM sudah ada! Silakan masukkan NIM lain.")
			continue
		}

		fmt.Print("Masukkan Nama: ")
		nama := readLine()
		fmt.Print("Masukkan Prodi: ")
		prodi := readLine()
		mahasiswa = append(mahasiswa, Mahasiswa{nim, nama, prodi})
		fmt.Println("Mahasiswa berhasil ditambahkan")
		break
	}
}

func updateMahasiswa() {
	fmt.Print("Masukkan NIM yang akan diubah: ")
	nimStr := readLine()
	nim, err := strconv.Atoi(nimStr)
	if err != nil {
		fmt.Println("NIM harus berupa angka.")
		return
	}
	for i, m := range mahasiswa {
		if m.NIM == nim {
			fmt.Print("Masukkan nama baru: ")
			nama := readLine()
			fmt.Print("Masukkan prodi baru: ")
			prodi := readLine()
			mahasiswa[i].Nama = nama
			mahasiswa[i].Prodi = prodi
			fmt.Printf("Data mahasiswa dengan NIM %d berhasil di perbarui\n", nim)
			return
		}
	}
	fmt.Println("Data tidak ditemukan.")
}

func deleteMahasiswa() {
	fmt.Print("Masukkan NIM yang akan dihapus: ")
	nimStr := readLine()
	nim, err := strconv.Atoi(nimStr)
	if err != nil {
		fmt.Println("NIM harus berupa angka.")
		return
	}
	for i, m := range mahasiswa {
		if m.NIM == nim {
			mahasiswa = append(mahasiswa[:i], mahasiswa[i+1:]...)
			fmt.Printf("Data mahasiswa dengan NIM %d berhasil di hapus\n", nim)
			return
		}
	}
	fmt.Println("Data tidak ditemukan.")
}

func searchMahasiswa() {
	fmt.Print("Masukkan nama mahasiswa: ")
	nama := readLine()
	found := false
	fmt.Printf("\n%-20s %-20s %-20s\n", "NIM", "Nama", "Prodi")
	fmt.Println(strings.Repeat("-", 50))
	for _, m := range mahasiswa {
		if strings.HasPrefix(strings.ToLower(m.Nama), strings.ToLower(nama)) {
			fmt.Printf("%-20d %-20s %-20s\n", m.NIM, m.Nama, m.Prodi)
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ditemukan.")
	}
}

func sortMahasiswa(asc bool) {
	for i := 0; i < len(mahasiswa)-1; i++ {
		for j := 0; j < len(mahasiswa)-i-1; j++ {
			if (asc && mahasiswa[j].NIM > mahasiswa[j+1].NIM) || (!asc && mahasiswa[j].NIM < mahasiswa[j+1].NIM) {
				mahasiswa[j], mahasiswa[j+1] = mahasiswa[j+1], mahasiswa[j]
			}
		}
	}
}

// =====================
// MENU MATA KULIAH
// =====================
func menuMataKuliah() {
	for {
		fmt.Println("\n--- Data Mata Kuliah ---")
		fmt.Println("a. Lihat")
		fmt.Println("b. Tambah")
		fmt.Println("c. Ubah")
		fmt.Println("d. Hapus")
		fmt.Println("e. Cari")
		fmt.Println("f. Lihat Terurut (asc/desc)")
		fmt.Println("g. Kembali")
		fmt.Print("Input: ")
		input := readLine()
		switch input {
		case "a":
			tampilMataKuliah()
		case "b":
			addMataKuliah()
		case "c":
			updateMataKuliah()
		case "d":
			deleteMataKuliah()
		case "e":
			searchMataKuliah()
		case "f":
			for {
				fmt.Print("Urutan (asc/desc): ")
				order := strings.ToLower(readLine())
				if order != "asc" && order != "desc" {
					fmt.Println("Input salah! Harap masukkan 'asc' atau 'desc'.")
					continue
				}
				sortMataKuliah(order == "asc")
				tampilMataKuliah()
				break
			}
		case "g":
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tampilMataKuliah() {
	fmt.Printf("\n%-10s %-30s %-5s %-10s %-10s\n", "KodeMK", "NamaMK", "SKS", "Hari", "Jam")
	fmt.Println(strings.Repeat("-", 60))
	for _, mk := range mataKuliah {
		fmt.Printf("%-10s %-30s %-5d %-10s %-10s\n", mk.KodeMK, mk.NamaMK, mk.SKS, mk.Hari, mk.Jam)
	}
}

func addMataKuliah() {
	for {
		fmt.Print("Masukkan Kode MK: ")
		kode := readLine()

		duplikat := false
		for _, m := range mataKuliah {
			if m.KodeMK == kode {
				duplikat = true
				break
			}
		}
		if duplikat {
			fmt.Println("Kode MK sudah ada! Silakan masukkan Kode MK lain.")
			continue
		}

		fmt.Print("Masukkan nama MK: ")
		nama := readLine()
		fmt.Print("Masukkan SKS: ")
		sksStr := readLine()
		sks, err := strconv.Atoi(sksStr)
		if err != nil {
			fmt.Println("SKS harus berupa angka.")
			continue
		}
		fmt.Print("Masukkan Hari (cth: Senin): ")
		hari := readLine()
		fmt.Print("Masukkan Jam (cth: 08:00): ")
		jam := readLine()

		mataKuliah = append(mataKuliah, MataKuliah{kode, nama, sks, hari, jam})

		fmt.Println("Mata Kuliah berhasil ditambahkan")
		break
	}
}

func updateMataKuliah() {
	fmt.Print("Masukkan KodeMK yang akan diubah: ")
	kode := readLine()
	for i, mk := range mataKuliah {
		if mk.KodeMK == kode {
			fmt.Print("Masukkan Nama baru: ")
			nama := readLine()
			fmt.Print("Masukkan SKS baru: ")
			sksStr := readLine()
			sks, err := strconv.Atoi(sksStr)
			if err != nil {
				fmt.Println("SKS harus berupa angka.")
				return
			}
			fmt.Print("Masukkan Hari baru (cth: Senin): ")
			hari := readLine()
			fmt.Print("Masukkan Jam baru (cth: 08:00): ")
			jam := readLine()

			mataKuliah[i] = MataKuliah{kode, nama, sks, hari, jam}
			fmt.Printf("Data mata kuliah dengan kode %s berhasil di perbarui\n", kode)
			return
		}
	}
	fmt.Println("Data tidak ditemukan.")
}

func deleteMataKuliah() {
	fmt.Print("Masukkan KodeMK yang akan dihapus: ")
	kode := readLine()
	for i, mk := range mataKuliah {
		if mk.KodeMK == kode {
			mataKuliah = append(mataKuliah[:i], mataKuliah[i+1:]...)
			fmt.Printf("Data mata kuliah dengan kode %s berhasil di hapus\n", kode)
			return
		}
	}
	fmt.Println("Data tidak ditemukan.")
}

func searchMataKuliah() {
	fmt.Print("Masukkan nama mata kuliah: ")
	nama := readLine()
	found := false
	fmt.Printf("\n%-10s %-30s %-5s %-10s %-10s\n", "KodeMK", "NamaMK", "SKS", "Hari", "Jam")
	fmt.Println(strings.Repeat("-", 60))
	for _, mk := range mataKuliah {
		if strings.Contains(strings.ToLower(mk.NamaMK), strings.ToLower(nama)) {
			fmt.Printf("%-10s %-30s %-5d %-10s %-10s\n", mk.KodeMK, mk.NamaMK, mk.SKS, mk.Hari, mk.Jam)
			found = true
		}
	}

	if !found {
		fmt.Println("Tidak ditemukan.")
	}
}

func sortMataKuliah(asc bool) {
	for i := 0; i < len(mataKuliah)-1; i++ {
		for j := 0; j < len(mataKuliah)-i-1; j++ {
			if (asc && mataKuliah[j].KodeMK > mataKuliah[j+1].KodeMK) || (!asc && mataKuliah[j].KodeMK < mataKuliah[j+1].KodeMK) {
				mataKuliah[j], mataKuliah[j+1] = mataKuliah[j+1], mataKuliah[j]
			}
		}
	}
}

// =====================
// MENU JADWAL
// =====================
func menuJadwalKuliah() {
	for {
		fmt.Println("\n--- Data Jadwal Kuliah ---")
		fmt.Println("a. Lihat")
		fmt.Println("b. Tambah")
		fmt.Println("c. Ubah")
		fmt.Println("d. Hapus")
		fmt.Println("e. Lihat Terurut (asc/desc)")
		fmt.Println("f. Cari (berdasarkan NIM/KodeMK)")
		fmt.Println("g. Kembali")
		fmt.Print("Input: ")
		input := readLine()
		switch input {
		case "a":
			tampilJadwal()
		case "b":
			addJadwal()
		case "c":
			updateJadwal()
		case "d":
			deleteJadwal()
		case "e":
			for {
				fmt.Print("Urutan (asc/desc): ")
				order := strings.ToLower(readLine())
				if order != "asc" && order != "desc" {
					fmt.Println("Input salah! Harap masukkan 'asc' atau 'desc'.")
					continue
				}
				sortJadwal(order == "asc")
				tampilJadwal()
				break
			}
		case "f":
			searchJadwal()
		case "g":
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tampilJadwal() {
	fmt.Printf("\n%-20s %-10s %-10s %-10s\n", "NIM", "KodeMK", "Hari", "Jam")
	fmt.Println(strings.Repeat("-", 40))
	for _, jd := range jadwalKuliah {
		fmt.Printf("%-20d %-10s %-10s %-10s\n", jd.NIM, jd.KodeMK, jd.Hari, jd.Jam)
	}
}

func addJadwal() {
	for {
		fmt.Print("Masukkan NIM Mahasiswa: ")
		nimStr := readLine()
		nim, err := strconv.Atoi(nimStr)
		if err != nil {
			fmt.Println("NIM harus berupa angka.")
			continue
		}

		validNIM := false
		for _, mhs := range mahasiswa {
			if mhs.NIM == nim {
				validNIM = true
				break
			}
		}
		if !validNIM {
			fmt.Println("NIM mahasiswa tidak ditemukan. Silakan masukkan NIM yang valid.")
			continue
		}

		fmt.Print("Masukkan KodeMK: ")
		kode := readLine()

		var hari, jam string
		found := false
		for _, mk := range mataKuliah {
			if mk.KodeMK == kode {
				hari = mk.Hari
				jam = mk.Jam
				found = true
				break
			}
		}

		if !found {
			fmt.Println("KodeMK tidak ditemukan. Silakan masukkan kode yang valid.")
			continue
		}

		duplicate := false
		for _, jd := range jadwalKuliah {
			if jd.NIM == nim && jd.KodeMK == kode {
				duplicate = true
				break
			}
		}
		if duplicate {
			fmt.Println("Jadwal untuk NIM dan KodeMK ini sudah ada.")
			continue
		}

		bentrok := false
		for _, jd := range jadwalKuliah {
			if jd.NIM == nim && jd.Hari == hari && jd.Jam == jam {
				bentrok = true
				break
			}
		}
		if bentrok {
			fmt.Println("Jadwal bentrok! Mahasiswa sudah memiliki jadwal pada hari dan jam yang sama.")
			continue
		}

		jadwalKuliah = append(jadwalKuliah, JadwalKuliah{nim, kode, hari, jam})
		fmt.Printf("Jadwal berhasil ditambahkan: NIM %d, KodeMK %s, Hari %s, Jam %s\n", nim, kode, hari, jam)
		break
	}
}

func updateJadwal() {
	fmt.Print("Masukkan NIM: ")
	nimStr := readLine()
	nim, err := strconv.Atoi(nimStr)
	if err != nil {
		fmt.Println("NIM harus berupa angka.")
		return
	}

	// Cek validitas NIM di data mahasiswa
	validNIM := false
	for _, mhs := range mahasiswa {
		if mhs.NIM == nim {
			validNIM = true
			break
		}
	}
	if !validNIM {
		fmt.Println("NIM mahasiswa tidak ditemukan.")
		return
	}

	fmt.Print("Masukkan KodeMK lama: ")
	kodeLama := readLine()

	// Cek validitas KodeMK lama di data mataKuliah
	validKodeLama := false
	for _, mk := range mataKuliah {
		if mk.KodeMK == kodeLama {
			validKodeLama = true
			break
		}
	}
	if !validKodeLama {
		fmt.Println("KodeMK lama tidak ditemukan.")
		return
	}

	// Cari indeks jadwal berdasarkan NIM dan KodeMK lama
	index := -1
	for i, jd := range jadwalKuliah {
		if jd.NIM == nim && jd.KodeMK == kodeLama {
			index = i
			break
		}
	}
	if index == -1 {
		fmt.Println("Jadwal dengan NIM dan KodeMK lama tidak ditemukan.")
		return
	}

	fmt.Print("Masukkan KodeMK baru: ")
	kodeBaru := readLine()

	// Cek validitas KodeMK baru di data mataKuliah dan dapatkan hari & jam
	var hariBaru, jamBaru string
	validKodeBaru := false
	for _, mk := range mataKuliah {
		if mk.KodeMK == kodeBaru {
			validKodeBaru = true
			hariBaru = mk.Hari
			jamBaru = mk.Jam
			break
		}
	}
	if !validKodeBaru {
		fmt.Println("KodeMK baru tidak ditemukan.")
		return
	}

	// Cek duplikat jadwal (NIM + KodeMK baru), kecuali jadwal yang sedang diupdate
	for i, jd := range jadwalKuliah {
		if i != index && jd.NIM == nim && jd.KodeMK == kodeBaru {
			fmt.Println("Jadwal dengan NIM dan KodeMK baru sudah ada.")
			return
		}
	}

	// Cek bentrok jadwal (hari & jam baru) untuk NIM yang sama, kecuali jadwal yang sedang diupdate
	for i, jd := range jadwalKuliah {
		if i != index && jd.NIM == nim && jd.Hari == hariBaru && jd.Jam == jamBaru {
			fmt.Println("Jadwal bentrok! Mahasiswa sudah memiliki jadwal pada hari dan jam yang sama.")
			return
		}
	}

	// Update jadwal
	jadwalKuliah[index] = JadwalKuliah{nim, kodeBaru, hariBaru, jamBaru}
	fmt.Printf("Jadwal berhasil diperbarui: NIM %d, KodeMK Lama %s, KodeMK Baru %s\n", nim, kodeLama, kodeBaru)
}

func deleteJadwal() {
	fmt.Print("Masukkan NIM untuk jadwal yang dihapus: ")
	nimStr := readLine()
	nim, err := strconv.Atoi(nimStr)
	if err != nil {
		fmt.Println("NIM harus berupa angka.")
		return
	}

	found := false
	for _, jd := range jadwalKuliah {
		if jd.NIM == nim {
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Tidak ditemukan jadwal dengan NIM tersebut.")
		return
	}

	fmt.Print("Masukkan KodeMK yang ingin dihapus (atau ketik 'hapus semua' untuk menghapus semua jadwal dari NIM ini): ")
	kodeMK := readLine()

	if kodeMK == "hapus semua" {
		newJadwal := []JadwalKuliah{}
		for _, jd := range jadwalKuliah {
			if jd.NIM != nim {
				newJadwal = append(newJadwal, jd)
			}
		}
		jadwalKuliah = newJadwal
		fmt.Printf("Semua jadwal dari NIM %d berhasil dihapus\n", nim)
		return
	}

	deleted := false
	for i := 0; i < len(jadwalKuliah); i++ {
		if jadwalKuliah[i].NIM == nim && jadwalKuliah[i].KodeMK == kodeMK {
			jadwalKuliah = append(jadwalKuliah[:i], jadwalKuliah[i+1:]...)
			deleted = true
			fmt.Printf("Jadwal dari NIM %d dengan kode MK %s berhasil dihapus\n", nim, kodeMK)
			break
		}
	}

	if !deleted {
		fmt.Println("KodeMK tidak ditemukan untuk NIM tersebut.")
	}
}

func searchJadwal() {
	fmt.Print("Cari berdasarkan NIM/KodeMK: ")
	query := readLine()
	found := false
	fmt.Printf("\n%-20s %-10s %-10s %-10s\n", "NIM", "KodeMK", "Hari", "Jam")
	fmt.Println(strings.Repeat("-", 40))
	for _, jd := range jadwalKuliah {
		if strings.Contains(strconv.Itoa(jd.NIM), query) || strings.Contains(strings.ToLower(jd.KodeMK), strings.ToLower(query)) {
			fmt.Printf("%-20d %-10s %-10s %-10s\n", jd.NIM, jd.KodeMK, jd.Hari, jd.Jam)
			found = true
		}
	}

	if !found {
		fmt.Println("Tidak ditemukan.")
	}
}

func sortJadwal(asc bool) {
	for i := 0; i < len(jadwalKuliah)-1; i++ {
		for j := 0; j < len(jadwalKuliah)-i-1; j++ {
			if (asc && jadwalKuliah[j].NIM > jadwalKuliah[j+1].NIM) || (!asc && jadwalKuliah[j].NIM < jadwalKuliah[j+1].NIM) {
				jadwalKuliah[j], jadwalKuliah[j+1] = jadwalKuliah[j+1], jadwalKuliah[j]
			}
		}
	}
}

func readLine() string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
