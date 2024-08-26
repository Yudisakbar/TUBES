package main

import "fmt"

const NMAX int = 1000

type TabPenerimaan struct {
	NamaMahasiswa, NomorInduk                           string
	KotaTinggal, NamaJalan                              string
	TanggalLahir, BulanLahir, TahunLahir, UmurMahasiswa int
}

type TabRegis struct {
	Program, JalurSeleksi, Jurusan string
	NilaiTes                       float64
}

type KonfirmasiAdmin struct {
	Konfirmasi bool
}

type TabArr [NMAX]TabPenerimaan
type TabArr2 [NMAX]TabRegis
type TabKonf [NMAX]KonfirmasiAdmin

func main() {
	var arr TabArr
	var arr2 TabArr2
	var arr3 TabKonf
	var banyakMahasiswa int
	var chooseMenu, chooseMenuAdmin, chooseMenuUser int

	banyakMahasiswa = 0

	for {
		MenuAdminUser()
		fmt.Print("Pilih Menu: ")
		fmt.Scan(&chooseMenu)
		fmt.Println()
		if chooseMenu > 2 {
			break
		}
		switch chooseMenu {
		case 1:
			MenuAdmin()
			fmt.Print("Pilih Menu: ")
			fmt.Scan(&chooseMenuAdmin)
			fmt.Println()
			switch chooseMenuAdmin {
			case 1:
				CekPendaftar(arr, arr2, banyakMahasiswa, arr3)
				fmt.Println()
			case 2:
				BeriKonfirmasi(arr, arr2, banyakMahasiswa, &arr3)
				fmt.Println()
			case 3:
				EditMahasiswa(&arr, &arr2, &banyakMahasiswa, &arr3)
				fmt.Println()
			}
		case 2:
			MenuUser()
			fmt.Print("Pilih Menu: ")
			fmt.Scan(&chooseMenuUser)
			fmt.Println()
			switch chooseMenuUser {
			case 1:
				IsInput(&arr, &arr2, &banyakMahasiswa)
				fmt.Println()
			case 2:
				var CariNama string
				fmt.Print("Masukan Nama Anda: ")
				fmt.Scan(&CariNama)
				fmt.Println()
				CekStatusMahasiswa(arr, banyakMahasiswa, CariNama)
				fmt.Println()
			case 3:
				var CariNama string
				fmt.Print("Masukan Nama Anda: ")
				fmt.Scan(&CariNama)
				fmt.Println()
				LihatStatusKelulusan(arr, arr2, banyakMahasiswa, arr3, CariNama)
				fmt.Println()
			}
		}
	}
}


func CekPendaftar(arr TabArr, arr2 TabArr2, banyakMahasiswa int, arr3 TabKonf) {
	var sorting int
	for i := 0; i < banyakMahasiswa; i++ {
		fmt.Println("==========================================")
		fmt.Printf(" Mahasiswa ke %d: 				\n", i+1)
		fmt.Printf(" Nama		: %s 		\n", arr[i].NamaMahasiswa)
		fmt.Printf(" Nomor Induk	: %s 			\n", arr[i].NomorInduk)
		fmt.Printf(" Tanggal Lahir	: %d-%d-%d 			\n", arr[i].TanggalLahir, arr[i].BulanLahir, arr[i].TahunLahir)
		fmt.Printf(" Umur		: %d 				\n", arr[i].UmurMahasiswa)
		fmt.Printf(" Asal Kota	: %s 			\n", arr[i].KotaTinggal)
		fmt.Printf(" Alamat		: %s 		\n", arr[i].NamaJalan)
		fmt.Printf(" Program Studi	: %s %s\n", arr2[i].Program, arr2[i].Jurusan)
		fmt.Printf(" Nilai Tes 	: %0.1f\n", arr2[i].NilaiTes)
		fmt.Println("==========================================")
		fmt.Println()
	}
	MenuSortingAdmin()
	fmt.Print("Pilih Menu: ")
	fmt.Scan(&sorting)
	fmt.Println()
	switch sorting {
	case 1:
		NilaiAscending(&arr, &arr2, banyakMahasiswa)
	case 2:
		NilaiDescending(&arr, &arr2, banyakMahasiswa)
	case 3:
		SortingJurusan(arr, arr2, banyakMahasiswa)
	case 4:
		SortingLulus(arr, arr3, banyakMahasiswa)
	case 5:
		SortingDitolak(arr, arr3, banyakMahasiswa)
	case 6:
		SortingUmur(arr, banyakMahasiswa)
	}
}

func BeriKonfirmasi(arr TabArr, arr2 TabArr2, banyakMahasiswa int, arr3 *TabKonf) {
	var kelulusan string
	for i := 0; i < banyakMahasiswa; i++ {
		fmt.Printf(" Mahasiswa ke %d: 				\n", i+1)
		fmt.Printf(" Nama		: %s		\n", arr[i].NamaMahasiswa)
		fmt.Printf(" Nomor Induk	: %s 			\n", arr[i].NomorInduk)
		fmt.Printf(" Program Studi	: %s %s\n", arr2[i].Program, arr2[i].Jurusan)
		fmt.Printf(" Nilai Tes 	: %0.1f\n", arr2[i].NilaiTes)
		fmt.Println()
		fmt.Print(" Status Kelulusan (Ya/Tidak): ")
		fmt.Scan(&kelulusan)
		if kelulusan == "Ya" {
			arr3[i].Konfirmasi = true
		} else {
			arr3[i].Konfirmasi = false
		}
	}
}

func EditMahasiswa(arr *TabArr, arr2 *TabArr2, banyakMahasiswa *int, arr3 *TabKonf) {
	/* Mengedit mahasiswa yang sudah terdaftar dengan memasukan Nama, NIK, Tanggal lahir, Umur, Asal kota, Alamat yang baru */
	var menuEdit string
	var edit, nomorIndukEdit, kotaTinggalEdit string
	var namaJalanEdit string
	var tanggalLahirEdit, bulanLahirEdit, tahunLahirEdit, umurMahasiswaEdit int
	var noIndukHapus string
	var kelulusan string
	var programStud int
	var nilaiTesEdit float64
	var namaHapus string

	fmt.Print("Menu: (tambah / edit / hapus): ")
	fmt.Scan(&menuEdit)
	fmt.Println()
	switch menuEdit {
	case "tambah":
		i := *banyakMahasiswa
		fmt.Printf("Nama			     : ")
		fmt.Scan(&arr[i].NamaMahasiswa)

		fmt.Printf("Nomor Induk Kependudukan     : ")
		fmt.Scan(&arr[i].NomorInduk)

		fmt.Printf("Tanggal Lahir (Ex : 8 9 2004): ")
		fmt.Scan(&arr[i].TanggalLahir, &arr[i].BulanLahir, &arr[i].TahunLahir)

		fmt.Print("Umur			     : ")
		fmt.Scan(&arr[i].UmurMahasiswa)

		fmt.Printf("Asal Kota (Ex : Bandung)     : ")
		fmt.Scan(&arr[i].KotaTinggal)

		fmt.Printf("Alamat Rumah 		     : ")
		fmt.Scan(&arr[i].NamaJalan)
		fmt.Println()
		MenuJurusan()
		fmt.Print("Pilih Jurusan: ")
		fmt.Scan(&programStud)
		switch programStud {
		case 1:
			arr2[i].Program = "S1"
			arr2[i].Jurusan = "Informatika"
		case 2:
			arr2[i].Program = "S1"
			arr2[i].Jurusan = "Sistem Informasi"
		case 3:
			arr2[i].Program = "D4"
			arr2[i].Jurusan = "Informatika"
		case 4:
			arr2[i].Program = "D4"
			arr2[i].Jurusan = "Sistem Informasi"
		case 5:
			arr2[i].Program = "D3"
			arr2[i].Jurusan = "Informatika"
		case 6:
			arr2[i].Program = "D3"
			arr2[i].Jurusan = "Sistem Informasi"
		}

		fmt.Print("Masukan nilai anda: ")
		fmt.Scan(&arr2[i].NilaiTes)
		fmt.Println()
		fmt.Print("Status Kelulusan (Ya/Tidak): ")
		fmt.Scan(&kelulusan)
		if kelulusan == "Ya" {
			arr3[i].Konfirmasi = true
		} else {
			arr3[i].Konfirmasi = false
		}

		i += 1
		fmt.Println()
		*banyakMahasiswa = i
	case "edit":
		fmt.Print("Nama mahasiswa yang akan diganti :  ")
		fmt.Scan(&edit)
		fmt.Println()
		fmt.Printf("Masukan data baru\n")
		for i := 0; i < *banyakMahasiswa; i++ {
			if arr[i].NamaMahasiswa == edit {
				fmt.Printf("Nama		                 : ")
				fmt.Scan(&edit)

				fmt.Printf("Nomor Induk Kependudukan         : ")
				fmt.Scan(&nomorIndukEdit)

				fmt.Printf("Tanggal Lahir (Ex : 8 9 2004)    : ")
				fmt.Scan(&tanggalLahirEdit, &bulanLahirEdit, &tahunLahirEdit)

				fmt.Print("Umur			         : ")
				fmt.Scan(&umurMahasiswaEdit)

				fmt.Printf("Asal Kota (Ex : Bandung)         : ")
				fmt.Scan(&kotaTinggalEdit)

				fmt.Printf("Alamat Rumah 		         : ")
				fmt.Scan(&namaJalanEdit)
				fmt.Println()
				MenuJurusan()
				fmt.Print("Pilih Jurusan: ")
				fmt.Scan(&programStud)
				switch programStud {
				case 1:
					arr2[i].Program = "S1"
					arr2[i].Jurusan = "Informatika"
				case 2:
					arr2[i].Program = "S1"
					arr2[i].Jurusan = "Sistem Informasi"
				case 3:
					arr2[i].Program = "D4"
					arr2[i].Jurusan = "Informatika"
				case 4:
					arr2[i].Program = "D4"
					arr2[i].Jurusan = "Sistem Informasi"
				case 5:
					arr2[i].Program = "D3"
					arr2[i].Jurusan = "Informatika"
				case 6:
					arr2[i].Program = "D3"
					arr2[i].Jurusan = "Sistem Informasi"
				}

				fmt.Print("Masukan nilai anda: ")
				fmt.Scan(&nilaiTesEdit)
				fmt.Println()
				fmt.Print("Status Kelulusan (Ya/Tidak): ")
				fmt.Scan(&kelulusan)
				if kelulusan == "Ya" {
					arr3[i].Konfirmasi = true
				} else {
					arr3[i].Konfirmasi = false
				}
				arr[i].NamaMahasiswa = edit
				arr[i].NomorInduk = nomorIndukEdit
				arr[i].KotaTinggal = kotaTinggalEdit
				arr[i].TanggalLahir = tahunLahirEdit
				arr[i].BulanLahir = bulanLahirEdit
				arr[i].TahunLahir = tahunLahirEdit
				arr[i].NamaJalan = namaJalanEdit
				arr[i].UmurMahasiswa = umurMahasiswaEdit
				arr2[i].NilaiTes = nilaiTesEdit
			}
		}
	case "hapus":
		fmt.Print("Ketikan Nama dan Nomor Induk yang akan dihapus: ")
		fmt.Scan(&namaHapus, &noIndukHapus)
		hasilFind := FindDelete(*arr, *banyakMahasiswa, namaHapus, noIndukHapus) // Memanggil fungsi find delete untuk mencari index data yang akan dihapus //

		if hasilFind != -1 {
			for i := hasilFind; i < *banyakMahasiswa; i++ {
				arr[i] = arr[i+1]
				arr2[i] = arr2[i+1]
			}
			fmt.Println("Penghapusan berhasil !")
			*banyakMahasiswa = *banyakMahasiswa - 1
			fmt.Println("Data hasil penghapusan: ")
			for i := 0; i < *banyakMahasiswa; i++ {
				fmt.Println("==========================================")
				fmt.Printf(" Mahasiswa ke %d: 				\n", i+1)
				fmt.Printf(" Nama		: %s		\n", arr[i].NamaMahasiswa)
				fmt.Printf(" Nomor Induk	: %s 			\n", arr[i].NomorInduk)
				fmt.Printf(" Tanggal Lahir	: %d-%d-%d 			\n", arr[i].TanggalLahir, arr[i].BulanLahir, arr[i].TahunLahir)
				fmt.Printf(" Umur		: %d 				\n", arr[i].UmurMahasiswa)
				fmt.Printf(" Asal Kota	: %s 			\n", arr[i].KotaTinggal)
				fmt.Printf(" Alamat		: %s  		\n", arr[i].NamaJalan)
				fmt.Printf(" Program Studi	: %s %s\n", arr2[i].Program, arr2[i].Jurusan)
				fmt.Printf(" Nilai Tes 	: %0.1f\n", arr2[i].NilaiTes)
				fmt.Println("==========================================")
				fmt.Println()
			}
		} else {
			fmt.Println("Penghapusan gagal karena tidak ditemukan data yang cocok")
		}
	}
}

func FindDelete(arr TabArr, banyakMahasiswa int, x, y string) int {
	// { Fungsi mengembalikan index data yang akan dihapus } //
	idx := -1
	i := 0
	for i < banyakMahasiswa && idx == -1 {
		if arr[i].NomorInduk == y && arr[i].NamaMahasiswa == x {
			idx = i
		}
		i += 1
	}
	return idx
}

func NilaiAscending(arr *TabArr, arr2 *TabArr2, banyakMahasiswa int) {
	for pass := 1; pass < banyakMahasiswa; pass++ {
		idx := pass - 1
		for j := pass; j < banyakMahasiswa; j++ {
			if (*arr2)[idx].NilaiTes > (*arr2)[j].NilaiTes {
				idx = j
			}
		}
		temp := (*arr2)[idx].NilaiTes
		temp1 := (*arr)[idx].NamaMahasiswa
		temp2 := (*arr)[idx].NomorInduk

		(*arr2)[idx].NilaiTes = (*arr2)[pass-1].NilaiTes
		(*arr)[idx].NamaMahasiswa = (*arr)[pass-1].NamaMahasiswa
		(*arr)[idx].NomorInduk = (*arr)[pass-1].NomorInduk

		(*arr2)[pass-1].NilaiTes = temp
		(*arr)[pass-1].NamaMahasiswa = temp1
		(*arr)[pass-1].NomorInduk = temp2
	}

	fmt.Println("========================================")
	for k := 0; k < banyakMahasiswa; k++ {
		fmt.Printf("Nama	 : %s\nNo.Induk : %s\nNilai	 : %0.2f\n", arr[k].NamaMahasiswa, arr[k].NomorInduk, arr2[k].NilaiTes)
		fmt.Println("========================================")
	}
}

func NilaiDescending(arr *TabArr, arr2 *TabArr2, banyakMahasiswa int) {
	for pass := 1; pass < banyakMahasiswa; pass++ {
		j := pass - 1
		temp := (*arr2)[pass].NilaiTes
		temp1 := (*arr)[pass].NamaMahasiswa
		temp2 := (*arr)[pass].NomorInduk
		for j >= 0 && temp > (*arr2)[j].NilaiTes {
			(*arr)[j+1].NamaMahasiswa = (*arr)[j].NamaMahasiswa
			(*arr)[j+1].NomorInduk = (*arr)[j].NomorInduk
			(*arr2)[j+1].NilaiTes = (*arr2)[j].NilaiTes
			j--
		}
		(*arr2)[j+1].NilaiTes = temp
		(*arr)[j+1].NamaMahasiswa = temp1
		(*arr)[j+1].NomorInduk = temp2
	}

	fmt.Println("========================================")
	for k := 0; k < banyakMahasiswa; k++ {
		fmt.Printf("Nama	 : %s\nNo.Induk : %s\nNilai	 : %0.2f\n", arr[k].NamaMahasiswa, arr[k].NomorInduk, arr2[k].NilaiTes)
		fmt.Println("========================================")
	}
}

func SortingJurusan(arr TabArr, arr2 TabArr2, banyakMahasiswa int) {
	var Jurusan int
	MenuJurusan()
	fmt.Print("Pilih Menu Program dan Jurusan: ")
	fmt.Scan(&Jurusan)
	fmt.Println()
	switch Jurusan {
	case 1:
		fmt.Println("========================================")
		for k := 0; k < banyakMahasiswa; k++ {
			if arr2[k].Program == "S1" && arr2[k].Jurusan == "Informatika" {
				fmt.Printf("Nama		: %s\nNomor Induk	: %s\nProgram Studi	: S1 Informatika\n", arr[k].NamaMahasiswa, arr[k].NomorInduk)
				fmt.Println("========================================")
			}
		}

	case 2:
		fmt.Println("========================================")
		for k := 0; k < banyakMahasiswa; k++ {
			if arr2[k].Program == "S1" && arr2[k].Jurusan == "Sistem Informasi" {
				fmt.Printf("Nama		: %s\nNomor Induk	: %s\nProgram Studi	: S1 Sistem Informasi\n", arr[k].NamaMahasiswa, arr[k].NomorInduk)
				fmt.Println("========================================")
			}
		}

	case 3:
		fmt.Println("========================================")
		for k := 0; k < banyakMahasiswa; k++ {
			if arr2[k].Program == "D4" && arr2[k].Jurusan == "Informatika" {
				fmt.Printf("Nama		: %s\nNomor Induk	: %s\nProgram Studi	: D4 Informatika\n", arr[k].NamaMahasiswa, arr[k].NomorInduk)
				fmt.Println("========================================")
			}
		}

	case 4:
		fmt.Println("========================================")
		for k := 0; k < banyakMahasiswa; k++ {
			if arr2[k].Program == "D4" && arr2[k].Jurusan == "Sistem Informasi" {
				fmt.Printf("Nama		: %s\nNomor Induk	: %s\nProgram Studi	: D4 Sistem Informasi\n", arr[k].NamaMahasiswa, arr[k].NomorInduk)
				fmt.Println("========================================")
			}
		}

	case 5:
		fmt.Println("========================================")
		for k := 0; k < banyakMahasiswa; k++ {
			if arr2[k].Program == "D3" && arr2[k].Jurusan == "Informatika" {
				fmt.Printf("Nama		: %s\nNomor Induk	: %s\nProgram Studi	: D3 Informatika\n", arr[k].NamaMahasiswa, arr[k].NomorInduk)
				fmt.Println("========================================")
			}
		}

	case 6:
		fmt.Println("========================================")
		for k := 0; k < banyakMahasiswa; k++ {
			if arr2[k].Program == "D3" && arr2[k].Jurusan == "Sistem Informasi" {
				fmt.Printf("Nama		: %s\nNomor Induk	: %s\nProgram Studi	: D3 Sistem Informasi\n", arr[k].NamaMahasiswa, arr[k].NomorInduk)
				fmt.Println("========================================")
			}
		}

	}
}

func SortingLulus(arr TabArr, arr3 TabKonf, banyakMahasiswa int) {
	fmt.Println("========================================")
	for k := 0; k < banyakMahasiswa; k++ {
		if arr3[k].Konfirmasi == true {
			fmt.Printf("Nama	 : %s\nNo.Induk : %s\nStatus	 : Lulus\n\n", arr[k].NamaMahasiswa, arr[k].NomorInduk)
		}
	}
	fmt.Println("========================================")
}

func SortingDitolak(arr TabArr, arr3 TabKonf, banyakMahasiswa int) {
	fmt.Println("========================================")
	for k := 0; k < banyakMahasiswa; k++ {
		if arr3[k].Konfirmasi == false {
			fmt.Printf("Nama	 : %s\nNo.Induk : %s\nStatus	 : Tidak Lulus\n", arr[k].NamaMahasiswa, arr[k].NomorInduk)
			fmt.Println("========================================")
		}
	}
}

func SortingUmur(arr TabArr, banyakMahasiswa int) {
	var umur int
	fmt.Print("Masukan Umur yang akan dicari: ")
	fmt.Scan(&umur)
	fmt.Println()
	ki := 0
	ka := banyakMahasiswa
	idx := -1
	for ki <= ka && idx == -1 {
		te := (ka + ki) / 2
		if arr[te].UmurMahasiswa < umur {
			ka = te - 1
		} else if arr[te].UmurMahasiswa > umur {
			ki = te + 1
		} else if arr[te].UmurMahasiswa == umur {
			idx = te
		}
	}
	if idx == -1 {
		fmt.Printf("Tidak Terdapat Mahasiswa Dengan Umur %d\n", umur)
	} else {
		fmt.Println("========================================")
		for i := 0; i < banyakMahasiswa; i++ {
			if arr[i].UmurMahasiswa == umur {
				fmt.Printf("Nama	 : %s\nNo.Induk : %s\numur	 : %d\n", arr[i].NamaMahasiswa, arr[i].NomorInduk, arr[i].UmurMahasiswa)
				fmt.Println("========================================")
			}
		}
	}
}

func IsInput(arr *TabArr, arr2 *TabArr2, banyakMahasiswa *int) {
	/* input Nama, NIK, Tanggal lahir, Umur, Asal kota, Alamat dari Mahasiswa yang didaftarkan */
	var ProgramStud int
	i := *banyakMahasiswa
	fmt.Printf("Nama			      : ")
	fmt.Scan(&arr[i].NamaMahasiswa)

	fmt.Printf("Nomor Induk Kependudukan      : ")
	fmt.Scan(&arr[i].NomorInduk)

	fmt.Printf("Tanggal Lahir (Ex : 8 9 2004) : ")
	fmt.Scan(&arr[i].TanggalLahir, &arr[i].BulanLahir, &arr[i].TahunLahir)

	fmt.Print("Umur			      : ")
	fmt.Scan(&arr[i].UmurMahasiswa)

	fmt.Printf("Asal Kota 		      : ")
	fmt.Scan(&arr[i].KotaTinggal)

	fmt.Printf("Alamat Rumah		      : ")
	fmt.Scan(&arr[i].NamaJalan)

	fmt.Println()

	MenuJurusan()
	fmt.Print("Pilih Jurusan		: ")
	fmt.Scan(&ProgramStud)
	switch ProgramStud {
	case 1:
		arr2[i].Program = "S1"
		arr2[i].Jurusan = "Informatika"
	case 2:
		arr2[i].Program = "S1"
		arr2[i].Jurusan = "Sistem Informasi"
	case 3:
		arr2[i].Program = "D4"
		arr2[i].Jurusan = "Informatika"
	case 4:
		arr2[i].Program = "D4"
		arr2[i].Jurusan = "Sistem Informasi"
	case 5:
		arr2[i].Program = "D3"
		arr2[i].Jurusan = "Informatika"
	case 6:
		arr2[i].Program = "D3"
		arr2[i].Jurusan = "Sistem Informasi"
	}
	fmt.Print("Masukan nilai anda	: ")
	fmt.Scan(&arr2[i].NilaiTes)
	i += 1
	fmt.Println()
	*banyakMahasiswa = i
}

func LihatStatusKelulusan(arr TabArr, arr2 TabArr2, banyakMahasiswa int, arr3 TabKonf, CariNama string) {
	ada := false
	idx := -1
	for i := 0; i < banyakMahasiswa; i++ {
		if CariNama == arr[i].NamaMahasiswa {
			ada = true
			idx = i
		}
	}
	if ada == true {
		fmt.Printf(" Nama		  : %s		\n", arr[idx].NamaMahasiswa)
		fmt.Printf(" Nomor Induk	  : %s 			\n", arr[idx].NomorInduk)
		fmt.Printf(" Status Kelulusan : %t\n", arr3[idx].Konfirmasi)
	} else {
		fmt.Println("||	Nama tidak ditemukan.	||")
	}
}

func CekStatusMahasiswa(arr TabArr, banyakMahasiswa int, CariNama string) {
	// Binary Search //
	idx := 0
	for i := 0; i < banyakMahasiswa; i++ {
		if arr[i].NamaMahasiswa == CariNama {
			idx += 1
		}
	}
	if idx != 0 {
		fmt.Println("||	Nama Sudah Terdaftar	||")
	} else {
		fmt.Println("||	Nama Belum Terdaftar	||")
	}
}

func MenuAdminUser() {
	fmt.Println("========== Login Sebagai =========")
	fmt.Println("||	1. Admin 		||")
	fmt.Println("||	2. Mahasiswa		||")
	fmt.Println("||	3. Exit			||")
	fmt.Println("==================================")
} // Ditampilkan saat pertama kali program dijalankan

func MenuAdmin() {
	fmt.Println("==================================================")
	fmt.Println("Selamat Datang!")
	fmt.Println("==================================================")
	fmt.Println("||	1. Cek Daftar Mahasiswa			||")
	fmt.Println("||	2. Konfirmasi Kelulusan Mahasiswa	||")
	fmt.Println("||	3. Edit Daftar Mahasiswa		||")
	fmt.Println("||	4. Kembali				||")
	fmt.Println("==================================================")
} // Ditampilkan saat pengguna login sebagai admin

func MenuSortingAdmin() {
	fmt.Println("============= Urutkan berdasarkan ===========")
	fmt.Println("||	1. Nilai Tes Rendah - Tinggi	   ||")
	fmt.Println("||	2. Nilai Tes Tinggi - Rendah	   ||")
	fmt.Println("============= Cari berdasarkan ==============")
	fmt.Println("||	3. Jurusan			   ||")
	fmt.Println("||	4. Mahasiswa Diterima		   ||")
	fmt.Println("||	5. Mahasiswa Ditolak		   ||")
	fmt.Println("||	6. Cari Mahasiswa Dengan Umur	   ||")
	fmt.Println("||	7. Kembali  			   ||")
	fmt.Println("=============================================")
} // Ditampilkan saat pengguna memilih menu Cek Pendaftar di halaman admin

func MenuJurusan() {
	fmt.Println("======= Pilih Program Studi=======")
	fmt.Println("|| 1. S1 Informatika		||")
	fmt.Println("|| 2. S1 Sistem Informasi	||")
	fmt.Println("|| 3. D4 Informatika		||")
	fmt.Println("|| 4. D4 Sistem Informasi 	||")
	fmt.Println("|| 5. D3 Informatika		||")
	fmt.Println("|| 6. D3 Sistem Informasi	||")
	fmt.Println("==================================")
	fmt.Println()
} // Ditampilkan setelah mahasiswa mengisi informasi data diri

func MenuUser() {
	fmt.Println("===================================")
	fmt.Println("Selamat Datang,")
	fmt.Println("Portal Pendaftaran Calon Mahasiswa")
	fmt.Println("===================================")
	fmt.Println("||1. Daftar			||")
	fmt.Println("||2. Cek Status Pendaftaran	||")
	fmt.Println("||3. Cek Status Kelulusan	||")
	fmt.Println("||4. Kembali			||")
	fmt.Println("==================================")
} // Ditampilkan di halaman Mahasiswa setelah mengisi data diri

func Jurusan() {
	fmt.Println("======= Pilih Berdasarkan ========")
	fmt.Println("|| 1. S1 Informatika		||")
	fmt.Println("|| 2. S1 Sistem Informasi	||")
	fmt.Println("|| 3. D4 Informatika		||")
	fmt.Println("|| 4. D4 Sistem Informasi 	||")
	fmt.Println("|| 5. D3 Informatika		||")
	fmt.Println("|| 6. D3 Sistem Informasi	||")
	fmt.Println("==================================")
} // Ditampilkan di halaman admin Cek Pendaftar
