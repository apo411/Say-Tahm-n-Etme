package welcome

import "fmt"

func Welcome() {
	ismin := " "
	fmt.Println("Sana hitap edebilmem için ismini yazarmısın?")
	fmt.Scanln(&ismin)

	fmt.Printf("Selam %v en başta senden bir sayı söylemini isteyeceğim ve bundan 2saniye sonra konsol tamamen silinecek.", ismin)
	fmt.Println(" ")
	fmt.Println("Sende o sırada arkadaşına yazdığın sayıyı tahmin ettireceksin bu şekilde senin ile bir oyun oynamıs olacağız.")
}
