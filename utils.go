package utils

import (
	"fmt"
	"strconv"
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
