package main

import "fmt"

func main() {
var (
nama, kelas, nim string
)
fmt.Print("Masukkan Nama Anda:")
fmt.Scan(&nama)
fmt.Println("Masukkan Nomor Kelas Anda:")
fmt.Scan(&kelas)
fmt.Print("Masukkan NIM Anda:")
fmt.Scan(&nim)
fmt.Println("Haloo telyutizen kenalin nama ku " + nama + " aku adalah salah satu mahasiswa informatika & aku dari kelas IF- " + kelas + " dengan Nomor Induk Mahasiswa: " + nim)

}
