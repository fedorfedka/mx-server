package main

import (
	"fmt"
	"net/smtp"
)

func main() {
	client, err := smtp.Dial("mx.yandex.ru:25")
	if err == nil {
		fmt.Print(client)
	} else {
		fmt.Print(err)
	}
}
