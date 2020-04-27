package main

import "fmt"

func main()  {
	rupiah := 15550
	usd := 2

	fmt.Println("Selamat Datang di USD to IDR konverter")
	fmt.Printf("Kurs tukar untuk %d USD ke IDR saat ini adalah %d Rupiah", usd, rupiah*usd)
}