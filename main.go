package main

import (
	"fmt"
	"math"
	"strings"
)

func doOperation(operator string, firstArgument int, secondArgument int) (int, string) {
	// converting string representation of operator and execute its meaning
	switch operator {
	case "*":
		return firstArgument * secondArgument, ""
	case "/":
		return firstArgument / secondArgument, ""
	case "+":
		return firstArgument + secondArgument, ""
	case "-":
		return firstArgument - secondArgument, ""
	default:
		return 0, "Invalid operator"
	}
}

func retrieveAsciiFromMayan(sLen int, L int, H int, mayanToAscii map[string]int) int {
	mayanNumber := ""
	// concatenate lines from inputs as mayan number representation
	for i := 0; i < sLen; i++ {
		var num1Line string
		_, err := fmt.Scan(&num1Line) // bind
		if err != nil {
			fmt.Println(err.Error()) // printing error will end in incorrect result and process end
		}
		mayanNumber += num1Line
	}
	//// debug log
	//fmt.Fprintf(os.Stderr, "mayanNumber: %v", mayanNumber)
	//fmt.Fprintln(os.Stderr, "")

	var mayanNumberSplit []string
	for i := 0; i < len(mayanNumber)/(L*H); i++ {
		index := i + 1
		mayanNumberSplit = append([]string{mayanNumber[L*H*(index-1) : (L * H * index)]}, mayanNumberSplit...)
	}
	//// debug log
	//fmt.Fprintf(os.Stderr, "mayanNumberSplit: %v", mayanNumberSplit)
	//fmt.Fprintln(os.Stderr, "")

	asciiNumber := 0
	for i, mayanKey := range mayanNumberSplit {
		asciiNumber += mayanToAscii[mayanKey] * int(math.Pow(20, float64(i)))
	}
	//// debug log
	//fmt.Fprintf(os.Stderr, "asciiNumber: %v", asciiNumber)
	//fmt.Fprintln(os.Stderr, "")

	return asciiNumber
}

func main() {
	var L, H int
	_, errLH := fmt.Scan(&L, &H) // bind
	if errLH != nil {
		fmt.Println(errLH.Error()) // printing error will end in incorrect result and process end
	}
	//// debug log
	//fmt.Fprintln(os.Stderr, L)
	//fmt.Fprintln(os.Stderr, H)

	var asciiToMayan = make(map[int]string)
	for i := 0; i < 20; i++ {
		asciiToMayan[i] = ""
	}

	for i := 0; i < H; i++ {
		var numeral string
		_, errNumeral := fmt.Scan(&numeral) // bind
		if errNumeral != nil {
			fmt.Println(errNumeral.Error()) // printing error will end in incorrect result and process end
		}
		mayanNumPartLen := 0
		mayanNum := ""
		asciiNum := 0
		for _, el := range strings.Split(numeral, "") {
			mayanNum += el
			mayanNumPartLen += 1
			if mayanNumPartLen >= L {
				asciiToMayan[asciiNum] += mayanNum
				asciiNum += 1
				mayanNum = ""
				mayanNumPartLen = 0
			}
		}
	}
	var mayanToAscii = make(map[string]int)
	for asciiKey, mayanKey := range asciiToMayan {
		mayanToAscii[mayanKey] = asciiKey
	}

	var S1 int
	_, errS1 := fmt.Scan(&S1) // bind
	if errS1 != nil {
		fmt.Println(errS1.Error()) // printing error will end in incorrect result and process end
	}
	firstAsciiNumber := retrieveAsciiFromMayan(S1, L, H, mayanToAscii)

	var S2 int
	_, errS2 := fmt.Scan(&S2)
	if errS2 != nil {
		fmt.Println(errS2.Error()) // printing error will end in incorrect result and process end
	} // bind
	secondAsciiNumber := retrieveAsciiFromMayan(S2, L, H, mayanToAscii)

	var operator string
	_, errOperator := fmt.Scan(&operator)
	if errOperator != nil {
		fmt.Println(errOperator.Error()) // printing error will end in incorrect result and process end
	} // bind
	//// debug log
	//fmt.Fprintf(os.Stderr, "operator: %v", operator)
	//fmt.Fprintln(os.Stderr, "")

	result, err := doOperation(operator, firstAsciiNumber, secondAsciiNumber)
	if err != "" {
		fmt.Println(err)
	} else {
		//// debug log
		//fmt.Fprintf(os.Stderr, "result ascii: %v", result)
		//fmt.Fprintln(os.Stderr, "")

		var mayanResult []string
		if result == 0 {
			mayanResult = []string{asciiToMayan[0]}
		}
		for result > 0 {
			remainder := result % 20
			mayanResult = append([]string{asciiToMayan[remainder]}, mayanResult...)
			result /= 20
		}

		for _, mayanResultDigit := range mayanResult {
			for i := 0; i < len(mayanResultDigit)/(L); i++ {
				index := i + 1
				//// debug log
				//fmt.Fprintln(os.Stderr, mayanResultDigit[(index - 1):(L*index)])
				fmt.Println(mayanResultDigit[(index-1)*L : (L * index)])
			}
		}
	}
}
