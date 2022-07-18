package main

import (
	"proje/konsol"
	"proje/oyun"
	"proje/welcome"
)

func main() {
	welcome.Welcome()
	oyun.Belirleme()
	konsol.Sil()
	oyun.Soru()
	oyun.Dongu()
	oyun.Bitis()

}
