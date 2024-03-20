package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Message struct {
	Content string
}

func main() {
	sms := []Message{
		{Content: "СЧЁТ7999 10:02 Перевод 350р от Ирина С. Баланс: 1150.28р"},
		{Content: "Перевод из Тинькофф Банк +7968р от АЛЕКСАНДР М. СЧЁТ7999 — Баланс: 9933.28р"},
		{Content: "Перевод из Газпромбанк +2778р от КИРИЛЛ П. СЧЁТ3550 — Баланс: 18044.43р"},
		{Content: "СЧЁТ3550 19:25 Перевод 10 790р от Людмила Д. Баланс: 15 266.43р"},
	}

	for _, v := range sms {
		NewSms(v)
	}

}

func NewSms(sms Message) {
	log.Println("new sms -", sms.Content)
	r, _ := regexp.Compile(`[+ ]((?:\d+ )*\d+)р`)
	f := r.FindAllStringSubmatch(sms.Content, -1)
	f[0][1] = strings.ReplaceAll(f[0][1], " ", "")

	if len(f) != 1 {
		log.Println("В теле сообщения отсутствует сумма")
		return
	}

	amount, err := strconv.ParseFloat(f[0][1], 64)

	if err == nil {
		log.Println("Идентификация суммы", amount)
	} else {
		log.Println("Сумма не идентифицирована")
		return
	}
}
