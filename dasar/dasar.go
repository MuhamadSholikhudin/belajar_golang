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

    // Konstanta
    const firstNameconst string = "john"
    fmt.Print("halo ", firstNameconst, "!\n")

    // Operator
    var value1 = (((2 + 6) % 3) * 4 - 2) / 3
    fmt.Println(value1)
    

    // Operator Perbandingan
    var value2 = (((2 + 6) % 3) * 4 - 2) / 3
    var isEqual2 = (value2 == 2)

    fmt.Printf("nilai %d (%t) \n", value2, isEqual2)
    

    // Operator Logika
    var left = false
    var right = true

    var leftAndRight = left && right
    fmt.Printf("left && right \t(%t) \n", leftAndRight)

    var leftOrRight = left || right
    fmt.Printf("left || right \t(%t) \n", leftOrRight)

    var leftReverse = !left
    fmt.Printf("!left \t\t(%t) \n", leftReverse)

    // Seleksi Kondisi

    // Seleksi Kondisi Menggunakan Keyword if, else if, & else
    
    var point1 = 8

    if point1 == 10 {
        fmt.Println("lulus dengan nilai sempurna")
    } else if point1 > 5 {
        fmt.Println("lulus")
    } else if point1 == 4 {
        fmt.Println("hampir lulus")
    } else {
        fmt.Printf("tidak lulus. nilai anda %d\n", point1)
        
    }

    // Variabel Temporary Pada if - else
    
    var point2 = 8840.0

    if percent2 := point2 / 100; percent2 >= 100 {
        fmt.Printf("%.1f%s perfect!\n", percent2, "%")
    } else if percent2 >= 70 {
        fmt.Printf("%.1f%s good\n", percent2, "%")
    } else {
        fmt.Printf("%.1f%s not bad\n", percent2, "%")
    }


    // Seleksi Kondisi Menggunakan Keyword switch - case

    var point3 = 6

    switch point3 {
    case 8:
        fmt.Println("perfect")
    case 7:
        fmt.Println("awesome")
    default:
        fmt.Println("not bad")
    }

    // Pemanfaatan case Untuk Banyak Kondisi

    var point4 = 6

    switch point4 {
    case 8:
        fmt.Println("perfect")
    case 7, 6, 5, 4:
        fmt.Println("awesome")
    default:
        fmt.Println("not bad")
    }

    // Kurung Kurawal Pada Keyword case & default


    var point5 = 6
    
    switch point5 {
    case 8:
        fmt.Println("perfect")
    case 7, 6, 5, 4:
        fmt.Println("awesome")
    default:
        {
            fmt.Println("not bad")
            fmt.Println("you can be better!")
        }
    }


    // Switch Dengan Gaya if - else
    
    var point6 = 6

    switch {
    case point6 == 8:
        fmt.Println("perfect")
    case (point6 < 8) && (point6 > 3):
        fmt.Println("awesome")
    default:
        {
            fmt.Println("not bad")
            fmt.Println("you need to learn more")
        }
    }

    // Penggunaan Keyword fallthrough Dalam switch
    
    var point7 = 6

    switch {
    case point7 == 8:
        fmt.Println("perfect")
    case (point7 < 8) && (point7 > 3):
        fmt.Println("awesome")
        fallthrough
    case point7 < 5:
        fmt.Println("you need to learn more")
    default:
        {
            fmt.Println("not bad")
            fmt.Println("you need to learn more")
        }
    }

    // Seleksi Kondisi Bersarang
    var point8 = 10

    if point8 > 7 {
        switch point8 {
        case 10:
            fmt.Println("perfect!")
        default:
            fmt.Println("nice!")
        }
    } else {
        if point8 == 5 {
            fmt.Println("not bad")
        } else if point8 == 3 {
            fmt.Println("keep trying")
        } else {
            fmt.Println("you can do it")
            if point8 == 0 {
                fmt.Println("try harder!")
            }
        }
    }

    // Perulangan

    // Perulangan Menggunakan Keyword for
    for i1 := 0; i1 < 5; i1++ {
        fmt.Println("Angka", i1)
    }

    // Penggunaan Keyword for Dengan Argumen Hanya Kondisi
    var i2 = 0

    for i2 < 5 {
        fmt.Println("Angka", i2)
        i2++
    }

    // Penggunaan Keyword for Tanpa Argumen
    var i3 = 0

    for {
        fmt.Println("Angka", i3)

        i3++
        if i3 == 5 {
            break
        }
    }

    // Penggunaan Keyword break & continue
    for i4 := 1; i4 <= 10; i4++ {
        if i4 % 2 == 1 {
            continue
        }

        if i4 > 8 {
            break
        }

        fmt.Println("Angka", i4)
    }

    // Perulangan Bersarang 
    for i5 := 0; i5 < 5; i5++ {
        for j5 := i5; j5 < 5; j5++ {
            fmt.Print(j5, " ")
        }

        fmt.Println()
    }

    // Pemanfaatan Label Dalam Perulangan

    outerLoop:
    for i6 := 0; i6 < 5; i6++ {
        for j6 := 0; j6 < 5; j6++ {
            if i6 == 3 {
                break outerLoop
            }
            fmt.Print("matriks [", i6, "][", j6, "]", "\n")
        }
    }

    // Array
    var names1 [4]string
    names1[0] = "trafalgar"
    names1[1] = "d"
    names1[2] = "water"
    names1[3] = "law"

    fmt.Println(names1[0], names1[1], names1[2], names1[3])

    // Pengisian Elemen Array yang Melebihi Alokasi Awal
    // var names [4]string
    // names[0] = "trafalgar"
    // names[1] = "d"
    // names[2] = "water"
    // names[3] = "law"
    // names[4] = "ez" // baris kode ini menghasilkan error


    // Inisialisasi Nilai Awal Array
    var fruits = [4]string{"apple", "grape", "banana", "melon"}

    fmt.Println("Jumlah element \t\t", len(fruits))
    fmt.Println("Isi semua element \t", fruits)

}