package main

import (
	"fmt"
)

func menu() {
	flag := true
	var choice string
	var ticker string

	for flag {
		fmt.Println("\nДобро пожаловать в приложение Терминал! \nС помощью данного простого приложения Вы можете " +
			"узнать рыночную стоимость акции, торгуемой на Московской бирже. \nДля этого Вам всего лишь надо " +
			"ввести её тикер. \nНапример, введя MTSS вы узнаете сколько стоит один лот акций МТС. \nЧтобы узнать полный список тикеров нажмите 2.")
		fmt.Println("Введите 1 чтобы узнать стоимость одного лота акции \nВведите 2 чтобы ознакомиться с полным списком тикеров")
		fmt.Scan(&choice)

		switch {
		case choice == "1":
			fmt.Println("Введите тикер")
			fmt.Scan(&ticker)
			input(ticker)
		case choice == "2":
			fmt.Println("MTSS")
		default:
			fmt.Println("Неверный выбор!")
			flag = false
		}

	}
}
