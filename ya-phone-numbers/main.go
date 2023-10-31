package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var phoneNumMain, phoneNumAdd1, phoneNumAdd2, phoneNumAdd3 string
	fmt.Scan(&phoneNumMain, &phoneNumAdd1, &phoneNumAdd2, &phoneNumAdd3) // standard input

	comparePhone(phoneNumMain, phoneNumAdd1, phoneNumAdd2, phoneNumAdd3)
}

func comparePhone(phoneNumMain string, phonesNumAdd ...string) {
	for _, phoneNumAdd := range phonesNumAdd {
		if convertToValidPhone(phoneNumMain) == convertToValidPhone(phoneNumAdd) {
			fmt.Println("YES") // standard output
		} else {
			fmt.Println("NO") // standard output
		}
	}
}

func convertToValidPhone(phoneNum string) (validPhone int) {
	re := regexp.MustCompile("[0-9]")
	phoneNum = strings.Join(re.FindAllString(phoneNum, -1), "")
	if len(phoneNum) == 7 {
		phoneNum = "7495" + phoneNum
	}
	if phoneNum[0:1] == "8" {
		phoneNum = "7" + phoneNum[1:]
	}
	validPhone, _ = strconv.Atoi(phoneNum)
	return
}
