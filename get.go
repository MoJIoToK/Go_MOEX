package main

import (
	"bytes"
	"fmt"
	"net/http"
	"unicode"
)

func httpGet(ticker string) {
	url := "https://iss.moex.com/iss/engines/stock/markets/shares/boards/TQBR/securities/" + ticker + ".csv?iss.only=securities&securities.columns=PREVPRICE"
	resp, err := http.Get(url)
	//resp, err := http.Get("https://iss.moex.com/iss/engines/stock/markets/shares/boards/TQBR/securities/SBER.csv?iss.only=securities&securities.columns=PREVPRICE")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	str := ""
	for true {

		bs := make([]byte, 1014)
		n, err := resp.Body.Read(bs)
		rs := bytes.Runes(bs)

		for i := 0; i < len(rs); i++ {
			if unicode.IsDigit(rs[i]) || unicode.IsPunct(rs[i]) {
				str = str + string(rs[i])
			}
		}

		if n == 0 || err != nil {
			break
		}
	}
	fmt.Println("\n" + str)
}
