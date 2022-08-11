package main

import (
	"fmt"
	"strconv"
)

//* INTERFACE ================================
type LapFAce interface {
	InfoFullLaptop() string
}

func main() {
	//* AUSUS LAPTOP===========================
	choseLaptop := ""
	configirationOfLaptop := 0
	fmt.Print("\nEnter the name of the Laptop that you are looking for : ")
	fmt.Scan(&choseLaptop)
	if choseLaptop == "Asus" || choseLaptop == "asus" {
		Asus1 := LapTop{
			Brand:     "ASUS",
			Model:     "Tuf gaming A15",
			Color:     "Black",
			Price:     "$1250",
			RAM:       8,
			SSD:       512,
			Processor: "AMD Ryzen 7 5800H, Inner DDR : 3.7GB ",
		}

		//? WHAT TO GET EXACTLY =================
		fmt.Println("\n1.Full name of the Laptop")
		fmt.Println("2.Color")
		fmt.Println("3.Price")
		fmt.Println("4.Configuration")
		fmt.Println("5.Processor Info")
		fmt.Println("6.FUll info about the laptop")
		fmt.Print("\nWhat exactly do you want to know choose above and write it down : ")
		fmt.Scan(&configirationOfLaptop)
		switch configirationOfLaptop {
		case 1:
			fmt.Println(Asus1.GetFullNameOfLaptop())
		case 2:
			fmt.Println(Asus1.Color)
		case 3:
			fmt.Println(Asus1.Price)
		case 4:
			fmt.Println("Capacity of RAM and SSD -> ", Asus1.RAM, " ", Asus1.SSD)
		case 5:
			fmt.Print("Processor Info : ", Asus1.Processor)
		case 6:
			asus1 := new(Asus)
			asus1.Brand = "ASUS"
			asus1.Model = "Tuf gaming F15"
			asus1.Color = "Black"
			asus1.Price = "$1250"
			asus1.RAM = 8
			asus1.SSD = 512
			asus1.Processor = "AMD Ryzen 7 5800H, Inner DDR : 3.7GB "
			fmt.Print("\n", asus1.InfoFullLaptop(), "\n\n")

		}

	}

	//* MacBook ===============================
	if choseLaptop == "Macbook" || choseLaptop == "macbook" {
		mac1 := new(MacBook)
		mac1.Brand = "MacBook"
		mac1.Model = "Pro"
		mac1.Color = "Space Grey"
		mac1.Price = "$1400"
		mac1.RAM = 16
		mac1.SSD = 512
		mac1.Processor = "M1"
		fmt.Print("\n", mac1.InfoFullLaptop())
	}

	//* Lenovo ================================
	if choseLaptop == "Lenovo" || choseLaptop == "lenovo" {
		lenovo1 := new(Lenovo)
		lenovo1.Brand = "Lenovo"
		lenovo1.Model = "Legion"
		lenovo1.Color = "Dark Blue"
		lenovo1.Price = "$1000"
		lenovo1.RAM = 16
		lenovo1.SSD = 512
		lenovo1.Processor = "intel core i 7 11th gen H"
		fmt.Print("\n", lenovo1.InfoFullLaptop())
	}
}

//* STRUCTS AND FUNCTIONS =====================
type LapTop struct {
	Brand, Model, Color, Price string
	RAM, SSD                   int
	Processor                  string
}

type MacBook struct {
	LapTop
}

func (m MacBook) InfoFullLaptop() string {
	result := "Brand : " + m.Brand + "\nColor : " + m.Color + "\nModel : " + m.Model + "\nPrice : " +
		m.Price + "\nRAM : " + strconv.Itoa(m.RAM) + "\nSSD : " + strconv.Itoa(m.SSD) + "\nProcessor : " + m.Processor
	return result
}

//? Asus LAPTOP=============
type Asus struct {
	LapTop
}

func (a Asus) InfoFullLaptop() string {
	asusFullInfo := "Brand : " + a.Brand + "\nModel : " + a.Model + "\nColor : " + a.Color +
		"\nPrice : " + a.Price + "\n" + "RAM : " + strconv.Itoa(a.RAM) + "\n" +
		"SSD : " + strconv.Itoa(a.SSD) + "\n" + "Processor : " + a.Processor
	return asusFullInfo

}

type Lenovo struct {
	LapTop
}

func (l Lenovo) InfoFullLaptop() string {
	asusFullInfo := "Brand : " + l.Brand + "\nModel : " + l.Model + "\nColor : " + l.Color +
		"\nPrice : " + l.Price + "\n" + "RAM : " + strconv.Itoa(l.RAM) + "\n" +
		"SSD : " + strconv.Itoa(l.SSD) + "\n" + "Processor : " + l.Processor
	return asusFullInfo
}

func (NameOfLap LapTop) GetFullNameOfLaptop() string {
	return NameOfLap.Brand + " " + NameOfLap.Model
}
