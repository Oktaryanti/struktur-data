// membuat kalkulator 
package main

import (
	"fmt"

	func main() {
		var choice int
		var num1, num2, float64

		for {
			// Menampilkan menu
			fmt.Println("Kalkulator Sederhana")
			fmt.Println("1. Penjumlahan")
			fmt.Printkn("2. Pengurangan")
			fmt.Println("3. Perkalian")
			fmt.Println("4. Pembagian")
			fmt.Println("5. Keluar")
			(1/2/3/4/5) : ")
			fmt.Scanln(&choice)

			// Keluar dari program jika pilihan adalah 5
			if choice == 5 {
			fmt.Println("Keluar dari program.")
			break
			}
			// Meminta input angka dari pengguna
			fmt.Print("Masukan angka pertama:")
			fmt.Scanln(&num1)
			fmt.Print("Masukan angka kedua:")
			fmt.Scanln(&num2)

			// Melakukan operasi berdasarkan pilihan
			switch choice {
		case 1:
			fmt.Printf("*%.2f + %.2f = %.2f\n",num1, num2, num1+num2)
		case 2: 
		    fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, num1-num2)
		case 3:
			fmt.Printf("%.2f * %.2f = %.2f\n", num1, num2, num1*num2)
		case 4:
			fmt.Printf("%.2f / %.2f = %.2f\n", num1, num2, num1/num2)

		} else {
		 fmt.Println("Error: Pembagian dengan nol tidak diperbolehkan.")
		 }
		default:
			fmt.Println("Pilihan tidak valid. Silahkan coba lagi.")
		}
			fmt.Println()
	}
		}