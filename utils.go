package utils

import (
	"encoding/hex"
	"fmt"
	"strconv"
)

// func MapKeyToCamelCase(m map[string]interface{}, d string) map[string]interface{} {
// 	for k, v := range m {
// 		nk := ""
// 		s := strings.Split(k, d)
// 		if len(s) > 1 {
// 			var r []string
// 			for _, v := range s {
// 				if v != "" {
// 					r = append(r, v)
// 				}
// 			}
// 			for i, kp := range s {
// 				kp = strings.ToLower(kp)
// 				if i == 0 {
// 					nk = nk + kp
// 				} else {
// 					nk = nk + strings.Title(kp)
// 				}
// 			}
// 		} else {
// 			nk = k
// 		}
// 		m[nk] = v
// 		if len(s) > 1 {
// 			delete(m, k)
// 		}
// 	}
// 	return m
// }

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
	src := []byte(s)
	dst := make([]byte, hex.DecodedLen(len(src)))
	if _, err := hex.Decode(dst, src); err != nil {
		return "", err
	}

	inputSlice := []rune(s)

	var inputInt []int64

	for i := 0; i < len(s); i = i + 4 {
		fmt.Println(string(inputSlice[i : i+4]))
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
	if n <= 32 {
		return n
	}
	var totalRemainder int64
	for n > 0 {
		totalRemainder = totalRemainder + n%32
		n = n / 32
	}
	return remainder(totalRemainder)
}
