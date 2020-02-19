package price

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	valid "github.com/asaskevich/govalidator"
)

func Round(x float64, place int) string {
	precision := math.Pow(0.1, float64(place))
	resultFormat := "%." + strconv.Itoa(place) + "f"

	return fmt.Sprintf(resultFormat, math.Round(x/precision)*precision)
}

func RoundHalfToEven(n interface{}) (int64, error) {
	numberString := valid.ToString(n)
	if i, err := valid.ToInt(numberString); err == nil {
		return i, nil
	}

	if f, err := valid.ToFloat(numberString); err == nil {
		rounded := Round(f, 1)
		splitted := strings.Split(rounded, ".")
		wholePlace, _ := strconv.ParseInt(splitted[0], 0, 64)
		decimalPlace, _ := strconv.ParseInt(splitted[1], 0, 64)
		if wholePlace >= 0 {
			if decimalPlace <= 4 {
				return wholePlace, nil
			} else if decimalPlace >= 6 {
				return wholePlace + 1, nil
			} else {
				if wholePlace%2 == 0 {
					return wholePlace, nil
				} else {
					return wholePlace + 1, nil
				}
			}
		} else {
			if decimalPlace <= 4 {
				return wholePlace, nil
			} else if decimalPlace >= 6 {
				return wholePlace - 1, nil
			} else {
				if wholePlace%2 == 0 {
					return wholePlace, nil
				} else {
					return wholePlace - 1, nil
				}
			}
		}
	}

	return 0, errors.New("failed to round")
}

func DecimalToWhole(n interface{}) (int64, error) {
	numberString := valid.ToString(n)
	f, err := valid.ToFloat(numberString)
	if err != nil {
		return 0, errors.New("not a number")
	}

	return RoundHalfToEven(f * 100)
}

func WholeToDecimal(n interface{}) (float64, error) {
	i, err := valid.ToInt(n)
	if err != nil {
		return 0, errors.New("not a whole number")
	}
	fs := fmt.Sprintf("%.2f", float64(i)/100)
	return valid.ToFloat(fs)
}

func CalculateFeeFromDistance(distance int64) (fee int64) {
	fee = 499
	if distance > 1000 && distance <= 4500 {
		fee += int64(math.Ceil((float64(distance)-1000)/500) * 75)
	}
	if distance > 4500 && distance <= 9500 {
		fee += int64(math.Ceil((float64(distance)-1000)/500) * 90)
	}
	if distance > 9500 && distance <= 14500 {
		fee += int64(math.Ceil((float64(distance)-1000)/500) * 115)
	}
	if distance > 14500 {
		fee += int64(math.Ceil((float64(distance)-1000)/500) * 125)
	}
	return fee
}

func CalculateDriverBenifit(distance int64) (int64, error) {
	fee := CalculateFeeFromDistance(distance)
	rate := int64(80)
	if distance >= 3000 && distance < 10000 {
		rate = int64(70)
	}
	if distance >= 10000 {
		rate = int64(65)
	}
	return RoundHalfToEven(fee * rate / 100)
}
