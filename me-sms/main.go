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
		{Content: "MIR-2778 10:02 Перевод 350р от Ирина С. Баланс: 1150.28р"},
		{Content: "Перевод из Тинькофф Банк +7968р от АЛЕКСАНДР М. MIR-2778 — Баланс: 9933.28р"},
		{Content: "Перевод из Газпромбанк +2778р от КИРИЛЛ П. MIR-2778 — Баланс: 18044.43р"},
		{Content: "MIR-2778 19:25 Перевод 10 790р от Людмила Д. Баланс: 15 266.43р"},
		{Content: "MIR-2778 13:35 перевод 400р Комиссия 2р TINKOFF Баланс: 18 062.43р"},
		{Content: "MIR-2778 13:38 Покупка 75.20р Urent Баланс: 17 987.23р"},
		{Content: "MIR-2778 17:26 Покупка 754р APTEKA Баланс: 18 396.23р"},
		{Content: "MIR-9550 13:18 Перевод 2650р от Элеонора К. Баланс: 19 579.59р"},
	}

	for _, v := range sms {
		NewSms(v)
	}

}

func NewSms(sms Message) {
	log.Println("new sms -", sms.Content)
	if !strings.Contains(strings.ToLower(sms.Content), "перевод") {
		log.Println("отсутствует информация о переводе")
		return
	}
	if !strings.Contains(strings.ToUpper(sms.Content), "MIR-2778") {
		log.Println("отсутствует информация о нужном счете")
		return
	}

	r := regexp.MustCompile(`(?s)\b(\d[\s\d]*(?:\.\d+)?)р`)

	f := r.FindAllStringSubmatch(sms.Content, -1)

	if f == nil || len(f[0]) == 0 {
		log.Println("отсутствует сумма")
		return
	}

	f[0][1] = strings.ReplaceAll(f[0][1], " ", "")
	amount, err := strconv.ParseFloat(f[0][1], 64)

	if err == nil {
		log.Println("идентификация суммы", amount)
	} else {
		log.Println("сумма не идентифицирована")
		return
	}
}

//СЧЁТ7999 10:02 Перевод 350р от Ирина С. Баланс: 1150.28р
//Перевод из Тинькофф Банк +7968р от АЛЕКСАНДР М. СЧЁТ7999 — Баланс: 9933.28р
//Перевод из Газпромбанк +2778р от КИРИЛЛ П. СЧЁТ3550 — Баланс: 18044.43р
//СЧЁТ3550 19:25 Перевод 10 790р от Людмила Д. Баланс: 15 266.43р
//MIR-9550 13:35 перевод 400р Комиссия 2р TINKOFF Баланс: 18 062.43р
//MIR-9550 13:38 Покупка 75.20р Urent Баланс: 17 987.23р
