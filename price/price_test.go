package price

import "testing"

func TestDecimalToWhole(t *testing.T) {
	var num interface{}
	num = float64(66666666.66)
	w, err := DecimalToWhole(num)
	if err != nil {
		t.Error(
			"For", num,
			"expected", 6666666666,
			"got", err.Error(),
		)
	}
	if w != 6666666666 {
		t.Error(
			"For", num,
			"expected", 6666666666,
			"got", w,
		)
	}
}

func TestWholeToDecimal(t *testing.T) {
	var wholeNum interface{}
	wholeNum = int64(6666666666)
	dec, err := WholeToDecimal(wholeNum)
	if err != nil {
		t.Error(
			"For", wholeNum,
			"expected", 66666666.66,
			"got", err.Error(),
		)
	}
	if dec != 66666666.66 {
		t.Error(
			"For", wholeNum,
			"expected", 66666666.66,
			"got", dec,
		)
	}

	wholeNum = "6666666666"
	dec, err = WholeToDecimal(wholeNum)
	if err != nil {
		t.Error(
			"For", wholeNum,
			"expected", 66666666.66,
			"got", err.Error(),
		)
	}
	if dec != 66666666.66 {
		t.Error(
			"For", wholeNum,
			"expected", 66666666.66,
			"got", dec,
		)
	}

	wholeNum = "0.1"
	dec, err = WholeToDecimal(wholeNum)
	if dec != 0 {
		t.Error(
			"For", wholeNum,
			"expected", 0,
			"got", dec,
		)
	}
}
