package main

import (
	"fmt"
)

func main() {

	//* USER METHOD 1 ======================================================
	var LookForSomone string
	var chooseWhatToGet int
	fmt.Println("\nWho do you want to find ? ")
	fmt.Print("Enter his first name : ")
	fmt.Scan(&LookForSomone)

	//? WHO LOOK FOR ======================================================
	if LookForSomone == "Boris" || LookForSomone == "boris" {
		user1 := User{
			Name:        "Asliddin",
			Surname:     "Mahamadjaniov",
			whereLives:  "Aviasozlar 2, 1/39",
			Birth:       "22.11.2002",
			phoneNumber: +998909771763,
			Ownings:     &Properties{item1: "Port", item2: "Car Business", item3: ""},
		}
		//? WHAT TO GET EXACTLY============================================
		fmt.Println("\n1. Full Name ")
		fmt.Println("2. Address ")
		fmt.Println("3. When was born ")
		fmt.Println("4. Phone Number ")
		fmt.Println("5. Properties ")
		fmt.Println("6. Full Info ")
		fmt.Print("\nWhat to get eaxctly ? -> ")
		fmt.Scan(&chooseWhatToGet)
		switch chooseWhatToGet {
		case 1:
			fmt.Println(user1.Name, user1.Surname)
		case 2:
			fmt.Println(user1.GetLocation())
		case 3:
			fmt.Println(user1.Birth)
		case 4:
			fmt.Println(user1.phoneNumber)
		case 5:
			fmt.Println(user1.Ownings.GetProperies())
		case 6:
			fmt.Println(user1)
		}
	}
	//* USER 2 METHOD GET FULL INFO ========================================
	if LookForSomone == "Diyorbek" {
		user2 := User{
			Name:        "Diyorbek",
			Surname:     "Anvarov",
			whereLives:  "Muruvvat ST, House : 80",
			Birth:       "03.02.2004",
			phoneNumber: +998950750305,
			Ownings:     &Properties{item1: "NO INFO"},
		}
		//? WHAT TO GET EXACTLY 2 ============================================
		fmt.Println("\n1. Full Name ")
		fmt.Println("2. Address ")
		fmt.Println("3. When was born ")
		fmt.Println("4. Phone Number ")
		fmt.Println("5. Properties ")
		fmt.Println("6. Full Info ")
		fmt.Print("\nWhat to get eaxctly ? -> ")
		fmt.Scan(&chooseWhatToGet)
		switch chooseWhatToGet {
		case 1:
			fmt.Println(user2.Name, user2.Surname)
		case 2:
			fmt.Println(user2.GetLocation())
		case 3:
			fmt.Println(user2.Birth)
		case 4:
			fmt.Println(user2.phoneNumber)
		case 5:
			fmt.Println(user2.Ownings.GetProperies())
		case 6:
			fmt.Println(user2)
		}
	}
}

func (UserLocation User) GetLocation() string {
	return UserLocation.whereLives
}

func (p Properties) GetProperies() string {
	return p.item1 + " " + p.item2 + " " + p.item3
}

type User struct {
	Name, Surname, whereLives, Birth string
	phoneNumber                      int
	Ownings                          *Properties
}

type Properties struct {
	item1, item2, item3 string
}
