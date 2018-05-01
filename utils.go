package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"encoding/hex"
	"strings"
)

func MapKeyToCamelCase(m map[string]interface{}, d string) map[string]interface{} {
	for k, v := range m {
		nk := ""
		s := strings.Split(k, d)
		if len(s) > 1 {
			var r []string
			for _, v := range s {
				if v != "" {
					r = append(r, v)
				}
			}
			for i, kp := range s {
				kp = strings.ToLower(kp)
				if i == 0 {
					nk = nk + kp
				} else {
					nk = nk + strings.Title(kp)
				}
			}
		} else {
			nk = k
		}
		m[nk] = v
		if len(s) > 1 {
			delete(m, k)
		}
	}
	return m
}

func Contains(s interface{}, v interface{}) bool {
	sSlice := s.([]interface{})
	vBytes, err := getBytes(v)
	if err != nil {
		return false
	}
	vMd5Bytes := md5.Sum(vBytes)
	vString := hex.EncodeToString(vMd5Bytes[:])
	set := make(map[string]struct{}, len(sSlice))
	for _, e := range sSlice {
		eBytes, err := getBytes(e)
		if err != nil {
			continue
		}
		eMd5Bytes := md5.Sum(eBytes)
		eString := hex.EncodeToString(eMd5Bytes[:])
		set[eString] = struct{}{}
	}

	_, ok := set[vString]
	return ok
}

func getBytes(value interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(value); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
