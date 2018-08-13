package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"github.com/nzgogo/utils/price"
	"time"
)

func MapKeySwap(m map[string]interface{}, n map[string]string) map[string]interface{} {
	r := make(map[string]interface{})
	for k, v := range m {
		if n[k] != "" {
			r[n[k]] = v
		}
	}
	return r
}

func ContainsString(s []string, v string) bool {
	set := make(map[string]struct{}, len(s))
	for _, e := range s {
		set[e] = struct{}{}
	}

	_, ok := set[v]
	return ok
}

func ShortID(s string) (string, error) {
	inputSlice := []rune(s)

	var inputInt []int64

	for i := 0; i < len(s); i = i + 4 {
		n, _ := strconv.ParseInt(string(inputSlice[i:i+4]), 16, 64)
		inputInt = append(inputInt, n)
	}

	for i, v := range inputInt {
		inputInt[i] = remainder(v)
	}

	table := "23456789ABCDEFGHJKLMNPQRSTUVWXYZ"

	var result string

	for i := len(inputInt) - 1; i >= 0; i-- {
		result = result + fmt.Sprintf("%c", table[inputInt[i]])
	}

	return result, nil
}

func remainder(n int64) int64 {
	if n < 32 {
		return n
	}
	var totalRemainder int64
	for n > 0 {
		totalRemainder = totalRemainder + n%32
		n = n / 32
	}
	return remainder(totalRemainder)
}

func FormatAddress(v interface{}) (string, error) {
	type Address struct {
		UnitNumber   string `json:"unitNumber"`
		StreetNumber string `json:"streetNumber"`
		Street       string `json:"street"`
		District     string `json:"district"`
		City         string `json:"city"`
		Zip          string `json:"zip"`
	}
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	a := Address{}
	if err := json.Unmarshal(b, &a); err != nil {
		return "", err
	}

	f := a.StreetNumber + " " + a.Street + ", " + a.District + ", " + a.City + " " + a.Zip

	if a.UnitNumber != "" {
		f = a.UnitNumber + "/" + f
	}

	return f, nil
}

func ReplacePrice(m map[string]interface{}, keys []string) map[string]interface{} {
	keysToBeReplaced := make(map[string]struct{})
	for _, v := range keys {
		keysToBeReplaced[v] = struct{}{}
	}
	for k, v := range m {
		if _, ok := keysToBeReplaced[k]; ok {
			if _, ok := v.(float64); ok {
				m[k], _ = price.WholeToDecimal(fmt.Sprintf("%.0f", v))
			} else {
				m[k], _ = price.WholeToDecimal(fmt.Sprintf("%v", v))
			}
			continue
		}
		iv := reflect.ValueOf(v)
		if iv.IsValid() {
			switch iv.Kind() {
			case reflect.Array, reflect.Slice:
				ns := make([]map[string]interface{}, 0)
				s, _ := v.([]map[string]interface{})
				for _, v := range s {
					siv := reflect.ValueOf(v)
					if !siv.IsValid() || siv.Kind() != reflect.Map {
						ns = s
						break
					}
					ns = append(ns, ReplacePrice(v, keys))
				}
				m[k] = ns
			case reflect.Map:
				nm, _ := v.(map[string]interface{})
				m[k] = ReplacePrice(nm, keys)
			}
		}
	}
	return m
}

// func Distinct(arr interface{}) (reflect.Value, bool) {
// 	// create a slice from our input interface
// 	slice, ok := takeArg(arr, reflect.Slice)
// 	if !ok {
// 		return reflect.Value{}, ok
// 	}
//
// 	// put the values of our slice into a map
// 	// the key's of the map will be the slice's unique values
// 	c := slice.Len()
// 	m := make(map[interface{}]bool)
// 	for i := 0; i < c; i++ {
// 		m[slice.Index(i).Interface()] = true
// 	}
// 	mapLen := len(m)
//
// 	// create the output slice and populate it with the map's keys
// 	out := reflect.MakeSlice(reflect.TypeOf(arr), mapLen, mapLen)
// 	i := 0
// 	for k := range m {
// 		v := reflect.ValueOf(k)
// 		o := out.Index(i)
// 		o.Set(v)
// 		i++
// 	}
//
// 	return out, ok
// }
//
// func takeArg(arg interface{}, kind reflect.Kind) (val reflect.Value, ok bool) {
// 	val = reflect.ValueOf(arg)
// 	if val.Kind() == kind {
// 		ok = true
// 	}
// 	return
// }

// func AppendUnique(slice []interface{}, elems ...interface{}) ([]interface{}, error) {
// 	for _, v := range elems {
// 		et := reflect.SliceOf(reflect.TypeOf(v))
// 		st := reflect.TypeOf(slice)
// 		if et != st {
// 			return nil, errors.New("type mismatch")
// 		}
// 	}
//
// 	set := make(map[interface{}]struct{})
//
// 	return nil, nil
// }

// The Go standard library "time" package has ISOWeek() function for getting ISO 8601 week
// number of a given Time, but there is no reverse functionality for getting a date from a
// week number. This method implements that.
func FirstDayOfISOWeek(year int, week int, timezone *time.Location) time.Time {
	date := time.Date(year, 0, 0, 0, 0, 0, 0, timezone)
	isoYear, isoWeek := date.ISOWeek()

	// iterate back to Monday
	for date.Weekday() != time.Monday {
		date = date.AddDate(0, 0, -1)
		isoYear, isoWeek = date.ISOWeek()
	}

	// iterate forward to the first day of the first week
	for isoYear < year {
		date = date.AddDate(0, 0, 7)
		isoYear, isoWeek = date.ISOWeek()
	}

	// iterate forward to the first day of the given week
	for isoWeek < week {
		date = date.AddDate(0, 0, 7)
		isoYear, isoWeek = date.ISOWeek()
	}

	return date
}
