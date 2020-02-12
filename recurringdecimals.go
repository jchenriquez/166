package recurringdecimals

import (
	"fmt"
	"math"
	"sort"
)

func longDivision(memo map[int][]string, currNum int, currDenom int, currStr string) string {
	if currNum == 0 {
		return currStr
	}

  if currNum < currDenom {
	  currNum *= 10
  }
  var nCurrNum int

	multiplier := sort.Search(10, func(i int) bool {
		return i*currDenom >= currNum
	})

	if multiplier*currDenom > currNum {
		multiplier--
	}

  if multiplier == 0 {
    nCurrNum = currNum*10
  } else {
	  nCurrNum = currNum - multiplier*currDenom
  }

  fmt.Printf("nCurrNum %d\n", nCurrNum)
	vals, saw := memo[nCurrNum]
	currStr = fmt.Sprintf("%s%d", currStr, multiplier)
  fmt.Printf("currStr %s\n", currStr)
  fmt.Printf("vals %v\n", vals)

	if saw {
    for _, val := range vals {
      if currStr[len(currStr)-(len(val)):] == val {
        return fmt.Sprintf("%s(%s)", currStr[:len(currStr)-(len(val)*2)], val)
      }
    }
	}

  vals = append(vals, fmt.Sprintf("%d", multiplier))
  vals = append(vals, currStr)
	memo[nCurrNum] = vals

	return longDivision(memo, nCurrNum, currDenom, currStr)
}

// FractionToDecimal will return string representation of decimal number
func FractionToDecimal(numerator, denominator int) (ret string) {
	if int(math.Mod(float64(numerator), float64(denominator))) == 0 {
		ret = fmt.Sprintf("%d", numerator/denominator)
		return
	}

	ret = fmt.Sprintf("0.%s", longDivision(make(map[int][]string), numerator, denominator, ""))
  fmt.Printf("ret %s\n", ret)

	return
}
