package main

import (
	"fmt"
	"net/http"
	"os"
)

//const (
//	baseURL = "https://iss.moex.com/iss/engines/stock/markets/shares/boards/TQBR/securities/%s/candles.json?from=%s&till=%s&interval=%d&start=%d"
//)

func main() {
	//security, dttmFrom, dttmTill, interval, startNum := "SBER", "2025-02-06 10:00:00", "2025-02-06 23:59:59", 1, 0
	//url := fmt.Sprintf(baseURL, security, dttmFrom, dttmTill, interval, startNum)

	url := "https://iss.moex.com/iss/engines/stock/markets/shares/boards/TQBR/securities/ABIO/candles.json?from=2025-02-06&till=2025-02-06&interval=1&start=0"
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	// Создание нового запроса
	//req, err := http.NewRequest("GET", url, nil)
	//if err != nil {
	//	fmt.Printf("Ошибка при создании запроса: %s\n", err)
	//	return
	//}

	// Добавление заголовков
	//req.Header.Add("Authorization", "Bearer YOUR_ACCESS_TOKEN")
	//req.Header.Add("Accept", "application/json")

	// Выполнение запроса
	//client := &http.Client{}
	//response, err := client.Do(req)
	//if err != nil {
	//	fmt.Printf("Ошибка при выполнении запроса: %s\n", err)
	//	return
	//}
	//defer response.Body.Close()

	// Чтение тела ответа
	//body, err := ioutil.ReadAll(response.Body)
	//if err != nil {
	//	fmt.Printf("Ошибка при чтении тела ответа: %s\n", err)
	//	return
	//}

	// Вывод тела ответа
	//fmt.Println(string(body))
}
