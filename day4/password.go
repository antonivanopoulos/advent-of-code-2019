package main

func PossiblePassword(input string) bool {
	var repeatingDigit int = -1
	repeating := make(map[byte]int)

	for index := 1; index < len(input); index++ {
		currentChar := input[index]
		currentNum := int(currentChar - '0')

		prevChar := input[index-1]
		prevNum := int(prevChar - '0')

		if currentNum < prevNum {
			return false
		}

		if currentNum == prevNum {
			if currentNum == repeatingDigit {
				repeating[currentChar]++
			} else {
				repeating[currentChar] = 2
			}

			repeatingDigit = currentNum
		} else {
			repeatingDigit = -1
		}
	}

	for _, value := range repeating {
		if value == 2 {
			return true
		}
	}

	return false
}
