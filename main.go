package main

import (
	"fmt"
	"strconv"
)

func reverse(s string) (result string) {
	for _,v := range s {
	  result = string(v) + result
	}
	return 
}

func toInt (s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err.Error())
	}
	return i
}

func isValidId (id string) bool {
	if len(id) != 13 {
		return false
	}
	reverseId := reverse(id)
	sum := 0
	fmt.Println(reverseId)

	for i, v := range reverseId {
		if i == 0 {
			continue
		}
		sum += (i + 1) * toInt(string(v))

	}
	moddulo := sum % 11
	validateChar := strconv.Itoa( (11 - moddulo ) % 10)[0]
	return validateChar == reverseId[0]
}

func main() {
	input := "1234567890121"
	fmt.Println(isValidId(input))

}