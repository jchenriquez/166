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

	if currNum < currDenom {
		currNum *= 10
	}

	fmt.Printf("currNum %d\n", currNum)
	var nCurrNum int

	multiplier := sort.Search(10, func(i int) bool {
		return i*currDenom >= currNum
	})

	if multiplier*currDenom > currNum {
		multiplier--
	}

	if multiplier == 0 {
		nCurrNum = currNum * 10
	} else {
		nCurrNum = currNum - multiplier*currDenom
	}

	fmt.Printf("nCurrNum %d\n", nCurrNum)
	currStr = fmt.Sprintf("%s%d", currStr, multiplier)
	val, saw := memo[nCurrNum]

	if saw {
		reptedVal := fmt.Sprintf("%s%s", val, val)

		if len(currStr) >= len(reptedVal) && reptedVal == currStr[len(currStr)-len(reptedVal):] {
			return fmt.Sprintf("%s(%s)", currStr[:len(currStr)-len(reptedVal)], val)
		}
	}

	if !saw {
		memo[nCurrNum] = currStr
	} else {
		memo[nCurrNum] = fmt.Sprintf("%d", multiplier)
	}

	return longDivision(memo, nCurrNum, currDenom, currStr)
}

// FractionToDecimal will return string representation of decimal number
func FractionToDecimal(numerator, denominator int) (ret string) {
	if int(math.Mod(float64(numerator), float64(denominator))) == 0 {
		ret = fmt.Sprintf("%d", numerator/denominator)
		return
	}

	ret = fmt.Sprintf("0.%s", longDivision(make(map[int]string), numerator, denominator, ""))
	fmt.Printf("ret %s\n", ret)

	return
}
