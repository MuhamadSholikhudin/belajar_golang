package main

// HEllO word
import "fmt"

func main() {
    fmt.Println("Hello world")

    // komentar kode
    // menampilkan pesan hello world

    /*
    komentar kode
    menampilkan pesan hello world
    */

    // Variabel 
    /*
    var <nama-variabel> <tipe-data>
    var <nama-variabel> <tipe-data> = <nilai>
    */
    var firstName string = "john"

    // tanpa var, tanpa tipe data, menggunakan perantara ":="
    lastName := "wick"  
    var lastName string
    lastName = "wick"

    // "halo %s %s!\n" => %s di ganti nilai variabel
    fmt.Printf("halo %s %s!\n", firstName, lastName)

    // Deklarasi Multi Variabel
    
    var first, second, third string
    first, second, third = "satu", "dua", "tiga"

    
    var fourth, fifth, sixth string = "empat", "lima", "enam"
    
     seventh, eight, ninth := "tujuh", "delapan", "sembilan"

     one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

    
    _ = "belajar Golang"
    _ = "Golang itu mudah"
    name, _ := "john", "wick"

    // Deklarasi Variabel Menggunakan Keyword new
    name := new(string)

    fmt.Println(name)   // 0x20818a220
    fmt.Println(*name)  // ""


    // TIPE DATA
    var positiveNumber uint8 = 89
    var negativeNumber = -1243423644
    // Template %d pada fmt.Printf() digunakan untuk memformat data numerik non-desimal.
    fmt.Printf("bilangan positif: %d\n", positiveNumber)
    fmt.Printf("bilangan negatif: %d\n", negativeNumber)

    // Tipe Data Numerik Desimal
    var decimalNumber = 2.62
    // Template %f digunakan untuk memformat data numerik desimal menjadi string
    fmt.Printf("bilangan desimal: %f\n", decimalNumber)
    fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

    // Tipe Data bool (Boolean)
    var exist bool = true
    // Gunakan %t untuk memformat data bool menggunakan fungsi fmt.Printf().
    fmt.Printf("exist? %t \n", exist)

    // Tipe Data string
    // var message string = "Halo"
    
    var message = `Nama saya "John Wick".
    Salam kenal.
    Mari belajar "Golang".`

    fmt.Println(message)
    fmt.Printf("message: %s \n", message)

}