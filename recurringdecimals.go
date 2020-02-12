package recurringdecimals

import (
	"fmt"
	"math"
	"sort"
)

func longDivision(memo map[int]string, currNum int, currDenom int, currStr string) string {
	if currNum == 0 {
		return currStr
	}

	currNum *= 10

	multiplier := sort.Search(10, func(i int) bool {
		return i*currDenom >= currNum
	})

	if multiplier*currDenom != currNum {
		multiplier--
	}

	nCurrDenom := currDenom - multiplier*currNum

	val, saw := memo[nCurrDenom]

	if saw {
		withoutLast := currStr[:len(currStr)-1]

		if withoutLast[len(withoutLast)-len(val):] == val {
			return fmt.Sprintf("%s.(%s)", withoutLast[:len(withoutLast)-len(val)], val)
		}
	}

	currStr = fmt.Sprintf("%s%d", currStr, multiplier)
	memo[nCurrDenom] = currStr

	return longDivision(memo, currNum, currDenom, currStr)
}

// FractionToDecimal will return string representation of decimal number
func FractionToDecimal(numerator, denominator int) (ret string) {
	if int(math.Mod(float64(numerator), float64(denominator))) == 0 {
		ret = fmt.Sprintf("%d", numerator/denominator)
		return
	}

	ret = fmt.Sprintf("0.%s", longDivision(make(map[int]string), numerator, denominator, ""))

	return
}
