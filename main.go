package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

var roman = map[string]string {
	"C":    "100",
	"XC":   "90",
	"L":    "50",
	"XL":   "40",
	"X":    "10",
	"IX":   "9",
	"VIII": "8",
	"VII":  "7",
	"VI":   "6",
	"V":    "5",
	"IV":   "4",
	"III":  "3",
	"II":   "2",
	"I":    "1",
}

var arab = map[string]string {
	"100": "C",  
	"90": "XC",
	"50": "L",    
	"40": "XL", 
	"10": "X", 
	"9": "IX",
	"8": "VIII",
	"7": "VII", 
	"6": "VI",
	"5": "V",   
	"4": "IV", 
	"3": "III", 
	"2": "II", 
	"1": "I",   
}

func IsCorrect(s string, arr [4]string) (string, string, string) {
	var proverka, arrPerem, arrOper []string
	var oper string
	proverka = strings.Split(s, " ")

	if len(proverka) == 3 {
		for i := 0; i < len(arr); i++ {
			if strings.Contains(s, arr[i]) == true {
				arrOper = append(arrOper, arr[i])
				oper = arr[i]
			}
		}

		text := strings.ReplaceAll(s, " ", "")

		if len(arrOper) < 1 {
			panic("Нет ни одного оператора")
		} else if len(arrOper) > 1 {
			panic("Не может быть больше одного оператора")
		} else {
			arrPerem = strings.Split(text, oper)
			return oper, arrPerem[0], arrPerem[1]
		}

	} else {
		panic("Неверный формат")
	}

}

func isRomainFunc(a string, b string, roman map[string]string) (bool, int, int) {
	var romain, arabic []string
	var two string
	for k, v := range roman {
		if len(romain) == 2 || len(arabic) == 2 {
			break
		} else if k == a && k == b {
			romain = append(romain, a, a)
		} else if v == a && v == b {
			arabic = append(arabic, a, a)
		} else if k == a {
			romain = append(romain, a)
			if two != "" {
				romain = append(romain, two)
			}
		} else if k == b {
			two = k
			if len(romain) == 1 {
				romain = append(romain, b)
			}
		} else if v == a {
			arabic = append(arabic, a)
			if two != "" {
				arabic = append(arabic, two)
			}
		} else if v == b {
			two = v 
			if len(arabic) == 1 {
				arabic = append(arabic, b)
			}
		} 
	}

	if len(arabic) == 2 {
		one, _ := strconv.Atoi(arabic[0])
		two, _ := strconv.Atoi(arabic[1])
		return false, one, two
	} else if len(romain) == 2 {
		one, _ := strconv.Atoi(roman[romain[0]])
		two, _ := strconv.Atoi(roman[romain[1]])
		return true, one, two
	} else {
		panic("Используются разные системы счисления")
	}
}

func cal(a int, b int, oper string, isRomain bool, arab map[string]string) {
	var otvet int
	if a > 0 && b > 0 && a <= 10 && b <= 10 {
		switch oper{
		case "+":
			otvet = a + b
		case "-":
			otvet = a - b
		case "*":
			otvet = a * b
		case "/":
			otvet = a / b	
		}
	} else {
		panic("Калькулятор работает только с числами от 1 до 10 включительно")
	}
	
	if isRomain == false {
		fmt.Println(otvet)
	} else {
		if otvet == 0 {
			panic("В римской системе нет 0")
		} else if otvet >= 1 {
			if otvet >= 1 && otvet <= 9 || otvet == 10 || otvet == 40 || otvet == 50 || otvet == 90 || otvet == 100 {
				fmt.Print(arab[strconv.Itoa(otvet)])
			} else if otvet > 11 && otvet < 40 {
				fmt.Println(strings.Repeat("X", otvet / 10) +  arab[strconv.Itoa(otvet % 10)])
			} else if otvet > 40 && otvet < 50 {
				fmt.Println("XL" +  arab[strconv.Itoa(otvet % 10)])
			} else if otvet > 50 && otvet < 90 {
				fmt.Println("L" + strings.Repeat("X", otvet / 10 - 5) +  arab[strconv.Itoa(otvet % 10)])
			} else if otvet > 90 && otvet < 99 {
				fmt.Println("XC" +  arab[strconv.Itoa(otvet % 10)])
			}
		} else {
			panic("В римской в римской системе нет отрицательных чисел")
		}
	}
	
}

func main() {
    reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите значения")
	operator := [4]string{"+", "-", "/", "*"} 
	text, _ := reader.ReadString('\n')
	text = strings.ToUpper(strings.TrimSpace(text))
	oper, a, b := IsCorrect(text, operator)
	isRomain, aInt, bInt := isRomainFunc(a, b, roman)
	cal(aInt, bInt, oper, isRomain, arab)
}