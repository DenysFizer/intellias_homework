package main

import "fmt"

func main() {
	// Одне яблуко коштує 5.99 грн. Ціна однієї груші - 7 грн.
	// Ми маємо 23 грн.
	// 1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?
	// 2. Скільки груш ми можемо купити?
	// 3. Скільки яблук ми можемо купити?
	// 4. Чи ми можемо купити 2 груші та 2 яблука?

	applePrice := 5.99
	pineapplePrice := 7.00
	ourmoney := 23.00
	// 1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?
	fmt.Println("1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?")
	moneyForbuying := 9*applePrice + 8*pineapplePrice
	fmt.Println("Грошей потрібно", moneyForbuying)
	// 2. Скільки груш ми можемо купити?
	fmt.Println("2. Скільки груш ми можемо купити?")
	numberofPineaples := int(ourmoney / pineapplePrice)
	fmt.Println("Груш купимо", numberofPineaples)
	// 3. Скільки яблук ми можемо купити?
	fmt.Println("3. Скільки яблук ми можемо купити?")
	numberofApples := int(ourmoney / applePrice)
	fmt.Println("Яблук купимо", numberofApples)
	// 4. Чи ми можемо купити 2 груші та 2 яблука?
	fmt.Println("Чи ми можемо купити 2 груші та 2 яблука?")
	moneyNeed := 2*applePrice + 2*pineapplePrice
	if ourmoney > moneyNeed {
		fmt.Println("Так зможемо купити")
	} else {
		fmt.Println("Ні не зможемо купити")
	}

}
