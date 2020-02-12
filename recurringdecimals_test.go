package recurringdecimals

import "testing"
import "encoding/json"
import "os"
import "bufio"
import "io"

type Test struct {
  Num int `json:"numerator"`
  Denom int `json:"denominator"`
  Output string `json:"output"`
}

func TestFractionToDecimal(testUtil *testing.T) {

  f, err := os.Open("./tests.json")

  if err != nil {
    testUtil.Error(err)
    return
  }

  defer f.Close()
  
  reader := bufio.NewReader(f)
  decoder := json.NewDecoder(reader)
  var tests map[string]Test

  for {

    err = decoder.Decode(&tests)

    if err == nil {

      for name, test := range tests {
        testUtil.Run(name, func (tstUtil *testing.T) {
          if FractionToDecimal(test.Num, test.Denom) != test.Output {
            tstUtil.Errorf("use case %v\n", test)
          }
        })
      }
    } else if err == io.EOF {
      break
    } else {
      testUtil.Error(err)
      return
    }
  }

}