package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	calcInput         string
	arabicNumbers     = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	romeNumbers       = []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	operations        = []string{"*", "/", "+", "-"}
	operation         string
	idxOperation      int
	firstNumber       string
	typeFirstNumber   int
	secondNumber      string
	typeSecondNumber  int
	calcResult        string
	err               error
	numberConverstion = map[int]string{1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X"}
)

func main() {
	fmt.Println("Добро пожаловать в калькулятор арабских и римских чисел")
	calcInput, _ = bufio.NewReader(os.Stdin).ReadString('\n')

	operation, idxOperation, err = findOperation(calcInput)
	if err != nil {
		fmt.Println(err)
	}

	if len(calcInput) < 3 || idxOperation == 0 || idxOperation+1 == len(calcInput) {
		fmt.Println(errors.New("нет всех необходимых значений для расчета"))
	} else {
		firstNumber, typeFirstNumber, err = findNumber(strings.TrimSpace(calcInput[:idxOperation]))
		if err != nil {
			fmt.Println(err)
		}

		secondNumber, typeSecondNumber, err = findNumber(strings.TrimSpace(calcInput[idxOperation+1:]))
		if err != nil {
			fmt.Println(err)
		}

		if typeFirstNumber == 0 || typeFirstNumber != typeSecondNumber {
			fmt.Println(errors.New("нет всех необходимых чисел для расчета или у них разный тип"))
		} else {
			calcResult, err = calculation(firstNumber, secondNumber, operation, typeFirstNumber)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Результат сложения: ", calcResult)
			}
		}
	}
}

func findOperation(input string) (result string, idxOp int, err error) {
	var countAllOperations int
	for _, el := range operations {
		countOperation := strings.Count(input, el)
		countAllOperations = countAllOperations + countOperation
	}

	switch {
	case countAllOperations == 0:
		return "", 0, errors.New("не обнаружено ни одного знака операции")
	case countAllOperations > 1:
		return "", 0, errors.New("количество операций больше одной")
	}

	for _, el := range operations {
		idx := strings.Index(input, el)
		if idx > 0 {
			result = input[idx : idx+1]
			idxOp = idx
			err = nil
		}
	}
	return
}

func findNumber(input string) (result string, typeNumber int, err error) {
	for _, el := range arabicNumbers {
		if el == input {
			return el, 1, nil
		}
	}

	for _, el := range romeNumbers {
		if el == input {
			return el, 2, nil
		}
	}
	return "", 0, errors.New("не обнаружено ни одного числа от 1 (I) до 10 (X)")
}

func calculation(firstNumber string, secondNumber string, operation string, typeNumber int) (result string, err error) {
	var firstNumberInt int
	var secondNumberInt int
	var resultInt int

	switch typeNumber {
	case 1:
		firstNumberInt, err = strconv.Atoi(firstNumber)
		if err != nil {
			return "", errors.New("ошибка с первым числом")
		}

		secondNumberInt, err = strconv.Atoi(secondNumber)
		if err != nil {
			return "", errors.New("ошибка со вторым числом")
		}
	case 2:
		for key, value := range numberConverstion {
			if value == firstNumber && value == secondNumber {
				firstNumberInt = key
				secondNumberInt = key
			} else if value == firstNumber {
				firstNumberInt = key
			} else if value == secondNumber {
				secondNumberInt = key
			}
		}
	}

	switch operation {
	case "*":
		resultInt = firstNumberInt * secondNumberInt
	case "/":
		resultInt = firstNumberInt / secondNumberInt
	case "+":
		resultInt = firstNumberInt + secondNumberInt
	case "-":
		resultInt = firstNumberInt - secondNumberInt
	}

	if resultInt < 1 && typeNumber == 2 {
		return "", errors.New("результат римских числа не может быть меньше 1")
	} else if resultInt > 10 && typeNumber == 2 {
		return "", errors.New("результат римских числа не может быть больше 10")
	} else if typeNumber == 1 {
		return strconv.Itoa(resultInt), nil
	} else {
		for key, value := range numberConverstion {
			if key == resultInt {
				return value, nil
			}
		}
	}
	return "", errors.New("неизвестная ошибка")
}
