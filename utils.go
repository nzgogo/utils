package utils

import (
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

func ContainsString(s []string, v string) bool {
	set := make(map[string]struct{}, len(s))
	for _, e := range s {
		set[e] = struct{}{}
	}

	_, ok := set[v]
	return ok
}
