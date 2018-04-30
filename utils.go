package utils

import "strings"

func MapKeyToCamelCase(m map[string]interface{}, d string) map[string]interface{} {
  for k, v := range m {
    nk := ""
    s := strings.Split(k, d)
    if len(s) > 1 {
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
    delete(m, k)
  }
  return m
}
