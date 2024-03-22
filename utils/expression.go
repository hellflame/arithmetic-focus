package utils

import (
	"errors"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateExpression(max int, hardness, multiplyRate float64) string {
	mid := max / 2
	numbers := []int{}
	for i := 0; i < 2; i++ {
		var num int
		if rand.Float64() < hardness {
			// generate bigger number
			num = randRange(mid, max)
		} else {
			num = randRange(1, mid)
		}
		numbers = append(numbers, num)
	}

	operatorCandidates := []string{"+", "-"}
	if rand.Float64() < multiplyRate {
		operatorCandidates = append(operatorCandidates, "x")
	}

	operator := operatorCandidates[rand.Intn(len(operatorCandidates))]

	if operator == "-" {
		sort.Slice(numbers, func(i, j int) bool { return numbers[j] < numbers[i] })
	}

	tmp := []string{}
	for _, v := range numbers {
		tmp = append(tmp, strconv.Itoa(v))
	}

	return strings.Join(tmp, " "+operator+" ")
}

func randRange(min, max int) int {
	return min + rand.Intn(max-min)
}

func ParseExpression(exp string) (int, error) {
	operation := map[string]func(int, int) int{
		"+": func(i1, i2 int) int { return i1 + i2 },
		"-": func(i1, i2 int) int { return i1 - i2 },
		"x": func(i1, i2 int) int { return i1 * i2 },
	}
	numbers := []int{}
	operator := ""
	for _, char := range strings.Split(exp, " ") {
		if _, found := operation[char]; found {
			operator = char
		} else {
			num, e := strconv.Atoi(char)
			if e != nil {
				return 0, e
			}
			numbers = append(numbers, num)
		}
	}
	if operator == "" || len(numbers) != 2 {
		return 0, errors.New("unknown operator")
	}
	return operation[operator](numbers[0], numbers[1]), nil
}
