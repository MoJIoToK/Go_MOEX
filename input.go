package main

import (
	"fmt"
)

func input(resp string) {
	fmt.Println("Вы ввели - %s", resp)
	httpGet(resp)
}
