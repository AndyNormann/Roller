package main

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage ./%s (<amount>)d<dice>(+<plus>)")
	}
	reg := regexp.MustCompile(`(?P<Amount>\d*)?d(?P<Dice>\d*)\+?(?P<Plus>\d*)?`)
	input := os.Args[1]
	var (
		dice    = 0
		amount  = 1
		plus    = 0
		matches = reg.FindStringSubmatch(input)
	)

	paramsMap := make(map[string]string)
	for i, name := range reg.SubexpNames() {
		paramsMap[name] = matches[i]
	}

	if paramsMap["Plus"] != "" {
		num, err := strconv.Atoi(paramsMap["Plus"])
		if err == nil {
			plus = num
		}
	}

	if paramsMap["Dice"] != "" {
		num, err := strconv.Atoi(paramsMap["Dice"])
		if err == nil {
			dice = num
		}
	}

	if paramsMap["Amount"] != "" {
		num, err := strconv.Atoi(paramsMap["Amount"])
		if err == nil {
			amount = num
		}
	}

	result := 0
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := 0; i < amount; i++ {
		current := r.Intn(dice) + 1
		result += current
		fmt.Printf("%d ", current)
	}
	if plus != 0 {
		result += plus
		fmt.Printf("+ %d", plus)
	}
	fmt.Printf(" = %d\n", result)
}
