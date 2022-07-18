package oyun

import "fmt"

var TahminEdileceksayı int = 0
var tahminEdilenSayı int = 0
var kackerede int = 0

func Belirleme() {
	fmt.Println("Tahmin edilecek sayıyı belirtin.")
	fmt.Scanln(&TahminEdileceksayı)
	fmt.Println("Belirtilen sayı kaydedildi konsol 5 saniye içinde silinecek!!")
}

func Soru() {
	fmt.Println("Merhabalar Senden bir sayı söylemini isteyeceğim")
	fmt.Scanln(&tahminEdilenSayı)
	kackerede = kackerede + 1
}

func Dongu() {
	for TahminEdileceksayı != tahminEdilenSayı {
		if tahminEdilenSayı > TahminEdileceksayı {
			fmt.Println("Daha küçük bir sayı belirtin.")
			fmt.Scanln(&tahminEdilenSayı)
			kackerede = kackerede + 1

		}
		if tahminEdilenSayı < TahminEdileceksayı {
			fmt.Println("Daha büyük bir sayı belirtin.")
			fmt.Scanln(&tahminEdilenSayı)
			kackerede = kackerede + 1

		}

	}

}

func Bitis() {
	if tahminEdilenSayı == TahminEdileceksayı {
		fmt.Println("Basarılı bir şekilde sonuca ulaştınız.")
		fmt.Printf("Sonuca toplamda %v bildin.", kackerede)
		fmt.Println(" ")
	}

}
