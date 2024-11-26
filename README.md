# struktur-data
                                  TugasOlinRyanti.go                               
package main

import "fmt"

// Interface Benda
type Benda interface {
        Identitas() string
}

// Interface Mamalia, yang merupakan bagian dari Benda
type Mamalia interface {
        Benda
        Bernafas() string
}

// Interface Manusia, yang merupakan bagian dari Mamalia
type Manusia interface {
        Mamalia
        Berbicara() string
}

// Struct Mahasiswa, yang mengimplementasikan semua interface di atas
type Mahasiswa struct {
        Nama string
}

func (m Mahasiswa) Identitas() string {
        return fmt.Sprintf("Saya adalah sebuah benda bernama %s.", m.Nama)
}

func (m Mahasiswa) Bernafas() string {
        return "Sebagai mamalia, saya bernafas menggunakan paru-paru."
}

func (m Mahasiswa) Berbicara() string {
        return fmt.Sprintf("Halo, nama saya %s. Saya bisa berbicara karena saya manusia.", m.Nama)
}

func main() {
        // Membuat instance Mahasiswa
        mahasiswa := Mahasiswa{Nama: "Olin"}

        // Mengakses metode dari berbagai interface
        fmt.Println(mahasiswa.Identitas())
        fmt.Println(mahasiswa.Bernafas())
        fmt.Println(mahasiswa.Berbicara())
}
